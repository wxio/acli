package newsubcmd

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/golang/glog"
	"github.com/iancoleman/strcase"
	"github.com/wxio/acli/internal/types"
)

type newsubcmdOpt struct {
	rt    *types.Root
	Debug bool

	Name       []string `opts:"mode=arg"`
	Parent     string   `help:"path of parent commands eg foo/bar"`
	Org        string
	Project    string
	ModulePath string `help:"the parent path of the internal src directory"`
	Overwrite  bool
	EntireReg  bool `help:"(only valid with --parent) print to standard output an entire go file to register the subcommand. If false only the func init is printed."`

	err error
}

// New constructor for init
func New(rt *types.Root) interface{} {
	parts := strings.Split(os.Args[0], "/")
	in := newsubcmdOpt{
		rt:      rt,
		Org:     "wxio",
		Project: parts[len(parts)-1],
	}
	absPath, err := os.Getwd()
	if _, err := os.Open(absPath + "/go.mod"); err != nil {
		err = fmt.Errorf("no go.mod in current directory")
	}
	if err == nil {
		in.ModulePath = absPath
	} else {
		in.err = err
	}
	return &in
}

//go:embed subcmd.tmpl
var fs embed.FS

func (in *newsubcmdOpt) Run() error {
	in.rt.Config(in)
	if in.err != nil {
		fmt.Fprintf(os.Stderr, "couldn't get executable's path %v\n", in.err)
		os.Exit(1)
	}
	if len(in.Name) == 0 {
		return fmt.Errorf("Name(s) required")
	}
	funcMap := template.FuncMap{
		"ToUpper":      strings.ToUpper,
		"ToCamel":      strcase.ToCamel,
		"ToLowerCamel": strcase.ToLowerCamel,
	}
	tmpl, err := template.New("").Funcs(funcMap).ParseFS(fs, "*.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Template error %v\n", err)
		glog.Fatalf("Template error %v", err)
	}
	data := struct {
		Name    string
		Parent  []string
		Org     string
		Project string
		Path    string
	}{
		// Name:    Name,
		Path:    in.Parent,
		Parent:  strings.Split(in.Parent, "/"),
		Org:     in.Org,
		Project: in.Project,
	}
	fmt.Fprintf(os.Stderr, "written starter code\n")
	for _, name := range in.Name {
		data.Name = name
		in.makeStarter(name, tmpl, data)
	}
	if in.Parent == "" && in.EntireReg {
		fmt.Fprintf(os.Stderr, "\nWarning --entire-reg only valid if --parent is specified\n")
	}
	if in.Parent == "" {
		fmt.Fprintf(os.Stderr, "\nMain needs to be manually modified, sample below\n")
		fmt.Fprintf(os.Stderr, "``` golang\n")
		for _, name := range in.Name {
			data.Name = name
			err := tmpl.Lookup("mainreg").Execute(os.Stdout, data)
			if err != nil {
				fmt.Printf("template exec error %v\n", err)
			}
		}
	} else {
		data := struct {
			Names   []string
			Parent  []string
			Path    string
			Org     string
			Project string
		}{
			Names:   in.Name,
			Parent:  strings.Split(in.Parent, "/"),
			Path:    in.Parent,
			Org:     in.Org,
			Project: in.Project,
		}
		if in.EntireReg {
			fmt.Fprintf(os.Stderr, "\nA go file in package main needs to be manually created, sample content below\n")
			fmt.Fprintf(os.Stderr, "``` golang\n")
			err = tmpl.Lookup("file_header").Execute(os.Stdout, data)
			if err != nil {
				fmt.Printf("template 'file_header' exec error %v\n", err)
				return err
			}
		} else {
			fmt.Fprintf(os.Stderr, "\nMain needs to be manually modified, sample below\n")
			fmt.Fprintf(os.Stderr, "``` golang\n")
		}
		err = tmpl.Lookup("mainregwithparent_addstruct").Execute(os.Stdout, data)
		if err != nil {
			fmt.Printf("template exec error %v\n", err)
		}
	}
	fmt.Fprintf(os.Stderr, "```\n")
	return nil
}

func (in *newsubcmdOpt) makeStarter(name string, tmpl *template.Template, data any) {
	dirname := in.ModulePath + "/internal/" + name
	if in.Parent != "" {
		dirname = in.ModulePath + "/internal/" + in.Parent + "/" + name
	}
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create dir error %v\n", err)
	}
	fname := dirname + "/" + name + ".go"
	if !in.Overwrite {
		if _, err = os.Open(fname); err == nil {
			fmt.Fprintf(os.Stderr, "Exiting. File already exists. Use --overwrite to ignore.\n")
			fmt.Fprintf(os.Stderr, "  %s\n", fname)
			os.Exit(3)
		}
	}
	fh, err := os.Create(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file error %v\n", err)
		os.Exit(1)
	}
	tmpl.Lookup("newsubcmd").Execute(fh, data)
	fh.Close()
	fmt.Fprintf(os.Stderr, "  %v\n", fname)
}

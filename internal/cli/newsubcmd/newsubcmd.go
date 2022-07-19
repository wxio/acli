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

	Name       string //`opts:"mode=arg"`
	Parent     string `help:"path of parent commands eg foo/bar"`
	Org        string
	Project    string
	ModulePath string `help:"the parent path of the internal src directory"`

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
	absPath, err := os.Executable()
	if err == nil {
		parts := strings.Split(absPath, "/")
		in.ModulePath = strings.Join(parts[0:len(parts)-1], "/")
	} else {
		in.err = err
	}
	return &in
}

//go:embed subcmd.tmpl
var fs embed.FS

func (in *newsubcmdOpt) Run() {
	in.rt.Config(in)
	if in.err != nil {
		fmt.Printf("could get executable's path %v\n", in.err)
		os.Exit(1)
	}
	if in.Name == "" {
		fmt.Printf("Name (--name) required\n")
		os.Exit(1)
	}
	funcMap := template.FuncMap{
		"ToUpper":      strings.ToUpper,
		"ToCamel":      strcase.ToCamel,
		"ToLowerCamel": strcase.ToLowerCamel,
	}
	tmpl, err := template.New("").Funcs(funcMap).ParseFS(fs, "*.tmpl")
	if err != nil {
		fmt.Printf("Template error %v\n", err)
		glog.Fatalf("Template error %v", err)
	}

	dirname := in.ModulePath + "/internal/" + in.Name
	if in.Parent != "" {
		dirname = in.ModulePath + "/internal/" + in.Parent + "/" + in.Name
	}
	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		fmt.Printf("create dir error %v\n", err)
	}
	fname := dirname + "/" + in.Name + ".go"
	fh, err := os.Create(fname)
	if err != nil {
		fmt.Printf("create file error %v\n", err)
		os.Exit(1)
	}
	data := struct {
		Name    string
		Parent  []string
		Org     string
		Project string
	}{
		Name:    in.Name,
		Parent:  strings.Split(in.Parent, "/"),
		Org:     in.Org,
		Project: in.Project,
	}
	tmpl.Lookup("newsubcmd").Execute(fh, data)
	fh.Close()
	fmt.Printf("written starter code to '%v'\n", fname)
	if in.Parent != "" {
		tmpl.Lookup("mainregwithparent").Execute(os.Stdout, data)
	} else {
		tmpl.Lookup("mainreg").Execute(os.Stdout, data)
	}
}

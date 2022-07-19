package newsubcmd

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/golang/glog"
	"github.com/wxio/acli/internal/types"
)

type newsubcmdOpt struct {
	rt    *types.Root
	Debug bool

	Name    string //`opts:"mode=arg"`
	Org     string
	Project string
}

// New constructor for init
func New(rt *types.Root) interface{} {
	parts := strings.Split(os.Args[0], "/")
	in := newsubcmdOpt{
		rt:      rt,
		Org:     "wxio",
		Project: parts[len(parts)-1],
	}
	return &in
}

//go:embed subcmd.tmpl
var fs embed.FS

func (in *newsubcmdOpt) Run() {
	in.rt.Config(in)
	if in.Name == "" {
		fmt.Printf("Name (--name) required\n")
		os.Exit(1)
	}
	funcMap := template.FuncMap{
		"ToUpper":      strings.ToUpper,
		"ToCamel":      ToCamel,
		"ToLowerCamel": ToLowerCamel,
	}
	tmpl, err := template.New("").Funcs(funcMap).ParseFS(fs, "*.tmpl")
	if err != nil {
		fmt.Printf("Template error %v\n", err)
		glog.Fatalf("Template error %v", err)
	}
	dirname := "internal/" + in.Name
	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		fmt.Printf("create dir error %v\n", err)
	}
	fname := dirname + "/" + in.Name + ".go"
	fh, err := os.Create("internal/" + in.Name + "/" + in.Name + ".go")
	if err != nil {
		fmt.Printf("create file error %v\n", err)
	}
	tmpl.Lookup("newsubcmd").Execute(fh, in)
	fh.Close()
	fmt.Printf("written starter code to '%v'\n", fname)
	tmpl.Lookup("mainreg").Execute(os.Stdout, in)
}

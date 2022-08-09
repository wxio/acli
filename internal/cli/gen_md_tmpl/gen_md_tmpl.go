package gen_md_tmpl

import (
	"fmt"
	"os"
	"strings"

	"github.com/wxio/acli/internal/types"
)

type genMdTmplOpt struct {
	rt *types.Root
}

func NewGenMdTmpl(rt *types.Root) interface{} {
	in := genMdTmplOpt{
		rt: rt,
	}
	return &in
}

func (in *genMdTmplOpt) Run() error {
	root := &types.RoseTree{Name: "acli"}
	root.Collect(in.rt.Cli.Children())
	fmt.Fprintf(os.Stderr, "Place the contents below in a readme and run md_tmpl eg `acli cli md_tmpl README.md`\n\n")
	fn := func(path []string, subcmd types.RoseTree) {
		fmt.Printf("## `%s %s`\n", strings.Join(path, " "), subcmd.Name)
		if !subcmd.Cmd.IsRunnable() {
			fmt.Printf("*sub-command is not runnable, only exists for grouping purposes*\n")
		}
		fmt.Printf("<!--tmpl,code:%s %s --help --><!--/tmpl-->\n\n", strings.Join(path, " "), subcmd.Name)
	}
	root.Call([]string{root.Name}, fn)
	return nil
}

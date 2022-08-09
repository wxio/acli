package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/jpillora/opts"
	"github.com/wxio/acli/internal/cli/md_tmpl"
	"github.com/wxio/acli/internal/cli/newsubcmd"
	"github.com/wxio/acli/internal/cli/rename"
	"github.com/wxio/acli/internal/types"
)

// Set by build tool chain by
// go build --ldflags '-X main.XXX=xxx -X main.YYY=yyy -X main.ZZZ=zzz'
var (
	ProjectName string = "?"
	Version     string = "dev"
	Date        string = "na"
	Commit      string = "na"
	ReleaseURL  string = "na"
)

type versionCmd struct{}

func (r *versionCmd) Run() {
	fmt.Printf(`%s
version: %s
date:    %s
commit:  %s
release: %s
`, ProjectName, Version, Date, Commit, ReleaseURL)
}

var (
	rflg    = &types.Root{}
	cliBldr = opts.New(rflg).
		Name("acli").
		EmbedGlobalFlagSet().
		Complete()
)

type roseTree struct {
	name     string
	runnable bool
	kids     []*roseTree
}

func (a roseTree) Len() int      { return len(a.kids) }
func (a roseTree) Swap(i, j int) { a.kids[i], a.kids[j] = a.kids[j], a.kids[i] }
func (a roseTree) Less(i, j int) bool {
	return strings.Compare(a.kids[i].name, a.kids[j].name) < 0
}
func (rt *roseTree) append(cmds map[string]opts.Opts) {
	for name, cmd := range cmds {
		parsed, _ := cmd.ParseArgsError([]string{})
		next := &roseTree{name: name, runnable: parsed.IsRunnable()}
		rt.kids = append(rt.kids, next)
		next.append(parsed.Children())
	}
}

func (rt *roseTree) print(path []string) {
	sort.Sort(rt)
	for _, subcmd := range rt.kids {
		fmt.Printf("## `%s %s`\n", strings.Join(path, " "), subcmd.name)
		if !subcmd.runnable {
			fmt.Printf("*sub-command is not runnable, only exists for grouping purposes*\n")
		}
		fmt.Printf("<!--tmpl,code:%s %s --help --><!--/tmpl-->\n\n", strings.Join(path, " "), subcmd.name)
		subcmd.print(append(path, subcmd.name))
	}
}

func main() {
	cli := cliBldr.Parse()
	a := reflect.ValueOf(cli.Selected()).Elem().Addr()
	b := reflect.ValueOf(cliBldr).Elem().Addr()
	if a == b {
		if rflg.GenDocs {
			root := &roseTree{name: "acli"}
			root.append(cli.Children())
			fmt.Fprintf(os.Stderr, "Place the contents below in a readme and run md_tmpl eg `acli cli md_tmpl README.md`\n\n")
			root.print([]string{root.name})
			os.Exit(0)
		}
	}
	err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\nError: %v\n\n", cli.Selected().Help(), err)
		os.Exit(1)
	}
}

func init() {
	cliBldr.AddCommand(opts.New(&versionCmd{}).Name("version"))
}

func init() {
	cliBldr.AddCommand(opts.NewPlaceholder("cli").
		Summary("Grouping of Command Line Interface subcommands").
		AddCommand(opts.New(rename.NewRename(rflg)).Name("rename").
			Summary("Does a search and replace on the name 'acli'* in all the source file.\nRespects .gitignore.")).
		AddCommand(opts.New(newsubcmd.New(rflg)).Name("new_sub_command").
			Summary("Creates starter code for new subcommands & print sample registration code for the main package.\n" +
				"\n" +
				"Starter code is written to  in `internal/parent/subcommand/subcommand.go`\n" +
				"\n" +
				"Use --parent to nest the subcommands.\n" +
				"\n" +
				"Sample code is written to stdOut, all other output is on stdErr, this allows for the following\n" +
				"`acli cli new_sub_command subcmd2 >> main.go`\n" +
				"or\n" +
				"`acli cli new_sub_command --entire-reg --parent topsy/turvy top1 top2 > reg02.go`",
			)).
		AddCommand(opts.New(md_tmpl.NewMdTmpl(rflg)).Name("md_tmpl").
			Summary("Simple markdown templating using shell commands. see https://github.com/jpillora/md-tmpl\n" +
				"General usage is `acli cli md_tmpl README.md`.\n" +
				"Works nicely with `acli --gen-docs >> README.md`")).
		End())
}

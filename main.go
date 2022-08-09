package main

import (
	"debug/buildinfo"
	"fmt"
	"os"

	"github.com/jpillora/opts"
	"github.com/wxio/acli/internal/cli/gen_md_tmpl"
	"github.com/wxio/acli/internal/cli/md_tmpl"
	"github.com/wxio/acli/internal/cli/newsubcmd"
	"github.com/wxio/acli/internal/cli/rename"
	"github.com/wxio/acli/internal/listtasks"
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

type versionCmd struct {
	ModuleInfo bool
}

func (r *versionCmd) Run() error {
	fmt.Printf(`%s
version: %s
date:    %s
commit:  %s
release: %s
`, ProjectName, Version, Date, Commit, ReleaseURL)
	if r.ModuleInfo {
		file, err := os.Executable()
		if err != nil {
			return err
		}
		bi, err := buildinfo.ReadFile(file)
		if err != nil {
			return err
		}
		fmt.Printf("\n%s\n", bi)
	}
	return nil
}

var (
	rflg    = &types.Root{}
	cliBldr = opts.New(rflg).
		Name("acli").
		EmbedGlobalFlagSet().
		Complete()
)

func main() {
	cli := cliBldr.Parse()
	rflg.Cli = cli
	err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\nError: %v\n\n", cli.Selected().Help(), err)
		os.Exit(1)
	}
}

func init() {
	cliBldr.
		AddCommand(opts.New(&versionCmd{}).Name("version").
			Summary("Prints version information injected by the build.")).
		AddCommand(opts.New(listtasks.NewListtasks(rflg)).Name("listtasks").
			Summary("List tasks with short summary").
			End())
}

func init() {
	cliBldr.AddCommand(opts.NewPlaceholder("cli").
		Summary("Grouping of Command Line Interface subcommands.").
		AddCommand(opts.New(gen_md_tmpl.NewGenMdTmpl(rflg)).Name("gen_md_tmpl").
			Summary("Generates markdown templates to add to readme.md")).
		AddCommand(opts.New(rename.NewRename(rflg)).Name("rename").
			Summary("Does a search and replace on the name in all the source file.").
			LongDescription("Respects .gitignore."),
		).
		AddCommand(opts.New(newsubcmd.New(rflg)).Name("new_sub_command").
			Summary("Creates starter code for subcommands & print sample registration code for the main package.").
			LongDescription("Starter code is written to  in `internal/parent/subcommand/subcommand.go`\n" +
				"\n" +
				"Use --parent to nest the subcommands.\n" +
				"\n" +
				"Sample code is written to stdOut, all other output is on stdErr, this allows for the following\n" +
				"`acli cli new_sub_command subcmd2 >> main.go`\n" +
				"or\n" +
				"`acli cli new_sub_command --entire-reg --parent topsy/turvy top1 top2 > reg02.go`",
			)).
		AddCommand(opts.New(md_tmpl.NewMdTmpl(rflg)).Name("md_tmpl").
			Summary("Simple markdown templating using shell commands. see https://github.com/jpillora/md-tmpl").
			LongDescription("General usage is `acli cli md_tmpl README.md`.\n" +
				"Works nicely with `acli --gen-docs >> README.md`")).
		End())
}

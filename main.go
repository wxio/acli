package main

import (
	"fmt"
	"os"

	"github.com/jpillora/opts"
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

func main() {
	cli := cliBldr.Parse()
	cmd, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v%s\n", err, cmd.Help())
		os.Exit(1)
	}
}

func init() {
	cliBldr.AddCommand(opts.New(&versionCmd{}).Name("version"))
}

func init() {
	cliBldr.AddCommand(opts.New(&struct{}{}).Name("cli").
		AddCommand(opts.New(rename.NewRename(rflg)).Name("rename")).
		AddCommand(opts.New(newsubcmd.New(rflg)).Name("new_sub_command")))
}

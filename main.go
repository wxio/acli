package main

import (
	"fmt"

	"github.com/jpillora/opts"
	"github.com/wxio/acli/internal/cli/newsubcmd"
	"github.com/wxio/acli/internal/cli/rename"
	"github.com/wxio/acli/internal/types"
)

func main() {
	rflg := &types.Root{}
	op := opts.New(rflg).
		Name("acli").
		EmbedGlobalFlagSet().
		Complete().
		AddCommand(opts.New(&versionCmd{}).Name("version")).
		AddCommand(
			opts.New(&struct{}{}).Name("cli").
				AddCommand(opts.New(rename.NewRename(rflg)).Name("rename")).
				AddCommand(opts.New(newsubcmd.New(rflg)).Name("new_sub_command"))).
		Parse()
	op.RunFatal()
}

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

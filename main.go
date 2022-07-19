package main

import (
	"fmt"

	"github.com/jpillora/opts"
	"github.com/wxio/acli/internal/newsubcmd"
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
				AddCommand(
					opts.New(newsubcmd.New(rflg)).Name("new_sub_command"), //.Summary(newsubcmd.Usage),
				),
		).
		Parse()
	op.RunFatal()
}

// Set by build tool chain by
// go build --ldflags '-X main.Version=xxx -X main.Date=xxx -X main.Commit=xxx'
var (
	Version string = "dev"
	Date    string = "na"
	Commit  string = "na"
)

type versionCmd struct{}

func (r *versionCmd) Run() {
	fmt.Printf("version: %s\ndate: %s\ncommit: %s\n", Version, Date, Commit)
}

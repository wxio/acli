package newsubcmd

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/wxio/acli/internal/types"
)

type newsubcmdOpt struct {
	rt    *types.Root
	Debug bool

	Name   string //`opts:"mode=arg"`
	Parent string `help:"forward slash (/) delimited path"`
}

// New constructor for init
func New(rt *types.Root) interface{} {
	in := newsubcmdOpt{
		rt: rt,
	}
	return &in
}

func (in *newsubcmdOpt) Run() {
	in.rt.Config(in)
	glog.Infof("%v\n", *in)
	fmt.Printf("todo implement code gen for new subcommand and print registration pattern or modify the main.go file\n")
}

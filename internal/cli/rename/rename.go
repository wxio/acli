package rename

import (
	"fmt"

	"github.com/wxio/acli/internal/types"
)

type renameOpt struct {
	rt    *types.Root
	Debug bool
	// public fields are flags by default.
	// Use annotations to adjust eg. `opts:"mode=arg"`
	Org  string
	Name string
}

func NewRename(rt *types.Root) interface{} {
	in := renameOpt{
		rt: rt,
	}
	return &in
}

func (in *renameOpt) Run() {
	in.rt.Config(in)
	fmt.Printf("todo implement rename\n")
}

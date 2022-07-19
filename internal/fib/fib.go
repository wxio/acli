package fib

import (
	"fmt"
	"os"

	"github.com/wxio/acli/internal/types"
)

type fibOpt struct {
	rt    *types.Root
	Debug bool

	Nth  int    `opts:"mode=arg"`
	Impl string `help:"i:iterative r:recursive c:channels"`
}

func NewFib(rt *types.Root) interface{} {
	in := fibOpt{
		rt:   rt,
		Impl: "i",
	}
	return &in
}

func (in *fibOpt) Run() {
	in.rt.Config(in)
	switch in.Impl {
	case "i":
		fmt.Printf("nth fib is %v\n", fibIterative(in.Nth))
	case "r":
		fmt.Printf("nth fib is %v\n", fibRecursive(in.Nth))
	case "c":
	default:
		fmt.Printf("only valid impl option are i|r|c not '%s'\n", in.Impl)
		os.Exit(2)
	}
}

func fibIterative(nth int) int {
	if nth <= 2 {
		return 1
	}
	n0, n1 := 1, 1
	idx := 3
	for {
		if idx >= nth {
			return n0 + n1
		}
		n0, n1 = n1, n0+n1
		idx++
	}
}

func fibRecursive(nth int) int {
	if nth <= 2 {
		return 1
	}
	return fibRecursive(nth-2) + fibRecursive(nth-1)
}

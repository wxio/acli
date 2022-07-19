package fib

import (
	"fmt"
	"log"
	"os"

	"net/http"
	_ "net/http/pprof"

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
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	switch in.Impl {
	case "i":
		fmt.Printf("interactive: nth fib is %v\n", fibIterative(in.Nth))
	case "r":
		fmt.Printf("recursive: nth fib is %v\n", fibRecursive(in.Nth))
	case "c":
		fmt.Printf("concurrent: nth fib is %v\n", fibChannel(in.Nth))
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

func nthCell(idx int, in0, in1 chan int) (out chan int) {
	out = make(chan int, 2)
	go func() {
		n0, n1 := <-in0, <-in1
		s := n0 + n1
		out <- s
		out <- s
	}()
	return
}

func fibChannel(nth int) int {
	in := make(chan int, 2)
	out0, out1 := in, nthCell(3, in, in)
	for i := 4; i <= nth; i++ {
		out0, out1 = out1, nthCell(i, out0, out1)
	}
	in <- 1
	in <- 1
	in <- 1
	_ = <-out0
	r1 := <-out1
	return r1
}

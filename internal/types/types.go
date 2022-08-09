package types

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jpillora/opts"
)

// Root struct for commands
// Field tags control where the values come from
// If opts:"-" yaml:"-" are set in object creation
//    opts:="-" come from config file
//    yaml:="-" come from command line flags
type Root struct {
	Cfg        string          `help:"Config file in json format (NOTE file entries take precedence over command-line flags & env)" json:"-"`
	DumpConfig bool            `help:"Dump the config to stdout and exits" json:"-"`
	Cli        opts.ParsedOpts `opts:"-"`
}

func (rt Root) Config(in interface{}) {
	if rt.Cfg != "" {
		fd, err := os.Open(rt.Cfg)
		// config is in its own func
		// this defer fire correctly
		//
		// won't fire if dump is used as os.Exit terminates program
		defer func() {
			fd.Close()
		}()
		if err != nil {
			log.Fatalf("error opening file %s %v", rt.Cfg, err)
		}
		dec := json.NewDecoder(fd)
		err = dec.Decode(in)
		if err != nil {
			log.Fatalf("json error %v", err)
		}
	}
	if rt.DumpConfig {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		err := enc.Encode(in)
		if err != nil {
			log.Fatalf("json encoding error %v", err)
		}
		os.Exit(0)
	}
}

type RoseTree struct {
	Cmd  opts.ParsedOpts
	Name string
	Kids []*RoseTree
}

func (a RoseTree) Len() int      { return len(a.Kids) }
func (a RoseTree) Swap(i, j int) { a.Kids[i], a.Kids[j] = a.Kids[j], a.Kids[i] }
func (a RoseTree) Less(i, j int) bool {
	return strings.Compare(a.Kids[i].Name, a.Kids[j].Name) < 0
}
func (rt *RoseTree) Collect(cmds map[string]opts.Opts) {
	for name, cmd := range cmds {
		parsed, _ := cmd.ParseArgsError([]string{})
		next := &RoseTree{Name: name, Cmd: parsed}
		rt.Kids = append(rt.Kids, next)
		next.Collect(parsed.Children())
	}
}

func (rt *RoseTree) Call(path []string, fn func([]string, RoseTree)) {
	sort.Sort(rt)
	for _, subcmd := range rt.Kids {
		fn(path, *subcmd)
		subcmd.Call(append(path, subcmd.Name), fn)
	}
}

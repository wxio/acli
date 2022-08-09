package listtasks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wxio/acli/internal/types"
)

type listtasksOpt struct {
	rt          *types.Root
	IncludeHelp bool
}

func NewListtasks(rt *types.Root) interface{} {
	in := listtasksOpt{
		rt: rt,
	}
	return &in
}

func (in *listtasksOpt) Run() error {
	in.rt.Config(in)
	root := &types.RoseTree{Name: "acli"}
	root.Collect(in.rt.Cli.Children())
	max := 0
	list := [][]string{}
	fn := func(path []string, subcmd types.RoseTree) {
		if subcmd.Cmd.IsRunnable() {
			cmd := strings.Join(append(path, subcmd.Name), " ")
			if len(cmd) > max {
				max = len(cmd)
			}
			list = append(list, []string{cmd, subcmd.Cmd.GetSummary()})
			if in.IncludeHelp {
				list = append(list, []string{"", subcmd.Cmd.Help()})
			}
		}
	}
	root.Call([]string{root.Name}, fn)
	for _, task := range list {
		fmt.Printf("%-"+strconv.Itoa(max)+"s %s\n", task[0], task[1])
	}
	return nil
}

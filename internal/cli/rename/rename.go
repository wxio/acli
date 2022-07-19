package rename

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	ignore "github.com/sabhiram/go-gitignore"
	"github.com/wxio/acli/internal/types"
)

type renameOpt struct {
	rt    *types.Root
	Debug bool
	// public fields are flags by default.
	// Use annotations to adjust eg. `opts:"mode=arg"`
	Org        string
	Name       string
	ModulePath string `help:"the parent path of the internal src directory"`
	err        error
}

func NewRename(rt *types.Root) interface{} {
	in := renameOpt{
		rt: rt,
	}
	absPath, err := os.Executable()
	if err == nil {
		parts := strings.Split(absPath, "/")
		in.ModulePath = strings.Join(parts[0:len(parts)-1], "/")
	} else {
		in.err = err
	}
	return &in
}

func (in *renameOpt) Run() {
	in.rt.Config(in)
	if in.err != nil {
		fmt.Printf("could get executable's path %v\n", in.err)
		os.Exit(1)
	}
	if in.Name == "" {
		fmt.Printf("Name (--name) required\n")
		os.Exit(1)
	}
	WalkDirSrc(in.ModulePath, func(path string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil {
			fmt.Printf("!!! %v\n", err)
			return err
		}
		m := info.Mode()
		fmt.Printf("path:'%v'   name:'%s' dir:%v mode:%v\n", path, d.Name(), d.Type().IsDir(), m.String())
		return nil
	})
	fmt.Printf("todo implement rename\n")
}

// walk a directory calling fn on src files.
// ie no directories, no files with match .gitignore files
func WalkDirSrc(root string, fn fs.WalkDirFunc) error {
	type GIS struct {
		gi   *ignore.GitIgnore
		path string
	}
	var gis []GIS
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err0 error) error {
		if err0 != nil {
			fmt.Printf("!! %v\n", err0)
			return err0
		}
		if d.Name() == ".git" {
			return fs.SkipDir
		}
		if d.IsDir() {
			gi, err := ignore.CompileIgnoreFile(path + "/.gitignore")
			if err == nil {
				gis = append(gis, GIS{gi, path})
			}
			return nil
		}
		for _, gi := range gis {
			if strings.HasPrefix(path, gi.path) {
				subpath := path[len(gi.path):]
				if gi.gi.MatchesPath(subpath) {
					return nil
				}
			}
		}
		return fn(path, d, err0)
	})
}

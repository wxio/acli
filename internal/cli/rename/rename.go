package rename

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
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
	FromTo     []string `help:"src dest eg wxio/acli freddo/frog or acli frog" opts:"mode=arg"`
	ModulePath string   `help:"the parent path of the internal src directory"`
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
	if len(in.FromTo) != 2 {
		fmt.Printf("'from_name to_name' must exist\n")
		os.Exit(1)
	}
	from, to := strings.Split(in.FromTo[0], "/"), strings.Split(in.FromTo[1], "/")
	if len(from) != len(to) && (len(to) == 1 || len(to) == 2) {
		fmt.Printf("'from_name to_name' must must look like 'a/b' or just 'b' and be the same\n")
		os.Exit(1)
	}
	err := WalkDirSrc(in.ModulePath, func(path string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil {
			fmt.Printf("!!! %v\n", err)
			return err
		}
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read file error %v", err)
		}
		for i := range from {
			contents = bytes.ReplaceAll(contents, []byte(from[i]), []byte(to[i]))
		}
		return ioutil.WriteFile(path, contents, info.Mode())
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
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

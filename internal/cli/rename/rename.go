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
	FromOrg    string   `opts:"short=o"`
	FromName   string   `opts:"short=n"`
	To         []string `help:"* to is either 'org name' or just 'name' eg 'freddo frog' or 'frog'" opts:"mode=arg"`
	ModulePath string   `help:"the parent path of the internal src directory"`
	err        error
}

func NewRename(rt *types.Root) interface{} {
	in := renameOpt{
		rt: rt,
	}
	absPath, err := os.Getwd()
	if _, err := os.Open(absPath + "/go.mod"); err != nil {
		err = fmt.Errorf("no go.mod in current directory")
	}
	if err == nil {
		in.ModulePath = absPath
		in.FromOrg = "wxio"
		in.FromName = "acli"
	} else {
		in.err = err
	}
	return &in
}

func (in *renameOpt) Run() error {
	in.rt.Config(in)
	if in.err != nil {
		return fmt.Errorf("could get executable's path %v", in.err)
	}
	if !(len(in.To) == 1 || len(in.To) == 2) {
		return fmt.Errorf("'to' must must look like 'a b' or just 'b'")
	}

	fmt.Printf("replacing ...\n")
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
		count := 0
		if len(in.To) == 1 {
			count += bytes.Count(contents, []byte(in.FromName))
			contents = bytes.ReplaceAll(contents, []byte(in.FromName), []byte(in.To[0]))
		}
		if len(in.To) == 2 {
			count += bytes.Count(contents, []byte(in.FromOrg))
			contents = bytes.ReplaceAll(contents, []byte(in.FromOrg), []byte(in.To[0]))
			count += bytes.Count(contents, []byte(in.FromName))
			contents = bytes.ReplaceAll(contents, []byte(in.FromName), []byte(in.To[1]))
		}
		fmt.Printf("  replaced %d occurrences in %v\n", count, path)
		return ioutil.WriteFile(path, contents, info.Mode())
	})
	if err != nil {
		return fmt.Errorf("error %v\n", err)
	}
	fmt.Printf("done\n")
	return nil
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

package md_tmpl

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jpillora/md-tmpl/mdtmpl"
	"github.com/wxio/acli/internal/types"
)

type mdTmplOpt struct {
	rt         *types.Root
	Filename   string `opts:"mode=arg"`
	WorkingDir string
	Preview    bool
}

func NewMdTmpl(rt *types.Root) interface{} {
	in := mdTmplOpt{
		WorkingDir: ".",
		rt:         rt,
	}
	return &in
}

func (gen *mdTmplOpt) Run() error {
	gen.rt.Config(gen)
	fp := filepath.Join(gen.WorkingDir, gen.Filename)
	if b, err := ioutil.ReadFile(fp); err != nil {
		return err
	} else {
		if gen.Preview {
			for i, c := range mdtmpl.Commands(b) {
				fmt.Printf("%18s#%d %s\n", gen.Filename, i+1, c)
			}
			return nil
		}
		b = mdtmpl.ExecuteIn(b, filepath.Join(gen.WorkingDir))
		if err := ioutil.WriteFile(fp, b, 0655); err != nil {
			return err
		}
		log.Printf("executed templates and rewrote '%s'", gen.Filename)
		return nil
	}
}

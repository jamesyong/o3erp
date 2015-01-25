package templating

import (
	"github.com/jamesyong/o3erp/go/config"
	"github.com/unrolled/render"
)

var (
	Render  *render.Render
	Options = render.Options{}
)

func Setup() {
	if Options.Directory == "" {
		Options.Directory = config.PATH_BASE_GOLANG_TEMPLATES
	}
	Render = render.New(Options)
}

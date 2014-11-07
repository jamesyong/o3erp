package templating

import (
	"github.com/jamesyong/o3erp/go/config"
	"github.com/unrolled/render"
)

var (
	Render *render.Render
)

func Setup() {
	Render = render.New(render.Options{Directory: config.PATH_BASE_GOLANG_TEMPLATES})
}

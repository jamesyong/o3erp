package templating

import (
	"github.com/jamesyong/o3erp/frontend/config"
	"github.com/unrolled/render"
)

var (
	Render *render.Render
)

func Setup() {
	Render = render.New(render.Options{Directory: config.PATH_BASE_FRONTEND_TEMPLATES})
}

package templating

import (
	"github.com/unrolled/render"
)

var (
	Render = render.New(render.Options{Directory: "frontend/templates"})
)

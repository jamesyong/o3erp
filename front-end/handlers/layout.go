package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"o3erp/front-end/config"
	"o3erp/front-end/templating"
)

func LayoutViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mapDetail := make(map[string]interface{})
	mapDetail["WebSocketHost"] = config.HOST

	templating.Render.HTML(w, http.StatusOK, "layout", mapDetail)
}

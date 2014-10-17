package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"o3erp/frontend/config"
	"o3erp/frontend/templating"
)

func LayoutViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mapDetail := make(map[string]interface{})
	mapDetail["WebSocketHost"] = config.HOST

	templating.Render.HTML(w, http.StatusOK, "layout", mapDetail)
}

package handlers

import (
	"github.com/jamesyong/o3erp/frontend/config"
	"github.com/jamesyong/o3erp/frontend/templating"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LayoutViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mapDetail := make(map[string]interface{})
	mapDetail["WebSocketHost"] = config.HOST

	templating.Render.HTML(w, http.StatusOK, "layout", mapDetail)
}

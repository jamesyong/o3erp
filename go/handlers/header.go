package handlers

import (
	"github.com/jamesyong/o3erp/go/templating"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mapDetail := make(map[string]interface{})
	templating.Render.HTML(w, http.StatusOK, "header", mapDetail)
}

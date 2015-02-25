package h_accounting

import (
	"github.com/jamesyong/o3erp/go/sessions"
	"github.com/jamesyong/o3erp/go/templating"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func AgreementViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mapDetail := make(map[string]interface{})

	session, _ := sessions.SessionStore.Get(r, "session-name")
	userLoginId := session.Values[sessions.USER_LOGIN_ID]
	mapDetail["userLoginId"] = userLoginId

	/*
		serviceContext := make(map[string]string)

		labelMap, err := helper.RunThriftService(helper.GetCallOfbizServiceFunction(userLoginId.(string), "", serviceContext))
		if err != nil {
			log.Println("error: ", err)
		}
		log.Println(labelMap)*/

	log.Println("Now at AgreementViewHandler")

	templating.Render.HTML(w, http.StatusOK, "acctg_agreement", mapDetail)
}

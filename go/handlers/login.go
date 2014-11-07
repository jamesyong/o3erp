package handlers

import (
	"fmt"
	"github.com/jamesyong/o3erp/go/config"
	"github.com/jamesyong/o3erp/go/helper"
	"github.com/jamesyong/o3erp/go/sessions"
	"github.com/jamesyong/o3erp/go/templating"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func LoginViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// no point going to login page if one is already logged in
	session, _ := sessions.SessionStore.Get(r, "session-name")
	if session.Values["partyId"] != nil {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	mapDetail := make(map[string]interface{})
	mapDetail["WebSocketHost"] = config.HOST

	templating.Render.HTML(w, http.StatusOK, "login", mapDetail)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	userLoginId := r.FormValue("email")
	password := r.FormValue("password")
	if userLoginId != "" && password != "" {

		msg, err := helper.RunThriftService(helper.GetLoginFunction(userLoginId, password))
		if err != nil {
			log.Println("error: ", err)
		} else {
			log.Println("success: ", msg, err)
		}
		if msg["responseMessage"] == "success" {
			// Get a session. We're ignoring the error resulted from decoding an
			// existing session: Get() always returns a session, even if empty.
			session, _ := sessions.SessionStore.Get(r, "session-name")

			// Set some session values.
			log.Println("partyId:", msg["partyId"])
			session.Values[sessions.PARTY_ID] = msg["partyId"]
			session.Values[sessions.USER_LOGIN_ID] = userLoginId
			// Save it.
			session.Save(r, w)

			// redirect to dashboard
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

	}

	fmt.Println("login credential incorrect")
	// redirect to login
	http.Redirect(w, r, "/login", http.StatusFound)
	return

}

func LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, _ := sessions.SessionStore.Get(r, "session-name")

	session.Values = nil

	// Save it.
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
	return

}

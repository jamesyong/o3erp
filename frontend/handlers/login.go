package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"o3erp/frontend/config"
	"o3erp/frontend/helper"
	"o3erp/frontend/sessions"
	"o3erp/frontend/templating"
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

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email != "" && password != "" {

		msg, err := helper.Login("admin", "ofbiz")
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

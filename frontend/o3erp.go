package frontend

import (
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"o3erp/frontend/config"
	"o3erp/frontend/handlers"
	"o3erp/frontend/sessions"
)

/* a middleware that redirects unauthenticated requests to login page
 */
func authenticationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path[1:]
		// log.Println("path:" + path)
		if path != "login" && path != "favicon.ico" {
			session, _ := sessions.SessionStore.Get(r, "session-name")
			if session.Values["partyId"] == nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// simple handler to redirect / to home view i.e. dashboard
func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func Startup() {
	mux := httprouter.New()
	mux.GET("/", HomeHandler)
	mux.GET("/login", handlers.LoginViewHandler)
	mux.POST("/login", handlers.LoginHandler)
	mux.GET("/logout", handlers.LogoutHandler)
	mux.GET("/dashboard", handlers.LayoutViewHandler)

	// handle menu
	mux.GET("/menu", handlers.MenuHandler)

	// handle static files
	mux.ServeFiles("/assets/*filepath", http.Dir("../o3erp/frontend/assets"))

	var u *url.URL
	var err error
	u, err = url.Parse(config.BACK_END_URL)
	if err != nil {
		log.Fatal(err)
	}
	// handle reserve proxy
	proxy := handler(httputil.NewSingleHostReverseProxy(u))
	mux.HandlerFunc("GET", "/ap/*any", proxy)
	mux.HandlerFunc("POST", "/ap/*any", proxy)
	mux.HandlerFunc("GET", "/ar/*any", proxy)
	mux.HandlerFunc("POST", "/ar/*any", proxy)
	mux.HandlerFunc("GET", "/accounting/*any", proxy)
	mux.HandlerFunc("POST", "/accounting/*any", proxy)
	mux.HandlerFunc("GET", "/catalog/*any", proxy)
	mux.HandlerFunc("POST", "/catalog/*any", proxy)
	mux.HandlerFunc("GET", "/content/*any", proxy)
	mux.HandlerFunc("POST", "/content/*any", proxy)
	mux.HandlerFunc("GET", "/facility/*any", proxy)
	mux.HandlerFunc("POST", "/facility/*any", proxy)
	mux.HandlerFunc("GET", "/humanres/*any", proxy)
	mux.HandlerFunc("POST", "/humanres/*any", proxy)
	mux.HandlerFunc("GET", "/manufacturing/*any", proxy)
	mux.HandlerFunc("POST", "/manufacturing/*any", proxy)
	mux.HandlerFunc("GET", "/marketing/*any", proxy)
	mux.HandlerFunc("POST", "/marketing/*any", proxy)
	mux.HandlerFunc("GET", "/ordermgr/*any", proxy)
	mux.HandlerFunc("POST", "/ordermgr/*any", proxy)
	mux.HandlerFunc("GET", "/partymgr/*any", proxy)
	mux.HandlerFunc("POST", "/partymgr/*any", proxy)
	mux.HandlerFunc("GET", "/sfa/*any", proxy)
	mux.HandlerFunc("POST", "/sfa/*any", proxy)
	mux.HandlerFunc("GET", "/workeffort/*any", proxy)
	mux.HandlerFunc("POST", "/workeffort/*any", proxy)
	mux.HandlerFunc("GET", "/bi/*any", proxy)
	mux.HandlerFunc("POST", "/bi/*any", proxy)
	mux.HandlerFunc("GET", "/webtools/*any", proxy)
	mux.HandlerFunc("POST", "/webtools/*any", proxy)

	mux.HandlerFunc("GET", "/tomahawk/*any", proxy)
	mux.HandlerFunc("GET", "/flatgrey/*any", proxy)
	mux.HandlerFunc("GET", "/images/*any", proxy)

	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		SSLHost:     config.HOST,
	})
	app := authenticationHandler(secureMiddleware.Handler(mux))

	log.Println("Starting...")

	// HTTP
	go func() {
		log.Fatal(http.ListenAndServe(config.PORT_HTTP, mux))
	}()

	// HTTPS
	log.Fatal(http.ListenAndServeTLS(config.PORT_HTTPS, "../o3erp/frontend/cert.pem", "../o3erp/frontend/key.pem", app))

}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		r.Header.Set("X-Forwarded-Proto", "https")

		session, _ := sessions.SessionStore.Get(r, "session-name")
		partyId := session.Values["partyId"]
		if partyId != nil {

			r.Header.Set("REMOTE_USER", partyId.(string))
		}
		p.ServeHTTP(w, r)

		newLoc := w.Header().Get("Location")
		if newLoc != "" {
			log.Println("Redirecting.. " + newLoc)
		}
	}
}

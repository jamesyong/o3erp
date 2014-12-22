package server

import (
	"github.com/jamesyong/o3erp/go/config"
	"github.com/jamesyong/o3erp/go/handlers"
	"github.com/jamesyong/o3erp/go/sessions"
	"github.com/jamesyong/o3erp/go/templating"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
	http.Redirect(w, r, "/main", http.StatusFound)
}

func getProxyPathListForGet() []string {
	return []string{
		"/ap/*any",
		"/ar/*any",
		"/accounting/*any",
		"/catalog/*any",
		"/content/*any",
		"/facility/*any",
		"/humanres/*any",
		"/manufacturing/*any",
		"/marketing/*any",
		"/ordermgr/*any",
		"/partymgr/*any",
		"/sfa/*any",
		"/workeffort/*any",
		"/bi/*any",
		"/webtools/*any",
		"/tomahawk/*any",
		"/flatgrey/*any",
		"/images/*any"}
}

func getProxyPathListForPost() []string {
	return []string{
		"/ap/*any",
		"/ar/*any",
		"/accounting/*any",
		"/catalog/*any",
		"/content/*any",
		"/facility/*any",
		"/humanres/*any",
		"/manufacturing/*any",
		"/marketing/*any",
		"/ordermgr/*any",
		"/partymgr/*any",
		"/sfa/*any",
		"/workeffort/*any",
		"/bi/*any",
		"/webtools/*any"}
}

func getPathMapForGet() map[string]httprouter.Handle {
	m := make(map[string]httprouter.Handle)
	m["/"] = HomeHandler
	m["/login"] = handlers.LoginViewHandler
	m["/logout"] = handlers.LogoutHandler
	m["/dashboard"] = handlers.DashboardViewHandler
	m["/main"] = handlers.LayoutViewHandler
	m["/menu"] = handlers.MenuHandler
	m["/header"] = handlers.HeaderHandler
	return m
}

func getPathMapForPost() map[string]httprouter.Handle {
	m := make(map[string]httprouter.Handle)
	m["/login"] = handlers.LoginHandler
	return m
}

var ProxyPathListForGet []string = getProxyPathListForGet()
var ProxyPathListForPost []string = getProxyPathListForPost()
var PathMapForGet map[string]httprouter.Handle = getPathMapForGet()
var PathMapForPost map[string]httprouter.Handle = getPathMapForPost()

func Startup() {

	templating.Setup()

	mux := httprouter.New()

	for path := range PathMapForGet {
		mux.GET(path, PathMapForGet[path])
	}
	for path := range PathMapForPost {
		mux.POST(path, PathMapForPost[path])
	}

	// handle static files
	mux.ServeFiles("/assets/*filepath", http.Dir(config.PATH_BASE_GOLANG_ASSETS))

	var u *url.URL
	var err error
	u, err = url.Parse(config.BACK_END_URL)
	if err != nil {
		log.Fatal(err)
	}
	// handle reserve proxy
	proxy := handler(httputil.NewSingleHostReverseProxy(u))

	// set GET paths
	for _, path := range ProxyPathListForGet {
		mux.HandlerFunc("GET", path, proxy)
	}

	// set POST paths
	for _, path := range ProxyPathListForPost {
		mux.HandlerFunc("POST", path, proxy)
	}

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

	log.Fatal(http.ListenAndServeTLS(config.PORT_HTTPS, config.PATH_BASE_GOLANG_CERT+"/cert.pem", config.PATH_BASE_GOLANG_CERT+"/key.pem", app))

}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		r.Header.Set("X-Forwarded-Proto", "https")

		session, _ := sessions.SessionStore.Get(r, "session-name")
		userLoginId := session.Values[sessions.USER_LOGIN_ID]
		if userLoginId != nil {
			r.Header.Set("REMOTE_USER", userLoginId.(string))
		}
		p.ServeHTTP(w, r)

		newLoc := w.Header().Get("Location")
		if newLoc != "" {
			log.Println("Redirecting.. " + newLoc)
		}
	}
}

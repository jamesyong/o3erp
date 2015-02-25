package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var (
	/* for validating the *any part of the url */
	validPath  = regexp.MustCompile("^/([a-zA-Z0-9]+)$")
	validPath2 = regexp.MustCompile("^/([a-zA-Z0-9]+)/([a-zA-Z0-9]+)$")
)

func getRangeFromReq(r *http.Request) (int, int) {

	from := 0
	to := 100
	dojoRange := r.Header.Get("Range")
	if dojoRange != "" {
		r := regexp.MustCompile("^items=([0-9]+)-([0-9]+)/([0-9]+)$").FindStringSubmatch(dojoRange)
		if len(r) == 4 {
			from, _ = strconv.Atoi(r[1])
			to, _ = strconv.Atoi(r[2])
			log.Println(r[1], r[2])
		}
	}
	return from, to
}

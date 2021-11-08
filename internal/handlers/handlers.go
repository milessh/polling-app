package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/milessh/polling-app/pkg/router"
)

func SetupHandlers() {
	router.NewRoute(router.HTTP_POST, "/create", createPoll)
	router.NewRoute(router.HTTP_GET, "/poll/([^/]+)", viewPoll)
	router.NewRoute(router.HTTP_POST, "/poll/([^/]+)/tally", endPoll)
	router.NewRoute(router.HTTP_GET, "(.*?)", defaultFallThrough)
	router.NewRoute(router.HTTP_POST, "(.*?)", defaultFallThrough)
}

func extractUserHash(r *http.Request) string {
	// assuming will always be populated as a single value by the reverse proxy
	clientIp := r.Header.Get("X-Forwarded-For")
	ua := r.Header.Get("User-Agent")

	hash := sha1.Sum([]byte(clientIp + ua))

	return hex.EncodeToString(hash[:])
}

func createPoll(w http.ResponseWriter, r *http.Request) {

	log.Println(r)

	switch method := r.Method; method {
	case router.HTTP_POST:
		log.Println("post got")
	default:
		log.Println("unknown method, ", method)

	}
}

func viewPoll(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}

func endPoll(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

	fields := router.GetFields(r)

	log.Println(fields)
}

func defaultFallThrough(w http.ResponseWriter, r *http.Request) {
	log.Println("url: ", r.URL)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(500)

	fmt.Fprintln(w, "Unsupported Content-Type")

}

package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
)

const HTTP_GET string = "GET"
const HTTP_POST string = "POST"

var routes []route

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type reqCtx struct{}

func NewRoute(method string, pattern string, handler http.HandlerFunc) {
	routes = append(routes, route{method, regexp.MustCompile("^" + pattern + "$"), handler})
}

func Serve(w http.ResponseWriter, r *http.Request) {
	log.Println("serving incoming request: ", r)

	for _, elem := range routes {

		matches := elem.regex.FindStringSubmatch(r.URL.Path)

		log.Println("see matches: ", matches)

		if len(matches) > 0 {

			if r.Method != elem.method {
				log.Println("missed method: ", r.Method)
				continue
			}

			log.Println("matches: ", matches[1:])

			// save any capture groups into the context for use in the handlers
			ctx := context.WithValue(r.Context(), reqCtx{}, matches[1:])

			elem.handler(w, r.WithContext(ctx))
			return
		}
	}

	// TODO: handle fall through case where we do not match anything
}

func GetFields(r *http.Request) []string {
	return r.Context().Value(reqCtx{}).([]string)
}

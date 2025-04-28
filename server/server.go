package server

import (
	"context"
	"net/http"
	"strings"
	router "vmwar/server/router"
)

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range router.Get_routes() {
		matches := route.Get_route_regex().FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.Get_route_method() {
				allow = append(allow, route.Get_route_method())
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.Get_route_handler()(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

package main

import "net/http"

func getUrlParam(r *http.Request, index int) string {
	fields := r.Context().Value(UrlParamsContextKey{}).([]string)
	return fields[index]
}

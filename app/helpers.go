package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (s *Server) ErrInternalServer(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	s.errorLog.Output(2, trace)
	status := http.StatusInternalServerError

	if s.debug {
		http.Error(w, trace, status)
		return
	}

	http.Error(w, http.StatusText(status), status)
}

func (s *Server) ErrBadRequest(w http.ResponseWriter, err error) {
	s.errorLog.Print(err.Error())
	status := http.StatusBadRequest
	http.Error(w, http.StatusText(status), status)
}

func (s *Server) ErrNotFound(w http.ResponseWriter) {
	status := http.StatusNotFound
	http.Error(w, http.StatusText(status), status)
}

func getUrlParam(r *http.Request, index int) string {
	fields := r.Context().Value(UrlParamsContextKey{}).([]string)
	return fields[index]
}

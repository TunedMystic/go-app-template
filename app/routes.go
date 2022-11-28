package main

import "net/http"

// Handler prepares and returns the Server's main http.Handler.
// .
func (s *Server) Handler() http.Handler {
	s.Route("GET", "/", s.HandleHome())
	s.Route("GET", "/health", s.HandleHealth())
	s.Route("GET", "/about", s.HandleAbout())
	s.Route("GET", "/notes", s.HandleNotes())
	s.Route("GET", "/notes/([0-9]+)", s.HandleNote())

	return s.Logger(s)
}

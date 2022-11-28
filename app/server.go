package main

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Server is the main application server.
// .
type Server struct {
	debug    bool
	infoLog  *log.Logger
	errorLog *log.Logger
	routes   map[string]Route
	users    UserModel
	notes    NoteModel
}

// UserModel defines the database functionality to manage Users.
// .
type UserModel interface {
	ActiveUsers() ([]*User, error)
	GetUser(string) (*User, error)
}

// NoteModel defines the database functionality to manage Notes.
// .
type NoteModel interface {
	AllNotes() ([]*Note, error)
}

// Route represents the mapping of an http method and pattern to a handler.
// .
type Route struct {
	method  string
	regex   *regexp.Regexp
	handler http.Handler
}

// UrlParamsContextKey is used to store url params in the request context.
// .
type UrlParamsContextKey struct{}

// Route maps the given method and patten to the http handler.
// .
func (s *Server) Route(method, pattern string, handler http.Handler) {
	if s.routes == nil {
		s.routes = make(map[string]Route)
	}

	routeKey := method + pattern

	// Add the route to the server if it does not exist.
	if _, ok := s.routes[routeKey]; !ok {

		s.infoLog.Printf("[setup-route] %s %s\n", method, pattern)

		s.routes[routeKey] = Route{
			method:  method,
			regex:   regexp.MustCompile("^" + pattern + "$"),
			handler: handler,
		}
	}
}

// ServeHTTP is the application's main http handler.
// .
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string

	// Check if the incoming url matches one of the routes.
	// If it matches, then serve the corresponding handler.
	// If it does not match, return a 404 / 405.
	for _, route := range s.routes {

		matches := route.regex.FindStringSubmatch(r.URL.Path)

		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}

			// The url and method both match, so we serve the handler.
			ctx := context.WithValue(r.Context(), UrlParamsContextKey{}, matches[1:])
			route.handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}

	// There was a url match, but the method was incorrect.
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// There was no url match.
	s.ErrNotFound(w)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>App home page!</h1>"))
	}
}

func (s *Server) HandleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}
}

func (s *Server) HandleAbout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>About</h1>"))
	}
}

func (s *Server) HandleNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notes, err := s.notes.AllNotes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(notes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func (s *Server) HandleNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID, err := strconv.Atoi(getUrlParam(r, 0))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		s.infoLog.Printf("Getting note id=%d", noteID)

		notes, err := s.notes.AllNotes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if noteID >= len(notes) {
			msg := fmt.Sprintf("note ID=%d not found", noteID)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		note := notes[noteID]

		content := fmt.Sprintf("note ID=%d is %+v", noteID, note)
		w.Write([]byte(content))
	}
}

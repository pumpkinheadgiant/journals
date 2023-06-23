package main

import (
	"github.com/bmizerany/pat"
	"journals/journalServer"
	"net/http"
)

func main() {
	server := pat.New()

	server.Get("/journals", authStub(journalServer.GetAllJournalsHandler))
	server.Put("/journal", authStub(journalServer.PostJournalHandler))
	http.Handle("/", server)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("server error: " + err.Error())
	}
}

func authStub(handlerFunc func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// perform auth here
		// if auth:
		handlerFunc(w, r)
		// else
		// http.Error(.....)
	}
}

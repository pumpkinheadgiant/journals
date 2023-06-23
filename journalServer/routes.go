package journalServer

import (
	"encoding/json"
	"net/http"
)

func GetAllJournalsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	journals, err := journalManager.GetAllJournals()
	if err != nil {
		// handle error
	}
	journalsJSON, err := json.Marshal(journals)
	if err != nil {
		// handle error
	}
	w.Write(journalsJSON)
}

func PostJournalHandler(w http.ResponseWriter, req *http.Request) {

}

func GetJournalHandler(w http.ResponseWriter, req *http.Request) {

}

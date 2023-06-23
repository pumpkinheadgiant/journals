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
	w.Header().Set("Content-Type", "application/json")
	newJournal := Journal{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newJournal)
	if err != nil {
		// handle error
	} else {
		if journalManager.DoesJournalExist(newJournal.Name) {
			// handle error
			w.WriteHeader(409)
		} else {
			err = journalManager.AddNewJournal(newJournal)
			if err != nil {
				// handle error
			}
			w.WriteHeader(201)
		}
	}
}

func GetJournalHandler(w http.ResponseWriter, req *http.Request) {

}

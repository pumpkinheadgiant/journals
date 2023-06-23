package journalServer

import "fmt"

var journalManager JournalManager

func init() {
	journalManager = JournalManager{}
}

type JournalManager struct {
	journals map[string]Journal
}

type Journal struct {
	name string
}

type JournalEntry struct {
	date  string
	entry string
}

func (j *JournalManager) GetAllJournals() ([]Journal, error) {
	journals := []Journal{}

	for _, journal := range j.journals {
		journals = append(journals, journal)
	}
	return journals, nil
}

func (j *JournalManager) GetJournalByName(name string) (Journal, error) {
	journal, found := j.journals[name]
	if !found {
		return Journal{}, fmt.Errorf("unknown journal: %v", name)
	}
	return journal, nil
}

func (j *JournalManager) AddNewJournal(name string) error {
	_, exists := j.journals[name]
	if exists {
		return fmt.Errorf("journal with name '%v' already exists", name)
	} else {
		j.journals[name] = Journal{name: name}
	}
	return nil
}

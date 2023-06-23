package journalServer

import "fmt"

var journalManager JournalManager

var initialJournal = Journal{
	Name:    "initial journal",
	Entries: []JournalEntry{"Welcome to my journal!", "It's been a while since my first post."},
}

func init() {
	journalManager = JournalManager{}
	journalManager.journals = map[string]Journal{}
	journalManager.journals[initialJournal.Name] = initialJournal
}

type JournalManager struct {
	journals map[string]Journal
}

type Journal struct {
	Name    string
	Entries []JournalEntry
}

type JournalEntry string

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
		return fmt.Errorf("journal with Name '%v' already exists", name)
	} else {
		j.journals[name] = Journal{Name: name}
	}
	return nil
}

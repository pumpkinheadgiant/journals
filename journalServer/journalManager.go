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

func (j *JournalManager) DoesJournalExist(name string) bool {
	_, exists := j.journals[name]
	return exists
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

func (j *JournalManager) AddNewJournal(journal Journal) error {
	_, exists := j.journals[journal.Name]
	if exists {
		return fmt.Errorf("journal with Name '%v' already exists", journal.Name)
	} else {
		// more validation here....
		j.journals[journal.Name] = journal
	}
	return nil
}

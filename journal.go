package imdb

type Journal struct {
	head    *Log
	current *Log
	pool    *LogPool
}

// NewJournal initilize the journal module
func NewJournal() {
}

func (imdb *Journal) SetValue(key string, value string) {
}

func (imdb *Journal) Begin() {
}

func (imdb *Journal) Commit() {
}

func (imdb *Journal) Rollback() {
}

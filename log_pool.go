package imdb

type Log struct {
	key       string
	old       string
	new       string
	committed bool
	next      *Log
	prev      *Log
}

type LogPool struct {
	head *Log
}

// NewLogPool initialize a Log pool
func NewLogPool(size int) (*LogPool, error) {
}

func (p *LogPool) GetNewLog() (*Log, error) {
}

func (p *LogPool) FreeLogs(h *Log) error {
}

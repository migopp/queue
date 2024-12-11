package queue

type Entry struct {
	ID     string
	Length Length
}

func (e *Entry) CreateEntry(id string, length Length) *Entry {
	return &Entry{
		ID:     id,
		Length: length,
	}
}

package storage

type storage struct {
	urls map[int]string
}

type Storage interface {
	GetURL(id int) (string, bool)
	AddURL(url string)
	CurrentIndex() int
}

func (s *storage) CurrentIndex() int {
	return len(s.urls) - 1
}

func (s *storage) GetURL(id int) (string, bool) {
	value, ok := s.urls[id]
	return value, ok
}

func (s *storage) AddURL(url string) {
	if s.urls == nil {
		s.urls = map[int]string{}
	}
	s.urls[len(s.urls)] = url
}

func New() *storage {
	return &storage{}
}

package storage

type storage struct {
	urls map[int]string
}

type Storage interface {
	GetUrl(id int) (string, bool)
	AddUrl(url string)
	CurrentIndex() int
}

func (s *storage) CurrentIndex() int {
	return len(s.urls) - 1
}

func (s *storage) GetUrl(id int) (string, bool) {
	value, ok := s.urls[id]
	return value, ok
}

func (s *storage) AddUrl(url string) {
	if s.urls == nil {
		s.urls = map[int]string{}
	}
	s.urls[len(s.urls)] = url
}

func New() *storage {
	return &storage{}
}

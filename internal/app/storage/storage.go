package storage

type storage struct {
	urls         map[int]string
	currentIndex int
}

type Storage interface {
	GetUrl(id int) (string, bool)
	AddUrl(url string)
	CurrentIndex() int
}

func (s *storage) CurrentIndex() int {
	return s.currentIndex - 1
}

func (s *storage) GetUrl(id int) (string, bool) {
	value, ok := s.urls[id]
	return value, ok
}

func (s *storage) AddUrl(url string) {
	if s.urls == nil {
		s.urls = map[int]string{
			0: url,
		}
	}
	s.currentIndex += 1
	s.urls[len(s.urls)] = url
}

func New() *storage {
	return &storage{}
}

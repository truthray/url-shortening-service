package storage

type Storage struct {
	urls         map[int]string
	currentIndex int
}

type IStorage interface {
	GetUrl(id int) string
	AddUrl(url string)
	CurrentIndex() int
}

func (s *Storage) CurrentIndex() int {
	return s.currentIndex - 1
}

func (s *Storage) GetUrl(id int) (string, bool) {
	value, ok := s.urls[id]
	return value, ok
}

func (s *Storage) AddUrl(url string) {
	if s.urls == nil {
		s.urls = map[int]string{
			0: url,
		}
	}
	s.currentIndex += 1
	s.urls[len(s.urls)] = url
}

func New() *Storage {
	return &Storage{}
}

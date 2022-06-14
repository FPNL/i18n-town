package anime

type Series struct {
	Name     string
	release  bool
	episodes []*Episode
}

func (s Series) Status() (status struct {
	status  []Status
	average int
}) {
	panic("Implement Me!")
	return
}

func (s *Series) AddEpisode(ep []*Episode) {
	s.episodes = ep
}

func (s *Series) isOverBound(n int) bool {
	return n < 0 || n > (len(s.episodes)-1)
}

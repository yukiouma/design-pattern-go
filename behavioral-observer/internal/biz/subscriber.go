package biz

type Subscriber interface {
	Receive(content string)
	EventHistory() []string
	CleanCache()
}

type userClient struct {
	id      string
	history []string
}

func NewUserClient(user string) Subscriber {
	return &userClient{
		id:      user,
		history: make([]string, 0),
	}
}

func (s *userClient) Receive(content string) {
	s.history = append(s.history, content)
}

func (s *userClient) EventHistory() []string {
	return s.history
}

func (s *userClient) CleanCache() {
	s.history = make([]string, 0)
}

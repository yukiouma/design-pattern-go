package cor

type State struct {
	data map[string]Role
}

func NewState() *State {
	return &State{data: make(map[string]Role)}
}

func (s *State) Register(username string, role Role) {
	s.data[username] = role
}

func (s *State) Findout(username string, role Role) bool {
	r, ok := s.data[username]
	if !ok {
		return false
	}
	return r == role
}

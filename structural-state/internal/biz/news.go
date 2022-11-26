package biz

type News struct {
	content string
	state   State
}

func NewNews(content string) *News {
	news := &News{
		content: content,
	}
	news.changeState(newDraftState(news))
	return news
}

func (n *News) changeState(state State) {
	n.state = state
}

func (n *News) changeContent(content string) {
	n.content = content
}

func (n *News) Publish() {
	n.state.Publish()
}

func (n *News) Update(content string) error {
	return n.state.Update(content)
}

func (n *News) Revert() {
	n.state.Revert()
}

func (n *News) State() stateType {
	switch n.state.(type) {
	case *draftState:
		return StateDraft
	case *moderationState:
		return StateModeration
	case *publishedState:
		return StatePublished
	default:
		return -1
	}
}

package biz

type moderationState struct {
	news *News
}

func newModerationState(news *News) State {
	return &moderationState{
		news: news,
	}
}

func (s *moderationState) Publish() {
	s.news.changeState(newPubishedState(s.news))
}

func (s *moderationState) Update(content string) error {
	return ErrUpdateInModeration
}

func (s *moderationState) Revert() {
	s.news.changeState(newDraftState(s.news))
}

package biz

type draftState struct {
	news *News
}

func newDraftState(news *News) State {
	return &draftState{
		news: news,
	}
}

func (s *draftState) Publish() {
	s.news.changeState(newModerationState(s.news))
}

func (s *draftState) Update(content string) error {
	s.news.changeContent(content)
	return nil
}

func (s *draftState) Revert() {

}

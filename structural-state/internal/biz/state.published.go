package biz

type publishedState struct {
	news *News
}

func newPubishedState(news *News) State {
	return &publishedState{
		news: news,
	}
}

func (s *publishedState) Publish() {

}

func (s *publishedState) Update(content string) error {
	return ErrUpdateInPublished
}

func (s *publishedState) Revert() {
	s.news.changeState(newDraftState(s.news))
}

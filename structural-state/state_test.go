package structuralstate

import (
	"structural-state/internal/biz"
	"testing"
)

const (
	contentV1 = "content version 1"
	contentV2 = "content version 2"
)

func TestState(t *testing.T) {
	news := biz.NewNews(contentV1)
	if news.State() != biz.StateDraft {
		t.Fatalf("expect %v, got %v", biz.StateDraft, news.State())
	}
	news.Publish()
	if news.State() != biz.StateModeration {
		t.Fatalf("expect %v, got %v", biz.StateModeration, news.State())
	}
	news.Revert()
	if news.State() != biz.StateDraft {
		t.Fatalf("expect %v, got %v", biz.StateDraft, news.State())
	}
	if err := news.Update(contentV2); err != nil {
		t.Fatal("news can be edit in draft state")
	}
	news.Publish()
	if err := news.Update(contentV2); err == nil {
		t.Fatal("news cannot be edit in moderation state")
	}
	news.Publish()
	if news.State() != biz.StatePublished {
		t.Fatalf("expect %v, got %v", biz.StatePublished, news.State())
	}
	if err := news.Update(contentV2); err == nil {
		t.Fatal("news cannot be edit in published state")
	}
	news.Revert()
	if news.State() != biz.StateDraft {
		t.Fatalf("expect %v, got %v", biz.StateDraft, news.State())
	}
}

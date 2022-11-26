package structuralproxy

import (
	"structural-proxy/internal/proxy"
	"testing"
)

func TestProxy(t *testing.T) {
	hosts := []string{"youtube.com", "google.com", "youtube.com", "google.com"}
	proxy := proxy.NewVPN()
	for _, v := range hosts {
		_, err := proxy.Access(v)
		if err != nil {
			t.Fatalf("did not pass because: %v", err)
		}
	}
	_, err := proxy.Access("github.com")
	if err == nil {
		t.Fatalf("you can not access undefined address")
	}
}

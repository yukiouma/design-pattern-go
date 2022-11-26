package proxy

import (
	"io"
	"log"
	"os"
)

const rootdir = "/root/playground/golang/design-pattern/structural-proxy/data/"

type vpnServer struct {
	cache map[string]string
}

func newVpnServer() VPN {
	return &vpnServer{
		cache: make(map[string]string),
	}
}

func (s *vpnServer) Access(url string) (string, error) {
	if data, ok := s.cache[url]; ok {
		log.Printf("%s is cache, return directly", url)
		return data, nil
	}
	f, err := os.Open(rootdir + url)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	data := string(b)
	log.Printf("caching %s", url)
	s.cache[url] = data
	return data, nil
}

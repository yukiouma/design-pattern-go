package proxy

type vpnClient struct {
	server VPN
}

func (c *vpnClient) Access(url string) (string, error) {
	return c.server.Access(url)
}

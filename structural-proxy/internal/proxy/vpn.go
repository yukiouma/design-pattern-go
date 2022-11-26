package proxy

type VPN interface {
	Access(url string) (string, error)
}

func NewVPN() VPN {
	return &vpnClient{
		server: newVpnServer(),
	}
}

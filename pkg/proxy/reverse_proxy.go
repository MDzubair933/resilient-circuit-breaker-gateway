package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxyGateway struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewReverseProxyGateway(targetURL string) (*ReverseProxyGateway, error) {
	u, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}
	return &ReverseProxyGateway{
		target: u,
		proxy:  httputil.NewSingleHostReverseProxy(u),
	}, nil
}

func (g *ReverseProxyGateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.proxy.ServeHTTP(w, r)
}

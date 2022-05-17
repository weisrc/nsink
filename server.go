package nsink

import "github.com/miekg/dns"

type Options struct {
	Ttl   uint32
	Proxy string
	Open  bool
}

type Server struct {
	Clients   map[string]Client
	Trees     map[string]Tree
	Options   Options
	DnsClient dns.Client
}

func NewServer(options Options) Server {
	s := Server{
		Options: options,
		Clients: make(map[string]Client),
		Trees:   make(map[string]Tree),
	}
	return s
}

func (s Server) Listen(addr string, network string) error {
	return dns.ListenAndServe(addr, network, s)
}

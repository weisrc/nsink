package nsink

import (
	"fmt"

	"github.com/miekg/dns"
)

func (s Server) Reply(question dns.Question, client Client) *dns.RR {
	if question.Qtype != dns.TypeA && question.Qtype != dns.TypeAAAA {
		return nil
	}

	for id := range client.Trees {
		node, ok := s.Trees[id]

		if !ok {
			delete(client.Trees, id)
			continue
		}
		ip := node.FindIP(question.Name)

		if ip == "0.0.0.0" {
			client.Trees[id]++
		} else {
			client.Total++
		}

		if ip != "" {
			response, err := dns.NewRR(fmt.Sprintf("%s A %s", question.Name, ip))
			response.Header().Ttl = s.Options.Ttl
			if err == nil {
				return &response
			}
		}
	}

	return nil
}

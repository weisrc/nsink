package nsink

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func (s Server) ServeDNS(writer dns.ResponseWriter, request *dns.Msg) {

	response := new(dns.Msg)
	response.SetReply(request)
	response.Compress = false

	ip := writer.RemoteAddr().(*net.UDPAddr).IP.String()

	client, ok := s.Clients[ip]
	if !ok && !s.Options.Open {
		return
	}

	if request.Opcode != dns.OpcodeQuery {
		writer.WriteMsg(response)
		return
	}

	questions := []dns.Question{}

	for _, question := range response.Question {
		if answer := s.Reply(question, client); answer != nil {
			response.Answer = append(response.Answer, *answer)
		} else {
			questions = append(questions, question)
		}
	}

	if len(questions) > 0 {
		proxyRequest := dns.Msg{}
		proxyRequest.Question = questions
		proxyResponse, _, err := s.DnsClient.Exchange(&proxyRequest, s.Options.Proxy)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if proxyResponse != nil {
			response.Answer = append(response.Answer, proxyResponse.Answer...)
		}
	}

	writer.WriteMsg(response)
}

package dns_resolver

import (
	"github.com/miekg/dns"
	"log"
	"time"
)

// DnsClientWrapper is created to have the ability to mock the `Lookup` function and enable unit testing
type DnsClientWrapper interface {
	Lookup(queryMsg *dns.Msg, nServer string) (time.Duration, string, []string, error)
}

type dnsClientDnsClientWrapper struct {
	dnsClient dns.Client
}

func (c dnsClientDnsClientWrapper) Lookup(query *dns.Msg, address string) (time.Duration, string, []string, error)  {
	log.Printf("Perform lookup against %s for %s \n", address, query)
	response, roundTripTime, err := c.dnsClient.Exchange(query, address)

	if err != nil { return roundTripTime, response.String(), nil, err }

	answer := response.Answer
	recordStrings := make([]string, len(answer))
	for i, rs := range answer{
		recordStrings[i] = rs.String()
	}
	log.Printf("Raw response: %s \n", response)
	return roundTripTime, response.String(), recordStrings, nil
}

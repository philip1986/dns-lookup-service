package dns_resolver

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const defaultDnsPort = "53"

type DnsResolver struct {
	DnsClient DnsClientWrapper
}

func NewDnsResolver() *DnsResolver {
	return &DnsResolver{
		DnsClient: dnsClientDnsClientWrapper{dns.Client{} },
	}
}

func (d *DnsResolver) Resolve(domain string, recordType RecordType, nServerSetting string) (*LookupResult, error) {
	nameServer := d.handleDefaultNServerSetting(nServerSetting)
	address := net.JoinHostPort(nameServer, defaultDnsPort)
	query := d.createQueryMsg(domain, recordType)
	roundTripTime, raw, recordString, err := d.DnsClient.Lookup(query, address)

	if err != nil { return nil, err }

	return d.transformToLookupResult(nameServer, raw, recordString, roundTripTime), nil
}

func (d *DnsResolver) handleDefaultNServerSetting(nServer string) string {
	if nServer != "default" {
		return nServer
	}
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	return config.Servers[0]
}

func (d *DnsResolver) transformToLookupResult(nServer string, raw string, recordStrings []string, rtt time.Duration) *LookupResult {
	var resourceRecords = make([]ResourceRecord, len(recordStrings))

	for i, rs := range recordStrings {
		resourceRecords[i] = d.transformToResourceRecord(rs)
	}

	return &LookupResult{
		NServer: nServer,
		RoundTripTime: rtt,
		ResourceRecords: resourceRecords,
		Raw: raw,
	}
}

func (d *DnsResolver) transformToResourceRecord(rr string) ResourceRecord {
	recordParts := strings.Split(rr, "\t")

	if len(recordParts) < 5 {
		log.Println(fmt.Errorf("record seems to be malformed: %s", rr))
		return ResourceRecord{
			Value: "Record parsing error",
		}
	}

	ttl, _ := strconv.Atoi(recordParts[1])
	return ResourceRecord{
		Domain: recordParts[0],
		TTL: ttl,
		RecordType: recordParts[3],
		Value: recordParts[4],
	}
}

func (d *DnsResolver) createQueryMsg(domain string, recordType RecordType) *dns.Msg {
	msg := new(dns.Msg)
	msg.RecursionAvailable = true
	msg.SetQuestion(dns.Fqdn(domain), recordType)
	return msg
}
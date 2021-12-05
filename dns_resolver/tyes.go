package dns_resolver

import "time"

type RecordType = uint16

// relevant record type mapping copied from github.com/miekg/dns to avoid exposing underlying implementation details to users of `dns_resolver` package
const (
	TypeA          uint16 = 1
	TypeNS         uint16 = 2
	TypeCNAME      uint16 = 5
	TypeSOA        uint16 = 6
	TypePTR        uint16 = 12
	TypeMX         uint16 = 15
	TypeTXT        uint16 = 16
	TypeAAAA       uint16 = 28
	TypeSRV        uint16 = 33
	TypeANY   	   uint16 = 255
)

type LookupResult struct {
	NServer 		string
	RoundTripTime 	time.Duration
	ResourceRecords []ResourceRecord
	Raw 			string
}

type ResourceRecord struct {
	Domain 		string
	TTL 		int // in sec
	RecordType 	string
	Value 		string
}
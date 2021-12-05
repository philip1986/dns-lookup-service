package test_dns_resolver

import (
	"dns-lookup-service/dns_resolver"
	"dns-lookup-service/dns_resolver/test/mocks"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net"
	"testing"
	"time"
)

var ctrl *gomock.Controller
var mockedDnsClientWrapper *mock_dns_resolver.MockDnsClientWrapper
var dnsResolver *dns_resolver.DnsResolver

func NewDnsResolver(mockClient dns_resolver.DnsClientWrapper) *dns_resolver.DnsResolver {
	return &dns_resolver.DnsResolver{
		DnsClient: mockClient,
	}
}

func (s *DnsResolverSuite) SetupTest()  {
	ctrl = gomock.NewController(s.T())
	mockedDnsClientWrapper = mock_dns_resolver.NewMockDnsClientWrapper(ctrl)
	dnsResolver = NewDnsResolver(mockedDnsClientWrapper)
}

func (s *DnsResolverSuite) TearDownTest()  {
	ctrl.Finish()
}

func (s *DnsResolverSuite) TestItForSingleRecordFound() {
	mockedDnsClientWrapper.
		EXPECT().
		Lookup(gomock.Any(), gomock.Eq(net.JoinHostPort("8.8.8.8", "53"))).
		Return(time.Duration(1), "raw response", []string{"tesla.com.\t300\tIN\tA\t199.66.11.62"}, nil)

	r, _ := dnsResolver.Resolve("tesla.com", dns_resolver.TypeA, "8.8.8.8")

	assert.Equal(s.T(), r.NServer, "8.8.8.8")
	assert.Equal(s.T(), r.Raw, "raw response")
	assert.Equal(s.T(), r.RoundTripTime, time.Duration(1))
	assert.Equal(s.T(), len(r.ResourceRecords), 1)
	assert.Equal(s.T(), r.ResourceRecords[0], dns_resolver.ResourceRecord{
		Domain: "tesla.com.",
		TTL: 300,
		RecordType: "A",
		Value: "199.66.11.62",
	})
}

func (s *DnsResolverSuite) TestItForMultipleRecordsFound() {
	mockedDnsClientWrapper.
		EXPECT().
		Lookup(gomock.Any(), gomock.Eq(net.JoinHostPort("8.8.8.8", "53"))).
		Return(
			time.Duration(1),
			"raw response",
			[]string{
				"tesla.com.\t300\tIN\tTXT\tbugcrowd-verification=40bd5dd89a6e4073ca9bc76feac3a47b",
				"tesla.com.\t300\tIN\tTXT\tlogmein-domain-confirmation=9zxwVn2buGWrLtU24J88",
				"tesla.com.\t300\tIN\tTXT\tadobe-idp-site-verification=321c026a-3a8c-4206-a1fa-391a59585c54",
			},
			nil,
		)

	r, _ := dnsResolver.Resolve("tesla.com", dns_resolver.TypeTXT, "8.8.8.8")

	assert.Equal(s.T(), r.NServer, "8.8.8.8")
	assert.Equal(s.T(), r.Raw, "raw response")
	assert.Equal(s.T(), r.RoundTripTime, time.Duration(1))
	assert.Equal(s.T(), len(r.ResourceRecords), 3)
	assert.Equal(s.T(), r.ResourceRecords[0], dns_resolver.ResourceRecord{
		Domain: "tesla.com.",
		TTL: 300,
		RecordType: "TXT",
		Value: "bugcrowd-verification=40bd5dd89a6e4073ca9bc76feac3a47b",
	})
	assert.Equal(s.T(), r.ResourceRecords[1], dns_resolver.ResourceRecord{
		Domain: "tesla.com.",
		TTL: 300,
		RecordType: "TXT",
		Value: "logmein-domain-confirmation=9zxwVn2buGWrLtU24J88",
	})
	assert.Equal(s.T(), r.ResourceRecords[2], dns_resolver.ResourceRecord{
		Domain: "tesla.com.",
		TTL: 300,
		RecordType: "TXT",
		Value: "adobe-idp-site-verification=321c026a-3a8c-4206-a1fa-391a59585c54",
	})
}

func (s *DnsResolverSuite) TestItForDefaultNServerSetting() {
	mockedDnsClientWrapper.
		EXPECT().
		Lookup(gomock.Any(), gomock.Any()).
		Return(time.Duration(1), "raw response", []string{"tesla.com.\t300\tIN\tA\t199.66.11.62"}, nil)

	r, _ := dnsResolver.Resolve("tesla.com", dns_resolver.TypeA, "default")

	assert.Equal(s.T(), r.NServer != "default", true)

}

func (s *DnsResolverSuite) TestItForNoRecordFound() {
	mockedDnsClientWrapper.
		EXPECT().
		Lookup(gomock.Any(), gomock.Eq(net.JoinHostPort("8.8.8.8", "53"))).
		Return(time.Duration(1), "raw response", make([]string, 0), nil)

	r, _ := dnsResolver.Resolve("tesla.com", dns_resolver.TypeA, "8.8.8.8")

	assert.Equal(s.T(), r.NServer, "8.8.8.8")
	assert.Equal(s.T(), r.Raw, "raw response")
	assert.Equal(s.T(), r.RoundTripTime, time.Duration(1))
	assert.Equal(s.T(), len(r.ResourceRecords), 0)

}

func (s *DnsResolverSuite) TestItForMalformedRecord() {
	mockedDnsClientWrapper.
		EXPECT().
		Lookup(gomock.Any(), gomock.Eq(net.JoinHostPort("8.8.8.8", "53"))).
		Return(time.Duration(1), "raw response", []string{"not a valid record"}, nil)

	r, _ := dnsResolver.Resolve("tesla.com", dns_resolver.TypeA, "8.8.8.8")

	assert.Equal(s.T(), r.NServer, "8.8.8.8")
	assert.Equal(s.T(), r.Raw, "raw response")
	assert.Equal(s.T(), r.RoundTripTime, time.Duration(1))
	assert.Equal(s.T(), len(r.ResourceRecords), 1)
	assert.Equal(s.T(), r.ResourceRecords[0], dns_resolver.ResourceRecord{
		TTL: 0,
		Value: "Record parsing error",
	})
}

// NOT A UNIT TEST
func (s *DnsResolverSuite) TestItForReal() {
	//s.T().Skip("test without mocking. mostly used during developing and setting up mocked tests")
	dnsResolver := dns_resolver.NewDnsResolver()

	r, _ := dnsResolver.Resolve("tesla.com", dns_resolver.TypeMX, "8.8.8.8")

	fmt.Println(r.Raw)

	assert.Equal(s.T(), r.NServer, "8.8.8.8")
	assert.Equal(s.T(), r.RoundTripTime > 0, true)
	assert.Equal(s.T(), len(r.ResourceRecords) > 0, true)
}

type DnsResolverSuite struct {
	suite.Suite
}

func TestDnsResolverSuite(t *testing.T) {
	suite.Run(t, new(DnsResolverSuite))
}
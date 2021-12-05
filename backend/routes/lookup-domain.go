package routes

import (
	"dns-lookup-service/dns_resolver"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HandleLookupDomain(dnsResolver *dns_resolver.DnsResolver ,context *gin.Context) {
	domainName := context.Param("domainName")
	recordType := context.Param("recordType")
	nServer := strings.Split(context.DefaultQuery("nserver", "default"), ",")

	lookupResults := make([]*dns_resolver.LookupResult, len(nServer))
	for i, s := range nServer {
		res, err := dnsResolver.Resolve(domainName, dns_resolver.RecordTypeNameToRecordType[recordType], s)

		if err != nil {
			context.Error(err)
			return
		}
		lookupResults[i] = res
	}

	context.JSON(http.StatusOK, lookupResults)
}
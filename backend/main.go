package main

import (
	"dns-lookup-service/dns_resolver"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {

	dnsResolver := dns_resolver.NewDnsResolver()

    router := gin.Default()

    router.Static("/ui", "./frontend/dist")
    router.GET("/", func(context *gin.Context) {
    	context.Redirect(301, "/ui")
	})

    v1 := router.Group("api/v1")
	v1.GET("/lookup/domain/:domainName/recordtype/:recordType", func(context *gin.Context) {
		domainName := context.Param("domainName")
		recordType := context.Param("recordType")
		nServer := strings.Split(context.DefaultQuery("nserver", "default"), ",")
		println("got it " + domainName + " " + recordType + " ")

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
	})

    router.Run("localhost:8080")
}


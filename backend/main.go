package main

import (
	"dns-lookup-service/dns_resolver"
	"dns-lookup-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	dnsResolver := dns_resolver.NewDnsResolver()

    router := gin.Default()
    v1 := router.Group("api/v1")
	v1.GET("/lookup/domain/:domainName/recordtype/:recordType", func(context *gin.Context) {
		routes.HandleLookupDomain(dnsResolver, context)
	})

    router.Run("0.0.0.0:8080")
}


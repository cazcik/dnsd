package handler

import (
	"log"
	"math"

	miekgdns "github.com/miekg/dns"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"github.com/projectdiscovery/retryabledns"
)

// fetch the DNSData for the specified domain
func GetDomain(domain string) (*retryabledns.DNSData, error) {
	dnsClient := setupClients()
	response, err := dnsClient.QueryMultiple(domain)
	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}

// configure dsnx client
func setupClients() (*dnsx.DNSX) {
	dnsClient, err := dnsx.New(dnsx.Options{
		BaseResolvers: dnsx.DefaultResolvers,
		MaxRetries: 3,
		QuestionTypes: []uint16{miekgdns.TypeA, miekgdns.TypeTXT, miekgdns.TypeNS, miekgdns.TypeMX},
		Trace: true,
		TraceMaxRecursion: math.MaxUint16,
		Hostsfile: true,
		OutputCDN:  true,
	})
	if err != nil {
		log.Fatal(err)
	}

	return dnsClient
}
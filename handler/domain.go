package handler

import (
	"encoding/json"
	"log"
	"math"

	"github.com/gofiber/fiber/v2"
	miekgdns "github.com/miekg/dns"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"github.com/projectdiscovery/retryabledns"
)

func GetDomain(c *fiber.Ctx) error {
	domain := c.Params("domain")
	dnsClient := SetupClients()

	response, err := getResponse(dnsClient, domain)
	if err != nil {
		log.Fatal(err)
	}

	resultsJsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	resultsJsonString := json.RawMessage(resultsJsonBytes)

	return c.JSON(resultsJsonString)
}

func SetupClients() (*dnsx.DNSX) {
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

func getResponse(client *dnsx.DNSX, domain string) (*retryabledns.DNSData, error) {
	response, err := client.QueryMultiple(domain)
	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}
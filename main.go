package main

import (
	"fmt"
	"log"
	"math"
	"os"

	miekgdns "github.com/miekg/dns"
	asnmap "github.com/projectdiscovery/asnmap/libs"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"github.com/projectdiscovery/retryabledns"
)

type ASNResult struct {
	ASN     int    `json:"asn"`
	Org     string `json:"org"`
	Country string `json:"country"`
}

func main() {
	domain := os.Args[1]

	dnsClient, err := dnsx.New(dnsx.Options{
		BaseResolvers:     dnsx.DefaultResolvers,
		MaxRetries:        5,
		QuestionTypes:     []uint16{miekgdns.TypeA, miekgdns.TypeAAAA, miekgdns.TypeCNAME, miekgdns.TypeTXT, miekgdns.TypeNS, miekgdns.TypeMX, miekgdns.TypeSOA, miekgdns.TypeSRV, miekgdns.TypePTR, miekgdns.TypeSPF, miekgdns.TypeCAA},
		TraceMaxRecursion: math.MaxUint16,
		Hostsfile:         true,
	})
	if err != nil {
		log.Fatal(err)
	}

	asnClient, err := asnmap.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	response, nil := getResponse(dnsClient, domain)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	host := getHost(response)
	ns := getNS(response)
	mx := getMX(response)
	txt := getTXT(response)
	a := getA(response)

	fmt.Printf("host: %s\n", host)
	fmt.Printf("nameservers: %s\n", ns)
	for _, i := range ns {
		ips, err := dnsClient.Lookup(i)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		for _, ip := range ips {
			asn := getASN(asnClient, ip)
			fmt.Printf("asn: %v\n", asn)
		}
	}
	fmt.Printf("mx: %s\n", mx)
	fmt.Printf("txt: %s\n", txt)
	fmt.Printf("a: %s\n", a)
	for _, ip := range a {
		asn := getASN(asnClient, ip)
		fmt.Printf("asn: %v\n", asn)
	}
}

func getResponse(client *dnsx.DNSX, domain string) (*retryabledns.DNSData, error) {
	response, err := client.QueryMultiple(domain)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	return response, nil
}

func getASN(client *asnmap.Client, ip string) *ASNResult {
	response, err := client.GetData(ip)
	if err != nil {
		log.Fatal(err)
	}

	results := new(ASNResult)
	results.ASN = response[0].ASN
	results.Org = response[0].Org
	results.Country = response[0].Country
	
	return results
}

func getHost(response *retryabledns.DNSData) string {
	return response.Host
}

func getNS(response *retryabledns.DNSData) []string {
	return response.NS
}

func getMX(response *retryabledns.DNSData) []string {
	return response.MX
}

func getTXT(response *retryabledns.DNSData) []string {
	return response.TXT
}

func getA(response *retryabledns.DNSData) []string {
	return response.A
}
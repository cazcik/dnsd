package handler

import (
	"log"
	"math"

	miekgdns "github.com/miekg/dns"
	asnmap "github.com/projectdiscovery/asnmap/libs"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"github.com/projectdiscovery/retryabledns"
)

type Host struct {
	Name string
	IP string
	ASNOrg string
	ASNCountry string
}

// fetch the DNSData for the specified domain
func GetDomain(domain string) (map[string]interface{}, error) {
	dnsClient, asnClient := setupClients()
	response, err := dnsClient.QueryMultiple(domain)
	if err != nil {
		log.Fatal(err)
	}

	nameservers := getNSRecords(response, dnsClient, asnClient)
	mx := getMXRecords(response, dnsClient, asnClient)
	txt := response.TXT
	hosts := getHostRecords(response, dnsClient, asnClient)

	r := map[string]interface{} {
		"nameservers": nameservers,
		"mx": mx,
		"txt": txt,
		"host": hosts,
	}

	return r, nil
}

func getNSRecords(r *retryabledns.DNSData, dnsClient *dnsx.DNSX, asnClient *asnmap.Client) []Host {

	var hosts []Host
	for _, a := range r.NS {
		var h Host

		ip, err := dnsClient.Lookup(a)
		if err != nil {
			log.Fatal(err)
		}

		h.Name = a
		h.IP = ip[0]
		h.ASNOrg = getAsnOrg(asnClient, ip[0])
		h.ASNCountry = getAsnCountry(asnClient, ip[0])

		hosts = append(hosts, h)
	}

	return hosts
}

func getMXRecords(r *retryabledns.DNSData, dnsClient *dnsx.DNSX, asnClient *asnmap.Client) []Host {
	var hosts []Host
	for _, a := range r.MX {
		var h Host

		ip, err := dnsClient.Lookup(a)
		if err != nil {
			log.Fatal(err)
		}

		h.Name = a
		h.IP = ip[0]
		h.ASNOrg = getAsnOrg(asnClient, ip[0])
		h.ASNCountry = getAsnCountry(asnClient, ip[0])

		hosts = append(hosts, h)
	}

	return hosts
}

func getHostRecords(r *retryabledns.DNSData, dnsClient *dnsx.DNSX, asnClient *asnmap.Client) []Host {
	var hosts []Host
	for _, a := range r.A {
		var h Host

		h.Name = r.Host
		h.IP = a
		h.ASNOrg = getAsnOrg(asnClient, a)
		h.ASNCountry = getAsnCountry(asnClient, a)

		hosts = append(hosts, h)
	}

	return hosts
}

func getAsnOrg (c *asnmap.Client, host string) string {
	asn, err := c.GetData(host)
	if err != nil {
		log.Fatal(err)
	}

	return asn[0].Org
}

func getAsnCountry (c *asnmap.Client, host string) string {
	asn, err := c.GetData(host)
	if err != nil {
		log.Fatal(err)
	}

	return asn[0].Country
}

// configure dsnx client
func setupClients() (*dnsx.DNSX, *asnmap.Client) {
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

	asnClient, err := asnmap.NewClient()
	if  err != nil {
		log.Fatal(err)
	}

	return dnsClient, asnClient
}
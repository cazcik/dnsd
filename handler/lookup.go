package handler

import (
	"log"
	"math"
	"regexp"

	"github.com/gofiber/fiber/v2"
	miekgdns "github.com/miekg/dns"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"github.com/projectdiscovery/retryabledns"
)

type Host struct {
	Name string `json:"name"`
	IP string `json:"ip"`
}

type LookupInput struct {
	Domain string `json:"domain"`
}

// function to validate domain
func isValidDomain(domain string) bool {
	pattern := `^(?i)([a-z0-9-]+\.){1,}([a-z]{2,})$`
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(domain)
}

// perform the query and return the json
func Lookup(c *fiber.Ctx) error {
	input := new(LookupInput)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid domain input", "data": nil, })
	}

	d := input.Domain
 
	if (!isValidDomain(d)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid domain name", "data": nil, })
	}

	response, err := GetDomain(d)
	if (err != nil) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "error on domain lookup", "data": err, })
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success",  "message": "domain lookup completed", "data": response, })
}

// fetch the DNSData for the specified domain and return data for template
func GetDomain(domain string) (map[string]interface{}, error) {
	dnsClient := setupClients()
	response, err := dnsClient.QueryMultiple(domain)
	if err != nil {
		log.Fatal(err)
	}

	nameservers := getNSRecords(response, dnsClient)
	mx := getMXRecords(response, dnsClient)
	txt := response.TXT
	hosts := getHostRecords(response, dnsClient)

	r := map[string]interface{} {
		"domain": domain,
		"nameservers": nameservers,
		"mx": mx,
		"txt": txt,
		"host": hosts,
	}

	return r, nil
}

// fetch all NS data and return array of hosts
func getNSRecords(r *retryabledns.DNSData, dnsClient *dnsx.DNSX) []Host {
	var hosts []Host
	for _, a := range r.NS {
		var h Host

		ip, err := dnsClient.QueryOne(a)
		if err != nil {
			log.Fatal(err)
		}

		h.Name = a
		h.IP = ip.A[0]

		hosts = append(hosts, h)
	}

	return hosts
}

// fetch all MX records and return array of hosts
func getMXRecords(r *retryabledns.DNSData, dnsClient *dnsx.DNSX) []Host {
	var hosts []Host
	for _, a := range r.MX {
		var h Host

		ip, err := dnsClient.Lookup(a)
		if err != nil {
			log.Fatal(err)
		}

		h.Name = a
		h.IP = ip[0]

		hosts = append(hosts, h)
	}

	return hosts
}

// fetch all host records and return array of hosts
func getHostRecords(r *retryabledns.DNSData, dnsClient *dnsx.DNSX) []Host {
	var hosts []Host
	for _, a := range r.A {
		var h Host

		h.Name = r.Host
		h.IP = a

		hosts = append(hosts, h)
	}

	return hosts
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
package services

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

var disposableDomains = make(map[string]struct{})

func init() {
	resp, err := http.Get("https://raw.githubusercontent.com/disposable/disposable-email-domains/master/domains_strict.txt")
	if err != nil {
		log.Fatal("Error getting disposable domains")
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error getting disposable domains")
		return
	}
	defer resp.Body.Close()
	var domains strings.Builder
	_, err = domains.Write(body)
	if err != nil {
		log.Fatal("Error getting disposable domains")
		return
	}
	for _, domain := range strings.Split(domains.String(), "\n") {
		disposableDomains[domain] = struct{}{}
	}
}

func ValidateMXRecord(domain string) string {
	records, err := net.LookupMX(domain)
	if err != nil {
		return "Invalid domain"
	}
	if len(records) > 0 {
		return "Domain available"
	}
	return "Invalid domain"
}

func CheckDisposableDomain(domain string) string {
	if _, exists := disposableDomains[domain]; exists {
		return "Domain is disposable"
	}
	return "Domain is not disposable"
}

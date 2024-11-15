package services

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"strings"
)

var disposableDomains = make(map[string]struct{})
var mxDomains []string

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
	for _, domain := range records {
		mxDomains = append(mxDomains, domain.Host)
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

func CheckMailDeliverable(mail string) string {
	for _, domain := range mxDomains {
		client, err := smtp.Dial(domain)
		defer client.Close()
		if err == nil {
			err := client.Verify(mail)
			if err == nil {
				return "Mail address is deliverable"
			}
		}
	}
	return "Mail address is undeliverable or we could not connect to server due to server privacy."
}

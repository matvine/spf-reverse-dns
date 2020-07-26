package main

import (
	"fmt"
	"net"
	"strings"
)

type spfRecord struct {
	recursiveDomains []string
	ip               []string
}

func (s *spfRecord) validateIPRecords() {
	for _, address := range s.ip {
		rDNSRecord, err := net.LookupAddr(address)
		if err != nil {
			fmt.Println("Error retrieving ReverseDNS record for " + address)
		} else {
			fmt.Println("Reverse DNS for IPv4 " + address + " : " + rDNSRecord[0])
			aDNSRecord, err := net.LookupHost(rDNSRecord[0])
			if err != nil {
				fmt.Println("Error retrieving A Record for " + rDNSRecord[0])
			} else {
				fmt.Println("A Record for " + rDNSRecord[0] + " : " + aDNSRecord[0])
			}
		}
	}

}

func buildAndValidateSpfRecord(domain string) {
	txtRecord, err := getTxtRecord(domain)
	if err != nil {
		fmt.Println("No TXT record for domain")
		return
	}

	record, err := getSpfRecord(txtRecord)
	if err != nil {
		fmt.Println("TXT Record does not contain an SPF Record")
		return
	}

	var s spfRecord
	fields := strings.Fields(record)
	for _, field := range fields {
		if strings.Contains(field, "include") {
			f := strings.Split(field, ":")
			s.recursiveDomains = append(s.recursiveDomains, f[1])
		}

		if strings.Contains(field, "ip4") {
			f := strings.Split(field, ":")
			s.ip = append(s.ip, f[1])
		}

		if strings.Contains(field, "ip6") {
			f := strings.Split(field, ":")
			s.ip = append(s.ip, f[1])
		}
	}

	if len(s.recursiveDomains) > 0 {
		for _, record := range s.recursiveDomains {
			buildAndValidateSpfRecord(record)
		}
	}

	s.validateIPRecords()
}

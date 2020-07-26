package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func getTxtRecord(domain string) ([]string, error) {
	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("No TXT record for domain: " + domain)
		return []string{}, errors.New("No TXT Record for Domain")
	}
	return txtRecord, nil
}

func getSpfRecord(txtRecord []string) (string, error) {
	fmt.Println("Parsing TXT for SPF Record")
	for _, record := range txtRecord {
		if isSpfRecord(record) {
			return record, nil
		}
	}
	return "", errors.New("no SPF record")
}

func isSpfRecord(record string) bool {
	if strings.Contains(record, "spf") {
		return true
	}
	return false
}

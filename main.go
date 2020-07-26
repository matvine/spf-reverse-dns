package main

import (
	"fmt"
	"os"
)

func main() {

	domain := os.Args[1]

	fmt.Println("Retrieving TXT record for domain: " + domain)
	buildAndValidateSpfRecord(domain)

}

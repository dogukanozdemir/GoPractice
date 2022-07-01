package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, sprRecord, hadDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: couldnt read input %v", err)
	}

}

func checkDomain(domain string) {

	var hasMx, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error :%v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	// https://www.youtube.com/watch?v=aX3-NynM7EM&list=PL5dTjWUk_cPY_xPnFbWWmoFCjvqrDDfh3
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			sprRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v %v %v %v %v %v", domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord)

}

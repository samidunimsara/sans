package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type ASN struct {
	ASN  string `json:"asn"`
	Name string `json:"name"`
}

type IPInfo struct {
	ASN         ASN    `json:"asn"`
	IP          string `json:"ip"`
	Organization string `json:"org"`
	Hostname     string `json:"hostname"`
}

func main() {
	// Command line flags
	subdomainList := flag.String("l", "", "Path to subdomain list file")
	apiToken := flag.String("t", "", "API Token for IP info")
	flag.Parse()

	// Validate the flags
	if *subdomainList == "" || *apiToken == "" {
		log.Fatal("Both -l (subdomain list) and -t (API token) are required.")
	}

	// Read subdomains from file
	subdomains, err := readSubdomains(*subdomainList)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare to save IPs and ASNs, remove duplicates using maps
	seenIPs := make(map[string]bool)
	seenASNs := make(map[string]bool)

	// Open files to save IPs and ASNs
	ipFile, err := os.OpenFile("ip.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening ip.txt:", err)
	}
	defer ipFile.Close()

	asnFile, err := os.OpenFile("asn.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening asn.txt:", err)
	}
	defer asnFile.Close()

	// Loop through subdomains and resolve IPs
	for _, subdomain := range subdomains {
		ipAddresses, err := net.LookupHost(subdomain)
		if err != nil {
			log.Printf("Error resolving subdomain %s: %v\n", subdomain, err)
			continue
		}
		for _, ip := range ipAddresses {
			ipInfo := getIPInfo(subdomain, ip, *apiToken)
			if ipInfo != nil {
				// Save unique IPs and ASNs
				if !seenIPs[ipInfo.IP] {
					seenIPs[ipInfo.IP] = true
					_, err := ipFile.WriteString(ipInfo.IP + "\n")
					if err != nil {
						log.Printf("Error writing IP to file: %v\n", err)
					}
				}
				if !seenASNs[ipInfo.ASN.ASN] {
					seenASNs[ipInfo.ASN.ASN] = true
					_, err := asnFile.WriteString(ipInfo.ASN.ASN + "\n")
					if err != nil {
						log.Printf("Error writing ASN to file: %v\n", err)
					}
				}
			}
		}
	}
}

// Function to get IP info from API
func getIPInfo(subdomain, ip, apiToken string) *IPInfo {
	// Define the API URL with the IP and API token
	apiURL := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, apiToken)

	// Make the HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error getting info for IP %s: %v\n", ip, err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body for IP %s: %v\n", ip, err)
		return nil
	}

	// Parse the JSON response
	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		log.Printf("Error unmarshalling response for IP %s: %v\n", ip, err)
		return nil
	}

	// Ensure correct values for ASN and organization
	if ipInfo.ASN.ASN == "" {
		ipInfo.ASN.ASN = "unknown"
	}
	if ipInfo.ASN.Name == "" {
		ipInfo.ASN.Name = "unknown"
	}
	if ipInfo.Organization == "" {
		ipInfo.Organization = "unknown"
	}

	// Print out the formatted and colored output without debug information
	formatAndPrint(subdomain, ipInfo)
	return &ipInfo
}

// Function to read subdomains from file
func readSubdomains(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	// Split the file content into subdomains
	subdomains := strings.Split(string(data), "\n")
	for i, subdomain := range subdomains {
		subdomains[i] = strings.TrimSpace(subdomain)
	}

	return subdomains, nil
}

// Function to format and print the output with colors
func formatAndPrint(subdomain string, ipInfo IPInfo) {
	// Create color styles
	subdomainColor := color.New(color.FgGreen).SprintFunc()
	ipColor := color.New(color.FgYellow).SprintFunc()
	asnColor := color.New(color.FgCyan).SprintFunc()
	orgColor := color.New(color.FgMagenta).SprintFunc()

	// Print formatted output with colors
	fmt.Printf("%s %s %s %s\n", subdomainColor(subdomain), ipColor(ipInfo.IP), asnColor(ipInfo.ASN.ASN), orgColor(ipInfo.ASN.Name))
}

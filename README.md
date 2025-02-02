# Sans

**Sans** is a lightweight Go tool that resolves subdomains from a provided list, retrieves their associated IP addresses, and fetches ASN (Autonomous System Number) information along with the owning company (organization) details using the [IPInfo.io](https://ipinfo.io/) API. It also saves unique IPs and ASNs to `ip.txt` and `asn.txt` respectively.

---

## Features

- **Subdomain Resolution**: Reads a list of subdomains from a file and resolves their IP addresses.
- **IP Information Retrieval**: For each IP, fetches ASN and organization details using the IPInfo.io API.
- **Formatted Output**: Displays output in a colorized, easy-to-read format.
- **Results Storage**: Saves unique IP addresses and ASN numbers into `ip.txt` and `asn.txt`.

---

## Installation

### Prerequisites

- **Go 1.18+**: Download and install Go from [the official Go downloads page](https://go.dev/dl/).
- **Git**: Make sure Git is installed on your system.

### Using `go install`

You can install **Sans** directly using `go install`:

```bash
go install github.com/samidunimsara/sans@latest
```

### Usage

Once installed, run **Sans** from the command line with the required flags.

### Command-line Options

- `-l`: Path to a file containing a list of subdomains (one per line).
- `-t`: Your IPInfo.io API token.

```
sans -l subdomains.txt -t YOUR_API_TOKEN

```


This will:

1. Resolve each subdomain in the `subdomains.txt` file.
2. Retrieve the IP information for each resolved subdomain, including ASN and organization (owned company) name.
3. Display output in the following format:

```bash
subdomain_name IP_address ASN organization_name
```

For example:

```text
breeno.test.com 123.345.345.444 AS222 ABC ORG
breenofs.test.testcom 222.33.3.54 AS234 IDC, US  Telecommunications Corporation
```

4. Save unique IP addresses to `ip.txt` and ASN numbers to `asn.txt`.
``` 

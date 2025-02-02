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

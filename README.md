# Sub ASN Hunter

A tool to resolve subdomains, get IP addresses, and retrieve IP-related information like ASN and organization name using the IPInfo.io API.

## Features

- Resolves subdomains and their associated IP addresses.
- Fetches ASN and organization details for each IP address using IPInfo.io.
- Saves the IPs to `ip.txt` and ASNs to `asn.txt`, without duplicates.
- Outputs the results in a clean, formatted manner with color-coded information.

## Requirements

- Go (1.18 or higher)
- An active API token from [IPInfo.io](https://ipinfo.io/account/token)

## Installation

### 1. Install the tool using `go install`:

To install the tool globally on your machine, run the following command:

```bash
go install github.com/samidunimsara/sans@latest
```

## Usage

### Run the program:

To use the tool, run the following command:

```bash
sans -l path_to_your_subdomain_list.txt -t YOUR_API_TOKEN
```

- `-l <path_to_your_subdomain_list.txt>`: Path to a file containing a list of subdomains (one subdomain per line).
- `-t <YOUR_API_TOKEN>`: Your IPInfo.io API token.

### Example:

```bash
sans -l subdomains.txt -t fhfhfhfhfh
```

This will:

1. Resolve each subdomain in the `subdomains.txt` file.
2. Retrieve the IP information for each resolved subdomain, including ASN and the organization (owned company) name.

3. Display the output in the terminal with the format:

```
sauapk-auto-cn.coloros.com 106.39.205.146 AS23724 IDC, China Telecommunications Corporation
sauapk-manual-cn.coloros.com 106.3.16.27 AS23724 IDC, China Telecommunications Corporation
saucfs-cost.coloros.com 106.3.16.27 AS23724 IDC, China Telecommunications Corporation
saucfs-cost.coloros.com 106.39.205.146 AS23724 IDC, China Telecommunications Corporation
saucfs.coloros.com 59.110.185.60 AS37963 Hangzhou Alibaba Advertising Co.,Ltd.
saufs-cost.coloros.com 106.3.16.27 AS23724 IDC, China Telecommunications Corporation
saufs-cost.coloros.com 106.39.205.146 AS23724 IDC, China Telecommunications Corporation
saufs.coloros.com 106.39.205.146 AS23724 IDC, China Telecommunications Corporation
saufs.coloros.com 106.3.16.27 AS23724 IDC, China Telecommunications Corporation
saufs1.coloros.com 106.39.205.146 AS23724 IDC, China Telecommunications Corporation
saufsg-cost.coloros.com 23.53.35.73 AS20940 Akamai International B.V.
saufsg-cost.coloros.com 23.53.35.79 AS20940 Akamai International B.V.
saufsg-cost.coloros.com 2600:1408:c400:26::17da:d90c AS20940 Akamai International B.V.
saufsg-cost.coloros.com 2600:1408:c400:26::17da:d931 AS20940 Akamai International B.V.
saufsg.coloros.com 18.136.0.190 AS16509 Amazon.com, Inc.
saufsg.coloros.com 52.220.200.36 AS16509 Amazon.com, Inc.
scc.apps.coloros.com 119.147.175.84 AS4134 CHINANET-BACKBONE
scc.apps.coloros.com 106.3.18.170 AS23724 IDC, China Telecommunications Corporation
sccf-eu.apps.coloros.com 13.37.76.202 AS16509 Amazon.com, Inc.
sccf-eu.apps.coloros.com 13.38.159.60 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 13.250.225.162 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 54.169.168.49 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 18.138.49.3 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 13.229.221.12 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 18.143.15.31 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 18.143.58.170 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 52.76.45.150 AS16509 Amazon.com, Inc.
sccf.apps.coloros.com 18.140.140.23 AS16509 Amazon.com, Inc.
sccfs.coloros.com 106.3.16.27 AS23724 IDC, China Telecommunications Corporation
sccfsf.coloros.com 23.222.16.43 AS20940 Akamai International B.V.

   ```

4. Save the IP addresses to `ip.txt` and ASN numbers to `asn.txt` (without duplicates).

### Files:
- **subdomains.txt**: A file containing subdomains to be resolved (one subdomain per line).
- **ip.txt**: A file where IP addresses are saved (one IP per line).
- **asn.txt**: A file where ASN numbers are saved (one ASN per line).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

If you’d like to contribute to **Subdomain IP Resolver**, follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a pull request.

## Acknowledgements

- Thanks to the creators of the [IPInfo.io](https://ipinfo.io/) API for providing IP and ASN data.

```

This format is now properly structured, and you can copy-paste it into your project’s `README.md`. Just replace the placeholders like `yourusername` and `your_api_token` with actual values. Let me know if you need any more adjustments!

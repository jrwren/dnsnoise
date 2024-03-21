# dnsnoise

DNSNoise is a DNS noise generator in the spirit of noisy.
Rather than use random strings, a list of domain names is referenced.

# Installation

Use `go install` to download and build.
Invoke with the -server flag and your upstream DNS server as argument.
Do not use your local DNS server unless you want its cache filled up with junk
and the useful entries evicted from cache.

```sh
go install github.com/jrwren/dnsnoise
dnsnoise -server 8.8.8.8  # If you use Google DNS.
dnsnoise -server 68.94.156.1  # If you use ATT DNS.
dnsnoise -server 1.1.1.1  # If you use Cloudflare DNS.
# ... etc etc
```

# Usage

```sh
$ dnsnoise -h
Usage of dnsnoise:
  -aaaarr
    	query for AAAA RR before A to look like IPv6 enabled client (default true)
  -csvdomainfile string
    	The top 1M file from https://s3-us-west-1.amazonaws.com/umbrella-static/index.html (default "top-1m.csv")
  -debug int
    	debug level
  -httpsrr
    	query for HTTPS RR before A to look like modern (Apple) client
  -server string
    	a dns server - not your local server, we don't want to bust its cache (default "68.94.156.1:53")

```

## Using the 1M domain file.

Go to https://s3-us-west-1.amazonaws.com/umbrella-static/index.html and download the top-1m.csv file.

Place it in current directory for dnsnoise to read, or run dnsnoise from your Download directory.

## Using the urlhaus malware domains

Go to https://urlhaus.abuse.ch/api/#csv and click one of the links.

Process this file to filter the IP address urls and be CSV:

```sh
awk -F, '{print $3}' csv.txt | awk -F/ '{print $3}' | sort | uniq | grep -nv '^\d' | sed 's/:/,/' > urlhaus.csv
```

The -csvdomainfile flag accepts comma separated list of files.

```sh
dnsnoise -csvdomainfile top-1m.csv,urlhause.csv
```

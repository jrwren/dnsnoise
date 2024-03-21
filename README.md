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
dnsnoise -server 8.8.8.8,8.8.4.4  # If you use Google DNS.
dnsnoise -server 68.94.156.1,68.94.157.1  # If you use ATT DNS.
dnsnoise -server 1.1.1.1,1.0.0.1  # If you use Cloudflare DNS.
# ... etc
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
  -pause duration
    	time to pause between queries (default 1s)
  -server string
    	a dns server - not your local server, we don't want to bust its cache (default "68.94.156.1:53,68.94.157.1:53")

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

# Screenshot

![a screenshot](https://private-user-images.githubusercontent.com/106443/314906981-69a7598b-f73f-40e8-93a8-aa1360dc3595.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5ODcxNTcsIm5iZiI6MTcxMDk4Njg1NywicGF0aCI6Ii8xMDY0NDMvMzE0OTA2OTgxLTY5YTc1OThiLWY3M2YtNDBlOC05M2E4LWFhMTM2MGRjMzU5NS5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjQwMzIxJTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI0MDMyMVQwMjA3MzdaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT05YzJiZjI3YjNhMWM3OWNkNDdmOWY4NjFjNGM0OWM1OTIxNzg0NTAzNzI3NzVjYzJmNTFmYmE1ZWJkYTQ5NWUyJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZhY3Rvcl9pZD0wJmtleV9pZD0wJnJlcG9faWQ9MCJ9.JhtSbCQW2hZ9eg5ePnZFzN9MY5nivvx22dOFsE8E11g "a screenshot")

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/miekg/dns"
)

func main() {
	var server, csvDomainFile string
	var aaaarr, httpsrr bool
	var debug int
	var pause time.Duration
	// TODO: Support multiple servers.
	flag.StringVar(&server, "server", "68.94.156.1:53",
		"a dns server - not your local server, we don't want to bust its cache")
	flag.StringVar(&csvDomainFile, "csvdomainfile", "top-1m.csv",
		"The top 1M file from https://s3-us-west-1.amazonaws.com/umbrella-static/index.html")
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.BoolVar(&aaaarr, "aaaarr", true, "query for AAAA RR before A to look like IPv6 enabled client")
	flag.BoolVar(&httpsrr, "httpsrr", false, "query for HTTPS RR before A to look like modern (Apple) client")
	flag.DurationVar(&pause, "pause", time.Second, "time to pause between queries")
	flag.Parse()
	if server == "" {
		fmt.Println("server must be not be blank.")
		return
	}
	var domains []string
	cdfs := strings.Split(csvDomainFile, ",")
	for i := range cdfs {
		ds, err := loadcsvDomainFile(cdfs[i])
		if err != nil {
			fmt.Printf("error loading domain file: %s\n", err)
			return
		}
		domains = append(domains, ds...)
	}
	rrTypes := []uint16{}
	if httpsrr {
		rrTypes = append(rrTypes, dns.TypeHTTPS)
	}
	if aaaarr {
		rrTypes = append(rrTypes, dns.TypeAAAA)
	}
	rrTypes = append(rrTypes, dns.TypeA)
	c := new(dns.Client)
	m := new(dns.Msg)
	for {
		i := rand.Intn(len(domains))
		for _, rrType := range rrTypes {
			m.SetQuestion(domains[i], rrType)
			msg, rtt, err := c.Exchange(m, server)
			if err != nil {
				log.Println(err)
			} else if debug > 3 {
				log.Printf("%s: msg: %v rtt: %v", domains[i], msg, rtt)
			} else if debug > 1 {
				log.Printf("%s: answer: %v rtt: %v", domains[i], msg.Answer, rtt)
			}
		}
		time.Sleep(pause)
	}
}

func loadcsvDomainFile(fname string) (domains []string, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		_, a, f := strings.Cut(line, ",")
		if !f {
			return domains, fmt.Errorf("comma not found on line %d: %s", i, line)
		}

		a = strings.TrimSpace(a)
		if len(a) == 0 {
			continue
		}
		if !strings.HasSuffix(a, ".") {
			a = a + "."
		}
		domains = append(domains, a)
	}
	if err := scanner.Err(); err != nil {
		return domains, err
	}
	return domains, nil
}

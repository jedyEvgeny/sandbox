package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	site := mustSite()
	ip, host := hostAndIP(site)
	fmt.Printf("ШЗ: %s\n", ip)
	fmt.Printf("Hostname: %s\n", host)
}

func mustSite() string {
	site := flag.String(
		"site",
		"",
		"хост или ip сайта",
	)
	flag.Parse()
	if *site == "" {
		log.Fatal("инфо о сайте не указано через флаг -site")
	}
	return *site
}

func hostAndIP(site string) (net.IP, string) {
	ipOk := net.ParseIP(site)

	var host []string
	var ip []net.IP
	var err error

	if ipOk != nil {
		host, err = net.LookupAddr(site)
		if err != nil {
			log.Fatal(err)
		}
		ip = []net.IP{ipOk}
	}
	if ipOk == nil {
		ip, err = net.LookupIP(site)
		if err != nil {
			log.Fatal(err)
		}
		host = []string{site}
	}
	return ip[0], host[0]
}

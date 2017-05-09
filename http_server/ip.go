package main

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/miekg/dns"
)

func main() {
	ip, _ := LookupIP("www.sina.com", 60)
	fmt.Println(ip)
}

func MyParseIp() {
	name := "192.168.1.1"
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
}

func MyGetIp() {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
}

func LookupIP(domain string, timeout int) ([]string, error) {
	var c dns.Client
	var err error
	ips := []string{}
	domain = strings.TrimRight(domain, ".") + "."
	c.DialTimeout = time.Duration(timeout) * time.Millisecond
	c.ReadTimeout = time.Duration(timeout) * time.Millisecond
	c.WriteTimeout = time.Duration(timeout) * time.Millisecond

	m := new(dns.Msg)
	m.SetQuestion(domain, dns.TypeA)

	ret, _, err := c.Exchange(m, "114.114.114.114:53")
	if err != nil {
		domain = strings.TrimRight(domain, ".")
		e := fmt.Sprintf("lookup error: %s, %s", domain, err.Error())
		return ips, errors.New(e)
	}

	for _, i := range ret.Answer {
		result := strings.Split(i.String(), "\t")
		if result[3] == "A" && IsIP(result[4]) {
			ips = append(ips, result[4])
		}
	}
	return ips, err
}

func IsIP(ip string) bool {
	if ip != "" {
		isOk, _ := regexp.MatchString(`^(\d{1,3}\.){3}\d{1,3}$`, ip)
		if isOk {
			return isOk
		}
	}
	return false
}

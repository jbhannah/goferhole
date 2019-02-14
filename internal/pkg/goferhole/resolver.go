package goferhole

import (
	"log"

	"github.com/miekg/dns"
)

// Resolver defines the interface for mutating DNS resolution handlers
type Resolver interface {
	AddHandler(string, dns.HandlerFunc)
	Forward(dns.ResponseWriter, *dns.Msg)
}

type resolver struct {
	c *dns.Client
}

// NewResolver returns a new instance of a Resolver
func NewResolver() Resolver {
	res := &resolver{c: &dns.Client{}}
	res.AddHandler(".", res.Forward)

	log.Printf("Initialzed resolver")
	return res
}

// AddHandler registers a new handler with the DNS resolver
func (res *resolver) AddHandler(pattern string, handler dns.HandlerFunc) {
	log.Printf("Adding handler for %s", pattern)
	dns.HandleFunc(pattern, handler)
}

// Forward forwards a DNS request to an external resolver
func (res *resolver) Forward(w dns.ResponseWriter, r *dns.Msg) {
	log.Printf("Forwarding question: %v", r.Question)
	rx, _, err := res.c.Exchange(r, "1.1.1.1:53")

	if err != nil {
		log.Printf("Lookup failed: %s\n", err.Error())
		return
	}

	log.Printf("Sending answer: %v", rx.Answer)
	w.WriteMsg(rx)
}

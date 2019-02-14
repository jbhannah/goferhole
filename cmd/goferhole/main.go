package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/miekg/dns"
)

func main() {
	port := 8053

	c := &dns.Client{}

	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		rx, _, err := c.Exchange(r, "1.1.1.1:53")
		if err != nil {
			log.Printf("Lookup failed: %s\n", err.Error())
		}
		w.WriteMsg(rx)
	})

	go func() {
		srv := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set udp listener: %s\n", err.Error())
		}
	}()

	go func() {
		srv := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "tcp"}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set tcp listener: %s\n", err.Error())
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Fatalf("Signal (%v) received, stopping\n", s)
}

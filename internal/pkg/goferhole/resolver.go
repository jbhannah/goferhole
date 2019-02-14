package goferhole

import (
	"log"
	"strconv"

	"github.com/miekg/dns"
)

// Listen starts udp and tcp DNS listeners on the specified port
func Listen(port int) {
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
}

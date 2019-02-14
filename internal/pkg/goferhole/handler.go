package goferhole

import (
	"log"

	"github.com/miekg/dns"
)

func init() {
	c := &dns.Client{}

	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		rx, _, err := c.Exchange(r, "1.1.1.1:53")

		if err != nil {
			log.Printf("Lookup failed: %s\n", err.Error())
			return
		}

		w.WriteMsg(rx)
	})
}

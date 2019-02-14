package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jbhannah/goferhole/internal/pkg/goferhole"
)

func main() {
	port := 8053
	goferhole.Listen(port)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Fatalf("Signal (%v) received, stopping\n", s)
}

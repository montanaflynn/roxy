package main

import (
	"github.com/montanaflynn/roxy"
	"log"
	"net/http"
	"time"
)

func main() {

	rp := roxy.Proxy()
	rp.AddMiddleware(roxy.Cors)
	rp.AddMiddleware(roxy.ConsoleLog)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      rp,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Fatal(s.ListenAndServe())

}

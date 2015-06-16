package main

import (  
  "net/http/httputil"
  "net/http"
  "log"
  "time"
  "github.com/montanaflynn/roxy"
)

func Cors(request *http.Request, response *http.Response) {
  response.Header.Set("Access-Control-Allow-Origin", "*")
}

func Logger(request *http.Request, response *http.Response) {
  reqLog, _ := httputil.DumpRequestOut(request, false)
  resLog, _ := httputil.DumpResponse(response, false)
  log.Printf("\nRequest: %s\nResponse: %s", reqLog, resLog)
}

func main() {
  middleware := []roxy.MiddlewareFunc{Cors, Logger}
  reverse_proxy := roxy.NewReverseProxy(middleware)

  s := &http.Server{
    Addr:           ":8080",
    Handler:        reverse_proxy,
    ReadTimeout:    30 * time.Second,
    WriteTimeout:   30 * time.Second,
  }

  log.Fatal(s.ListenAndServe())
}

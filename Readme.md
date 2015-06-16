# Roxy - A reverse proxy that supports custom middleware

```go
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
  proxy := roxy.NewReverseProxy(middleware)

  s := &http.Server{
    Addr:           ":8080",
    Handler:        proxy,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  log.Fatal(s.ListenAndServe())
}
```

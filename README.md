# Roxy - A proxy with middleware

Just a simple proxy that allows you to modify HTTP requests and responses with a familiar middleware implementation. Once the project is a little more stable there will be roxy-contribs with user contributed middleware. **Use at your own risk!**

### Usage

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

### Roadmap

- Write unit tests for everything
- Create user configuration options
- Integrate with etcd, consul, etc...
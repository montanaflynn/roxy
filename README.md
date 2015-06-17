# Roxy [![Build](https://img.shields.io/wercker/ci/montanaflynn/roxy.svg)](https://app.wercker.com/project/bykey/254e86288a0bbfe8a5aa791a89ff7beb) [![GoDoc](https://godoc.org/github.com/montanaflynn/roxy?status.svg)](https://godoc.org/github.com/montanaflynn/roxy)

A lightweight proxy with a familiar middleware implementation that allows you to modify HTTP requests and responses between clients and servers. Once the project is a little more stable there will be roxy-contribs with user contributed middleware. 

### Usage

Since things are in the alpha stage and quickly changing the best way to run a Roxy server is to build from source. **Use at your own risk!**

Start the example proxy which enables CORS and logging middleware:

```sh
git clone git@github.com:montanaflynn/roxy.git
cd roxy
go build
go run examples/main.go
```

Send a request:

```sh
curl -i localhost:8080 -H "host: anonfunction.com"
```

### Examples

Enable CORS and JSON Logging:

```go
package main

import (
	"net/http"
	"log"
	"time"
	"github.com/montanaflynn/roxy"
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

```

### Tips

You can set your system or browser or even twitter app to use a proxy like Roxy. Instructions vary but for example on a mac I found `networksetup` to be very useful.

First start Roxy on port 8080.

```
git clone git@github.com:montanaflynn/roxy.git
cd roxy
go build
go run examples/main.go
```

Then enable the system wide proxy to go through Roxy.

```sh
networksetup -setwebproxy Wi-Fi localhost 8080
```

You can replace `Wi-Fi` with any network connection which you can find with `networksetup -listnetworkserviceorder`, I always recommend taking a gander at the man page for any new commands (`man networksetup`).

When you want to turn it off it's this simple:

```
networksetup -setwebproxystate Wi-Fi off
```

### Roadmap

- Write unit tests for everything
- Create flags for configuration 
- Integrate with etcd or consul
- Allow for graceful restarts

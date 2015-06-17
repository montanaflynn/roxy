package roxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func ConsoleLog(request *http.Request, response *http.Response) {
	reqLog, _ := httputil.DumpRequestOut(request, false)
	resLog, _ := httputil.DumpResponse(response, false)
	log.Printf("\nRequest:\n%sResponse:\n%s", reqLog, resLog)
}

func JsonLog(request *http.Request, response *http.Response) {
	jsonlog := struct {
		Timestamp int64
		Request   *http.Request
		Response  *http.Response
	}{
		time.Now().Unix(),
		request,
		response,
	}

	l, err := json.Marshal(jsonlog)
	if err != nil {
		return
	}
	fmt.Println(string(l))
}

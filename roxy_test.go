package roxy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func StartRoxy() {

	rp := Proxy()
	rp.AddMiddleware(Cors)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      rp,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	s.ListenAndServe()

}

func TestServer(t *testing.T) {

	go StartRoxy()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"hello": "world"}`)
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Error(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Roxy-Host", u.Host)
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Didn't return a 200 response code")
	}

}

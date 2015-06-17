package roxy

import (
	"net/http"
)

func Cors(request *http.Request, response *http.Response) {
	response.Header.Set("Access-Control-Allow-Origin", "*")
}

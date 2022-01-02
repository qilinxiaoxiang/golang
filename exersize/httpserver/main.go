package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	os.Setenv("VERSION", "myGoVersion")
	http.HandleFunc("/headersEcho", headersEcho)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "200")
}

func headersEcho(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		for _, s := range v {
			writer.Header().Add(k, s)
		}
	}
	fmt.Printf("ip:%s,httpCode:%d\n", strings.Split(request.RemoteAddr, ":")[0], 200)
	writer.Header().Add("VERSION", os.Getenv("VERSION"))
	jsonResp, _ := json.Marshal("success")
	writer.Write(jsonResp)
}

package http_base

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	// URL.Path = "/"
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	// Header["Sec-Ch-Ua-Platform"] = ["\"Windows\""]
	// Header["Upgrade-Insecure-Requests"] = ["1"]
	// Header["Accept-Language"] = ["zh-CN,zh;q=0.9"]
}

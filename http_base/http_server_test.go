package http_base

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type TestHttpServer struct {
}

/**
自定义handler 处理; 验证是否所有请求都会路由到该方法
*/
func (ths TestHttpServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		fmt.Fprintf(response, "URL.Path = %q\n", request.URL.Path)

	case "/hello":
		for k, v := range request.Header {

			fmt.Fprintf(response, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(response, "404 NOT FOUND: %s\n", request.URL)
	}
}

func TestHandel(t *testing.T) {
	ths := new(TestHttpServer)

	err := http.ListenAndServe(":8080", ths)
	if err != nil {
		log.Fatal()
	}

}

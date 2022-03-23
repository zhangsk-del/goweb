package http_base

import (
	"fmt"
	"go-web/http_base/web"
	"net/http"
	"testing"
)

func TestWeb(t *testing.T) {
	web := web.New()

	web.Get("/", func(resp http.ResponseWriter, req *http.Request) {

		fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
	})

	web.Run(":8080")
}

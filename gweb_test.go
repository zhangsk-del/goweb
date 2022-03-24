package go_web

import (
	"fmt"
	"go-web/gweb"
	"net/http"
	"testing"
)

func TestWeb(t *testing.T) {
	engine := gweb.New()

	engine.Get("/", func(resp http.ResponseWriter, req *http.Request) {

		fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
	})

	engine.Run(":8080")
}

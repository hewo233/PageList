package tests

import (
	"github.com/hewo233/PageList/controller"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", controller.ShowIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := io.ReadAll(w.Body)
		pageOK := err == nil &&
			strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})

}

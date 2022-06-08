package tools

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type TestGroup map[string]func(t *testing.T)

func MakeJsonBody(v any) *bytes.Buffer {
	postBody, _ := json.Marshal(v)
	return bytes.NewBuffer(postBody)
}

func CreateHttpConnect(router *gin.Engine, method string, uri string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, body)
	router.ServeHTTP(w, req)
	return w
}

func RunGroup(t *testing.T, group TestGroup) {
	for s, f := range group {
		t.Run(s, f)
	}
}

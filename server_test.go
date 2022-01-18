package webhookForward

import (
	"io"
	"net/http/httptest"
)

type Header map[string]string

func handle(method, path string, headers Header, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	return resp
}

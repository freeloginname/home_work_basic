package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/get_user", nil)
	w := httptest.NewRecorder()
	GetUser(w, req)

}
func Test(t *testing.T) {
	testCases := []struct {
		desc        string
		ip          string
		port        string
		method      string
		path        string
		expectation string
	}{
		{
			desc:        "get user",
			ip:          "127.0.0.1",
			port:        "8080",
			method:      "GET",
			path:        "get_user",
			expectation: `{"id":1,"name":"John Doe","email":"xk0e5@example.com"}`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url := "http://" + tC.ip + ":" + tC.port
			server := server(tC.ip, tC.port)
			client := client(url, tC.method, tC.path)
		})
	}
}

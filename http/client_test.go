package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestClientGet(t *testing.T) {
	server := getTestServer()
	defer server.Close()

	port := getPort(server.URL)
	cli := NewSimpleHTTPClient(WithPort(port), WithInsecureConnection())
	resp, err := cli.Get("/")
	if err != nil {
		t.Fatal(err)
	}
	if "OK" != string(resp) {
		t.Fatal("error on response")
	}
}

func TestClientGetCognos(t *testing.T) {
	server := getTestServer()
	defer server.Close()

	port := getPort(server.URL)
	cli := NewSimpleHTTPClient(WithPort(port), WithInsecureConnection(), WithCognosIntegratedLogin())
	resp, err := cli.Get("/")
	if err != nil {
		t.Fatal(err)
	}
	if "OK Cognos" != string(resp) {
		t.Fatal("error on response")
	}
}

func TestClientPost(t *testing.T) {
	server := getTestServer()
	defer server.Close()

	message := "message"

	port := getPort(server.URL)
	cli := NewSimpleHTTPClient(WithPort(port), WithInsecureConnection())
	resp, err := cli.Post("/", []byte(message))
	if err != nil {
		t.Fatal(err)
	}
	if message != string(resp) {
		t.Fatal("error on response")
	}
}

func getTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if "CAMNamespace YWRtaW46YXBwbGU6" == req.Header["Authorization"][0] {
			rw.Write([]byte("OK Cognos"))
			return
		}

		body, _ := ioutil.ReadAll(req.Body)
		if "" == string(body) {
			rw.Write([]byte("OK"))
		}
		rw.Write(body)
	}))
}

func getPort(url string) int {
	pos := strings.LastIndex(url, ":")
	sub := url[pos+1:]
	p, err := strconv.Atoi(sub)
	if err != nil {
		panic(err)
	}
	return p
}

package http

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	gohttp "net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type simpleHTTPClient struct {
	http         *gohttp.Client
	hostname     string
	port         int
	username     string
	password     string
	namespace    string
	ssl          sslopt
	securityMode securityMode
	timeout      int
}

// NewSimpleHTTPClient creates the standard http client.
// It can be changed like we do in the tests.
func NewSimpleHTTPClient(opts ...option) *simpleHTTPClient {
	cookieJar, _ := cookiejar.New(nil)
	c := &gohttp.Client{
		Jar:     cookieJar,
		Timeout: time.Second * 10,
		Transport: &gohttp.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	cli := &simpleHTTPClient{
		http:         c,
		hostname:     "localhost",
		port:         50000,
		username:     "admin",
		password:     "apple",
		ssl:          true,
		securityMode: basicSecurityMode,
		timeout:      5,
	}

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}

func (cli *simpleHTTPClient) Get(path string) ([]byte, error) {
	response, err := cli.Do("GET", path, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (cli *simpleHTTPClient) Post(path string, data []byte) ([]byte, error) {
	response, err := cli.Do("POST", path, data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (cli *simpleHTTPClient) Do(reqType string, path string, data []byte) ([]byte, error) {
	req, err := cli.request(reqType, path, data)
	if err != nil {
		return nil, err
	}

	response, err := cli.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (cli *simpleHTTPClient) request(reqType string, path string, data []byte) (*gohttp.Request, error) {
	urlx := cli.fullURL(path)
	req, err := gohttp.NewRequest(reqType, urlx, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	cli.ensureCookie(req)

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (cli *simpleHTTPClient) ensureCookie(req *gohttp.Request) error {
	u, err := url.Parse(cli.baseURL())
	if err != nil {
		return err
	}

	authkey := cli.authkey()

	if len(cli.http.Jar.Cookies(u)) < 1 {
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", cli.securityMode, authkey))
	}
	return nil
}

func (cli *simpleHTTPClient) authkey() string {
	namespace := ""
	if cognosSecurityMode == cli.securityMode {
		namespace = ":" + cli.namespace
	}
	return base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s%s", cli.username, cli.password, namespace),
	))
}

func (cli *simpleHTTPClient) baseURL() string {
	u := fmt.Sprintf("http%s://%s:%d/api/", cli.ssl, cli.hostname, cli.port)
	return u
}

func (cli *simpleHTTPClient) fullURL(path string) string {
	baseURL := cli.baseURL()
	fullURL := baseURL + path[1:]
	return fullURL
}

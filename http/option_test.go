package http

import "testing"

func TestSSLOptString(t *testing.T) {
	opt := sslopt(false)
	if "" != opt.String() {
		t.Fatal("error on ssl option false")
	}

	opt = sslopt(true)
	if "s" != opt.String() {
		t.Fatal("error on ssl option true")
	}
}

func TestSecurityMode(t *testing.T) {
	if "Basic" != basicSecurityMode.String() {
		t.Fatal("erro on basic security mode")
	}

	if "CAMNamespace" != cognosSecurityMode.String() {
		t.Fatal("erro on cognos security mode")
	}
}

func TestWithHost(t *testing.T) {
	hostname := "tm1.example.com"
	cli := NewSimpleHTTPClient(WithHost(hostname))
	if hostname != cli.hostname {
		t.Fatal("error on hostname")
	}
}

func TestWithPort(t *testing.T) {
	port := 12345
	cli := NewSimpleHTTPClient(WithPort(port))
	if port != cli.port {
		t.Fatal("error on port number")
	}
}

func TestWithUsername(t *testing.T) {
	username := "maria"
	cli := NewSimpleHTTPClient(WithUsername(username))
	if username != cli.username {
		t.Fatal("error on username")
	}
}

func TestWithPassword(t *testing.T) {
	password := "131231234534"
	cli := NewSimpleHTTPClient(WithPassword(password))
	if password != cli.password {
		t.Fatal("error on password")
	}
}

func TestWithNamespace(t *testing.T) {
	namespace := "LDAP"
	cli := NewSimpleHTTPClient(WithNamespace(namespace))
	if namespace != cli.namespace {
		t.Fatal("error on namespace")
	}
}

func TestWithInsecureConnection(t *testing.T) {
	cli := NewSimpleHTTPClient(WithInsecureConnection())
	if cli.ssl {
		t.Fatal("error on secure connection")
	}
}

func TestWithTimeoutConnection(t *testing.T) {
	seconds := 123
	cli := NewSimpleHTTPClient(WithTimeoutConnection(seconds))
	if seconds != cli.timeout {
		t.Fatal("error on timeout")
	}
}

func TestWithCognosIntegratedLogin(t *testing.T) {
	cli := NewSimpleHTTPClient(WithCognosIntegratedLogin())
	if cognosSecurityMode != cli.securityMode {
		t.Fatal("error on cognosSecurityMode")
	}
}

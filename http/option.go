package http

type sslopt bool

func (sslopt sslopt) String() string {
	if sslopt {
		return "s"
	}
	return ""
}

type securityMode int

func (securityMode securityMode) String() string {
	if cognosSecurityMode == securityMode {
		return "CAMNamespace"
	}
	return "Basic"
}

var basicSecurityMode securityMode = 1
var cognosSecurityMode securityMode = 5

type option func(cli *simpleHTTPClient)

// WithHost provides the option to select the hostname of TM1 server.
func WithHost(hostname string) option {
	return func(cli *simpleHTTPClient) {
		cli.hostname = hostname
	}
}

// WithPort provides the option to select the hostname of TM1 server.
func WithPort(port int) option {
	return func(cli *simpleHTTPClient) {
		cli.port = port
	}
}

// WithUsername provides the option to select the username of TM1 server.
func WithUsername(username string) option {
	return func(cli *simpleHTTPClient) {
		cli.username = username
	}
}

// WithPassword provides the option to select the password of an user of TM1 server.
func WithPassword(password string) option {
	return func(cli *simpleHTTPClient) {
		cli.password = password
	}
}

// WithCognosIntegratedLogin if you are using Cognos Analytics.
func WithCognosIntegratedLogin() option {
	return func(cli *simpleHTTPClient) {
		cli.securityMode = cognosSecurityMode
	}
}

// WithNamespace provides the option to select the namespace of TM1 server.
func WithNamespace(namespace string) option {
	return func(cli *simpleHTTPClient) {
		cli.namespace = namespace
	}
}

// WithInsecureConnection provides insecure connection.
func WithInsecureConnection() option {
	return func(cli *simpleHTTPClient) {
		cli.ssl = false
	}
}

// WithTimeoutConnection provides the timeout in seconds to the connection.
func WithTimeoutConnection(seconds int) option {
	return func(cli *simpleHTTPClient) {
		cli.timeout = seconds
	}
}

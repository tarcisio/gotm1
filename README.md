<h1 align="center">GOTM1</h1>
<p align="center">
  <a href="https://travis-ci.org/tarcisio/gotm1"><img src="https://img.shields.io/travis/tarcisio/gotm1.svg?style=flat-square" alt="Build status"></a>
  <a href="https://goreportcard.com/report/github.com/tarcisio/gotm1"><img src="https://goreportcard.com/badge/github.com/tarcisio/gotm1?style=flat-square&e=2" alt="GoReport"></a>
  <a href="http://godoc.org/github.com/tarcisio/gotm1"><img src="http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square" alt="GoDoc"></a>
  <a href="https://codecov.io/gh/tarcisio/gotm1"><img src="https://img.shields.io/codecov/c/github/tarcisio/gotm1.svg?style=flat-square" /></a>
  <a href="https://codebeat.co/projects/github-com-tarcisio-gotm1-master"><img alt="codebeat badge" src="https://codebeat.co/badges/08cf054c-2b0e-465a-b17e-f6ed34398a59" /></a>
</p>
<hr>

Package gotm1 provides an idiomatic golang implementation of IBM TM1 Planning Analytics REST API,
allowing clients to manage processes, and query data that is stored in the model.

## Audience
To use this library effectively, you must be familiar with the following areas:
* TM1 and the TM1 architecture
* Dimensional data and modeling terminology and concepts
* Go (golang) programming language

## Installation and configuration
This library has no other dependency so you can `go get` like so:

```bash
$ go get -u https://github.com/tarcisio/gotm1
```

This project is compatible with `go modules`.

## How to use
The main package is `github.com/tarcisio/gotm1`, to make it nicely decoupled
and make easier to test, the http client was created.

```go
import (
	"fmt"
	"log"

	"github.com/tarcisio/gotm1"
	"github.com/tarcisio/gotm1/http"
)

tm1 := gotm1.New(
  http.NewSimpleHTTPClient(
    http.WithHost("tm1.example.com"),
    http.WithPort(8091),
    http.WithUsername("c6548654"),
    http.WithPassword("p@$$w0rd"),
    http.WithNamespace("CompanyLDAP"),
    http.WithCognosIntegratedLogin(),
    http.WithTimeoutConnection(5),
  ),
)
defer tm1.Logout()

cubes, err := tm1.GetCubes()

for _, cube := range cubes {
  fmt.Println(cube.Name)
}
```
package gotm1_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/tarcisio/gotm1"
	"github.com/tarcisio/gotm1/http"
)

func getRealIntegrated() *gotm1.TM1 {
	envHostname := os.Getenv("TM1_TEST_HOSTNAME")
	envPort := os.Getenv("TM1_TEST_PORT")
	envNamespace := os.Getenv("TM1_TEST_NAMESPACE")
	envUsername := os.Getenv("TM1_TEST_USERNAME")
	envPassword := os.Getenv("TM1_TEST_PASSWORD")

	port, err := strconv.Atoi(envPort)
	if err != nil {
		panic(err)
	}

	return gotm1.New(
		http.NewSimpleHTTPClient(
			http.WithHost(envHostname),
			http.WithPort(port),
			http.WithCognosIntegratedLogin(),
			http.WithNamespace(envNamespace),
			http.WithUsername(envUsername),
			http.WithPassword(envPassword),
			http.WithTimeoutConnection(10),
		),
	)
}

func getRealBasic() *gotm1.TM1 {

	envHostname := os.Getenv("TM1_TEST_HOSTNAME")
	envPort := os.Getenv("TM1_TEST_BASIC_PORT")
	envUsername := os.Getenv("TM1_TEST_BASIC_USERNAME")
	envPassword := os.Getenv("TM1_TEST_BASIC_PASSWORD")

	port, err := strconv.Atoi(envPort)
	if err != nil {
		panic(err)
	}

	return gotm1.New(
		http.NewSimpleHTTPClient(
			http.WithHost(envHostname),
			http.WithPort(port),
			http.WithUsername(envUsername),
			http.WithPassword(envPassword),
			http.WithInsecureConnection(),
			http.WithTimeoutConnection(10),
		),
	)
}

func TestTM1RealIntegrated(t *testing.T) {

	if os.Getenv("TM1_TEST_HOSTNAME") == "" {
		t.Skip("skipping test; $TM1_TEST_HOSTNAME not set")
	}

	testTM1Service := getRealIntegrated()
	defer testTM1Service.Logout()

	config, err := testTM1Service.GetConfiguration()
	if err != nil {
		t.Fatal(err)
	}

	if "CAM" != config.SecurityMode {
		t.Fatal("error on Real Integrated server security mode")
	}

	cubes, err := testTM1Service.GetCubes()
	if err != nil {
		t.Fatal(err)
	}

	found := false
	expected := "}Capabilities"

	for _, cube := range cubes {
		if expected == cube.Name {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Cube expected not found")
	}
}

func TestTM1RealBasic(t *testing.T) {

	if os.Getenv("TM1_TEST_HOSTNAME") == "" {
		t.Skip("skipping test; $TM1_TEST_HOSTNAME not set")
	}

	testTM1Service := getRealBasic()
	defer testTM1Service.Logout()

	config, err := testTM1Service.GetConfiguration()
	if err != nil {
		t.Fatal(err)
	}

	if "Basic" != config.SecurityMode {
		t.Fatal("error on Real Basic server security mode")
	}

	cubes, err := testTM1Service.GetCubes()
	if err != nil {
		t.Fatal(err)
	}

	found := false
	expected := "}Capabilities"

	for _, cube := range cubes {
		if expected == cube.Name {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Cube expected not found")
	}
}

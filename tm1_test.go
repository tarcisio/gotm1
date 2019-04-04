package gotm1_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/tarcisio/gotm1"
)

type testcli struct {
	testConfig     gotm1.Configuration
	testCubes      []gotm1.Cube
	testDimensions []gotm1.Dimension
	testChores     []gotm1.Chore
	testProcesses  []gotm1.Process
	testLoggers    []gotm1.Logger
	testSessions   []gotm1.Session
	testThreads    []gotm1.Thread
	testActiveUser gotm1.User

	testProcessExecuteName   string
	testReturnProcessExecute gotm1.ProcessExecuteResult
	testProcessEvenParams    []interface{}
	testProcessOddParams     []interface{}

	failOnConnection bool
	malformedJSON    bool
}

var defaultTestCli *testcli = &testcli{
	testConfig: gotm1.Configuration{
		ServerName: "gotm1",
	},
	testCubes: []gotm1.Cube{
		{Name: "one"},
		{Name: "}two"},
	},
	testDimensions: []gotm1.Dimension{
		{Name: "a"},
		{Name: "}b"},
	},
	testChores: []gotm1.Chore{
		{Name: "la"},
		{Name: "lo"},
	},
	testLoggers: []gotm1.Logger{
		{Name: "po"},
		{Name: "ko"},
	},
	testSessions: []gotm1.Session{
		{ID: 1},
		{ID: 2},
	},
	testThreads: []gotm1.Thread{
		{ID: 9},
		{ID: 8},
	},
	testActiveUser: gotm1.User{
		FriendlyName: "tarcisio",
	},
	testProcessExecuteName: "test_process",
	testReturnProcessExecute: gotm1.ProcessExecuteResult{
		ProcessExecuteStatusCode: "ok",
	},
	testProcessEvenParams: []interface{}{"p_ano", "2019"},
	testProcessOddParams:  []interface{}{"p_ano"},
}

func (cli *testcli) Get(path string) ([]byte, error) {
	if cli.failOnConnection {
		cli.failOnConnection = false
		return nil, errors.New("test cli connection error")
	}

	if cli.malformedJSON {
		cli.malformedJSON = false
		return []byte(`{"ServerName`), nil
	}

	if "/v1/Configuration" == path {
		data, _ := json.Marshal(cli.testConfig)
		return data, nil
	}

	if "/v1/Cubes" == path {
		resp := struct {
			Value []gotm1.Cube `json:"value"`
		}{Value: cli.testCubes}

		data, _ := json.Marshal(resp)
		return data, nil
	}

	if "/v1/Dimensions" == path {
		resp := struct {
			Value []gotm1.Dimension `json:"value"`
		}{Value: cli.testDimensions}

		data, _ := json.Marshal(resp)
		return data, nil
	}

	if "/v1/Chores" == path {
		resp := struct {
			Value []gotm1.Chore `json:"value"`
		}{Value: cli.testChores}

		data, _ := json.Marshal(resp)
		return data, nil
	}

	if "/v1/Processes" == path {
		resp := struct {
			Value []gotm1.Process `json:"value"`
		}{Value: cli.testProcesses}

		data, _ := json.Marshal(resp)
		return data, nil
	}

	if "/v1/Loggers" == path {
		resp := struct {
			Value []gotm1.Logger `json:"value"`
		}{Value: cli.testLoggers}

		data, _ := json.Marshal(resp)
		return data, nil
	}

	if "/v1/Sessions" == path {
		resp := struct {
			Value []gotm1.Session `json:"value"`
		}{Value: cli.testSessions}

		data, _ := json.Marshal(resp)
		return data, nil
	}

	if "/v1/Threads" == path {
		resp := struct {
			Value []gotm1.Thread `json:"value"`
		}{Value: cli.testThreads}

		data, _ := json.Marshal(resp)
		return data, nil
	}

	if "/v1/ActiveUser" == path {
		data, _ := json.Marshal(cli.testActiveUser)
		return data, nil
	}

	return nil, errors.New("error not expected. This is a test client.")
}

func (cli *testcli) Post(path string, data []byte) ([]byte, error) {
	if cli.failOnConnection {
		cli.failOnConnection = false
		return nil, errors.New("test cli connection error")
	}

	if cli.malformedJSON {
		cli.malformedJSON = false
		return nil, nil
	}

	if "/v1/ActiveSession/tm1.Close" == path {
		return nil, nil
	}

	if "/v1/Processes('"+cli.testProcessExecuteName+"')/tm1.ExecuteWithReturn" == path {
		data, _ := json.Marshal(cli.testReturnProcessExecute)
		return data, nil
	}

	return nil, errors.New("error not expected. This is a test client.")
}

func TestTM1GetConfiguration(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	config, err := testTM1Service.GetConfiguration()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	if config.ServerName != defaultTestCli.testConfig.ServerName {
		t.Fatal("GetConfiguration error. Server name mismatch. This is an unexpected error.")
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetConfiguration()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetConfiguration()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1GetCubes(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	cubes, err := testTM1Service.GetCubes()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	for i, c := range cubes {
		if c.Name != defaultTestCli.testCubes[i].Name {
			t.Fatal("GetCubes error. Cube name mismatch. This is an unexpected error.")
		}
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetCubes()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetCubes()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1GetDimensions(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	dimensions, err := testTM1Service.GetDimensions()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	for i, d := range dimensions {
		if d.Name != defaultTestCli.testDimensions[i].Name {
			t.Fatal("GetDimensions error. Cube name mismatch. This is an unexpected error.")
		}
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetDimensions()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetDimensions()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1GetChores(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	chores, err := testTM1Service.GetChores()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	for i, c := range chores {
		if c.Name != defaultTestCli.testChores[i].Name {
			t.Fatal("GetChores error. Chore name mismatch. This is an unexpected error.")
		}
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetChores()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetChores()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1GetProcess(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	processes, err := testTM1Service.GetProcesses()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	for i, p := range processes {
		if p.Name != defaultTestCli.testProcesses[i].Name {
			t.Fatal("GetProcesses error. Process name mismatch. This is an unexpected error.")
		}
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetProcesses()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetProcesses()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1GetLoggers(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	loggers, err := testTM1Service.GetLoggers()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	for i, l := range loggers {
		if l.Name != defaultTestCli.testLoggers[i].Name {
			t.Fatal("GetLoggers error. Logger name mismatch. This is an unexpected error.")
		}
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetLoggers()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetLoggers()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1GetSessions(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	sessions, err := testTM1Service.GetSessions()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	for i, l := range sessions {
		if l.ID != defaultTestCli.testSessions[i].ID {
			t.Fatal("GetSessions error. Session name mismatch. This is an unexpected error.")
		}
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetSessions()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetSessions()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1GetThreads(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	threads, err := testTM1Service.GetThreads()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	for i, th := range threads {
		if th.ID != defaultTestCli.testThreads[i].ID {
			t.Fatal("GetThreads error. Thread ID mismatch. This is an unexpected error.")
		}
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.GetThreads()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.GetThreads()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestActiveUser(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	activeUser, err := testTM1Service.ActiveUser()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	if activeUser.FriendlyName != defaultTestCli.testActiveUser.FriendlyName {
		t.Fatal("ActiveUser error. FriendlyName mismatch. This is an unexpected error.")
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.ActiveUser()
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.ActiveUser()
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestProcessExecuteWithReturn(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)

	processReturn, err := testTM1Service.ProcessExecuteWithReturn(defaultTestCli.testProcessExecuteName)
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	if processReturn.ProcessExecuteStatusCode != defaultTestCli.testReturnProcessExecute.ProcessExecuteStatusCode {
		t.Fatal("ProcessExecuteWithReturn error. ProcessExecuteStatusCode mismatch. This is an unexpected error.")
	}

	_, err = testTM1Service.ProcessExecuteWithReturn(
		defaultTestCli.testProcessExecuteName,
		defaultTestCli.testProcessEvenParams...)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testTM1Service.ProcessExecuteWithReturn(
		defaultTestCli.testProcessExecuteName,
		defaultTestCli.testProcessOddParams...)
	if err == nil {
		t.Fatal("error expected")
	}

	defaultTestCli.failOnConnection = true
	_, err = testTM1Service.ProcessExecuteWithReturn(defaultTestCli.testProcessExecuteName)
	if err == nil {
		t.Fatal("error expected on connection.")
	}

	defaultTestCli.malformedJSON = true
	_, err = testTM1Service.ProcessExecuteWithReturn(defaultTestCli.testProcessExecuteName)
	if err == nil {
		t.Fatal("error expected on malformed json.")
	}
}

func TestTM1Logout(t *testing.T) {
	var testTM1Service *gotm1.TM1 = gotm1.New(defaultTestCli)
	err := testTM1Service.Logout()
	if err != nil {
		t.Fatal("unexpected error. This is a test client.")
	}

	defaultTestCli.failOnConnection = true
	err = testTM1Service.Logout()
	if err == nil {
		t.Fatal("error expected on connection.")
	}
}

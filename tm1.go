package gotm1

// HTTPClient is responsible to call the REST API doing the low level
// interaction. This alleviate the code on the TM1 struct.
type HTTPClient interface {
	Get(path string) ([]byte, error)
	Post(path string, data []byte) ([]byte, error)
}

// TM1 Service
type TM1 struct {
	cli HTTPClient
}

// New creates a new TM1 rest API client with options.
func New(httpClient HTTPClient) *TM1 {
	return &TM1{httpClient}
}

// GetConfiguration of a TM1 server.
func (tm1 *TM1) GetConfiguration() (Configuration, error) {
	data, err := tm1.cli.Get("/v1/Configuration")
	if err != nil {
		return Configuration{}, err
	}
	return unmarshalConfiguration(data)
}

// GetCubes returns all cubes in a TM1 server.
func (tm1 *TM1) GetCubes() ([]Cube, error) {
	data, err := tm1.cli.Get("/v1/Cubes")
	if err != nil {
		return nil, err
	}
	return unmarshalCubes(data)
}

// GetDimensions get all dimensions in a TM1 server.
func (tm1 *TM1) GetDimensions() ([]Dimension, error) {
	data, err := tm1.cli.Get("/v1/Dimensions")
	if err != nil {
		return nil, err
	}
	return unmarshalDimensions(data)
}

// GetChores get all chores in a TM1 server.
func (tm1 *TM1) GetChores() ([]Chore, error) {
	data, err := tm1.cli.Get("/v1/Chores")
	if err != nil {
		return nil, err
	}
	return unmarshalChores(data)
}

// GetProcesses returns all processes in a TM1 server.
func (tm1 *TM1) GetProcesses() ([]Process, error) {
	data, err := tm1.cli.Get("/v1/Processes")
	if err != nil {
		return nil, err
	}
	return unmarshalProcesses(data)
}

// GetLoggers returns all Logger options in a TM1 server.
func (tm1 *TM1) GetLoggers() ([]Logger, error) {
	data, err := tm1.cli.Get("/v1/Loggers")
	if err != nil {
		return nil, err
	}
	return unmarshalLoggers(data)
}

// GetSessions get all sessions in a TM1 server.
func (tm1 *TM1) GetSessions() ([]Session, error) {
	data, err := tm1.cli.Get("/v1/Sessions")
	if err != nil {
		return nil, err
	}
	return unmarshalSessions(data)
}

// GetThreads get all threads in a TM1 server.
func (tm1 *TM1) GetThreads() ([]Thread, error) {
	data, err := tm1.cli.Get("/v1/Threads")
	if err != nil {
		return nil, err
	}
	return unmarshalThreads(data)
}

// ActiveUser get informations about the active user.
func (tm1 *TM1) ActiveUser() (User, error) {
	data, err := tm1.cli.Get("/v1/ActiveUser")
	if err != nil {
		return User{}, err
	}
	return unmarshalUser(data)
}

// ProcessExecuteWithReturn execute a process returning the resulting code.
func (tm1 *TM1) ProcessExecuteWithReturn(name string, parameters ...interface{}) (ProcessExecuteResult, error) {
	dpar, err := marshalProcessParameters(parameters)
	if err != nil {
		return ProcessExecuteResult{}, err
	}

	data, err := tm1.cli.Post("/v1/Processes('"+name+"')/tm1.ExecuteWithReturn", dpar)
	if err != nil {
		return ProcessExecuteResult{}, err
	}

	return unmarshalProcessExecuteResult(data)
}

// Logout the current session.
func (tm1 *TM1) Logout() error {
	_, err := tm1.cli.Post("/v1/ActiveSession/tm1.Close", []byte("{}"))
	return err
}

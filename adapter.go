package gotm1

import (
	"encoding/json"
	"errors"
)

func unmarshalConfiguration(data []byte) (conf Configuration, err error) {
	err = json.Unmarshal(data, &conf)
	return
}

func unmarshalCubes(data []byte) ([]Cube, error) {
	response := struct {
		Value []Cube `json:"value"`
	}{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Value, nil
}

func unmarshalDimensions(data []byte) ([]Dimension, error) {
	response := struct {
		Value []Dimension `json:"value"`
	}{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Value, nil
}

func unmarshalChores(data []byte) ([]Chore, error) {
	response := struct {
		Value []Chore `json:"value"`
	}{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Value, nil
}

func unmarshalProcesses(data []byte) ([]Process, error) {
	response := struct {
		Value []Process `json:"value"`
	}{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Value, nil
}

func unmarshalLoggers(data []byte) ([]Logger, error) {
	response := struct {
		Value []Logger `json:"value"`
	}{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Value, nil
}

func unmarshalSessions(data []byte) ([]Session, error) {
	response := struct {
		Value []Session `json:"value"`
	}{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Value, nil
}

func unmarshalThreads(data []byte) ([]Thread, error) {
	response := struct {
		Value []Thread `json:"value"`
	}{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Value, nil
}

func unmarshalUser(data []byte) (user User, err error) {
	err = json.Unmarshal(data, &user)
	if err != nil {
		return
	}
	return
}

func unmarshalProcessExecuteResult(data []byte) (result ProcessExecuteResult, err error) {
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

func marshalProcessParameters(parameters []interface{}) ([]byte, error) {
	if isEven(parameters) {
		return nil, errors.New("parameters is not even")
	}
	pstr := makeProcessParametersStruct(parameters)
	return json.Marshal(pstr)
}

func isEven(parameters []interface{}) bool {
	return len(parameters)%2 != 0
}

func makeProcessParametersStruct(parameters []interface{}) struct{ Parameters []map[string]interface{} } {
	pstr := struct{ Parameters []map[string]interface{} }{Parameters: []map[string]interface{}{}}
	for i := range parameters {
		if i%2 == 0 {
			pstr.Parameters = append(pstr.Parameters, map[string]interface{}{
				"Name":  parameters[i],
				"Value": parameters[i+1],
			})
		}
	}
	return pstr
}

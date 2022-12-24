package services

import (
	"errors"
	"fmt"

	"github.com/style77/appwrite-for-go"
)

type Functions struct {
	client *appwrite.Client
}

func NewFunctions(client *appwrite.Client) *Functions {
	return &Functions{client: client}
}

// List Functions
func (f *Functions) List(queries []string, search string) (interface{}, error) {
	url := fmt.Sprintf("/functions")
	payload := make(map[string]interface{})

	if len(queries) > 0 {
		payload["queries"] = queries
	}
	if search != "" {
		payload["search"] = search
	}

	res, err := f.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create Function
func (f *Functions) Create(functionId string, name string, execute []string, runtime string, events []string, schedule string, timeout *int, enabled *bool) (interface{}, error) {
	if functionId == "" {
		return nil, errors.New("Missing required parameter: \"functionId\"")
	}
	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}
	if len(execute) == 0 {
		return nil, errors.New("Missing required parameter: \"execute\"")
	}
	if runtime == "" {
		return nil, errors.New("Missing required parameter: \"runtime\"")
	}

	url := fmt.Sprintf("/functions")
	payload := make(map[string]interface{})

	payload["functionId"] = functionId
	payload["name"] = name
	payload["execute"] = execute
	payload["runtime"] = runtime

	if len(events) > 0 {
		payload["events"] = events
	}
	if schedule != "" {
		payload["schedule"] = schedule
	}
	if timeout != nil {
		payload["timeout"] = timeout
	}
	if enabled != nil {
		payload["enabled"] = enabled
	}

	res, err := f.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// List runtimes
func (f *Functions) ListRuntimes() (interface{}, error) {
	url := fmt.Sprintf("/functions/runtimes")
	payload := make(map[string]interface{})

	res, err := f.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Function
func (f *Functions) Get(functionId string) (interface{}, error) {
	if functionId == "" {
		return nil, errors.New("Missing required parameter: \"functionId\"")
	}

	url := fmt.Sprintf("/functions/%s", functionId)
	payload := make(map[string]interface{})

	res, err := f.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update Function
func (f *Functions) Update(functionId string, name string, execute []string, events []string, schedule string, timeout *int, enabled *bool) (interface{}, error) {
	if functionId == "" {
		return nil, errors.New("Missing required parameter: \"functionId\"")
	}
	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}
	if len(execute) == 0 {
		return nil, errors.New("Missing required parameter: \"execute\"")
	}

	url := fmt.Sprintf("/functions/%s", functionId)
	payload := make(map[string]interface{})

	payload["name"] = name
	payload["execute"] = execute

	if len(events) > 0 {
		payload["events"] = events
	}
	if schedule != "" {
		payload["schedule"] = schedule
	}
	if timeout != nil {
		payload["timeout"] = timeout
	}
	if enabled != nil {
		payload["enabled"] = enabled
	}

	res, err := f.client.Call("put", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete Function
func (f *Functions) Delete(functionId string) (interface{}, error) {
	if functionId == "" {
		return nil, errors.New("Missing required parameter: \"functionId\"")
	}

	url := fmt.Sprintf("/functions/%s", functionId)
	payload := make(map[string]interface{})

	res, err := f.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// List deployments
func (f *Functions) ListDeployments(functionId string, queries []string, search string) (interface{}, error) {
	if functionId == "" {
		return nil, errors.New("Missing required parameter: \"functionId\"")
	}

	url := fmt.Sprintf("/functions/%s/deployments", functionId)
	payload := make(map[string]interface{})

	if len(queries) > 0 {
		payload["queries"] = queries
	}
	if search != "" {
		payload["search"] = search
	}

	res, err := f.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create deployment TODO buffer
func (f *Functions) CreateDeployment(functionId string, entryPoint string, code *appwrite.InputFile, activate *bool) (interface{}, error) {
	if functionId == "" {
		return nil, errors.New("Missing required parameter: \"functionId\"")
	}
	if entryPoint == "" {
		return nil, errors.New("Missing required parameter: \"entryPoint\"")
	}
	if code == nil {
		return nil, errors.New("Missing required parameter: \"code\"")
	}
	if activate == nil {
		return nil, errors.New("Missing required parameter: \"activate\"")
	}

	url := fmt.Sprintf("/functions/%s/deployments", functionId)
	payload := make(map[string]interface{})

	payload["entryPoint"] = entryPoint
	payload["code"] = code
	payload["activate"] = activate

	res, err := f.client.Call("post", url, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (f *Functions) GetDeployment(functionId string, deploymentId string) (interface{}, error) {
	if functionId == "" {
		return nil, errors.New("Missing required parameter: \"functionId\"")
	}
	if deploymentId == "" {
		return nil, errors.New("Missing required parameter: \"deploymentId\"")
	}

	url := fmt.Sprintf("/functions/%s/deployments/%s", functionId, deploymentId)
	payload := make(map[string]interface{})

	res, err := f.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// WIP todo https://github.com/appwrite/sdk-for-node/blob/master/lib/services/functions.js#L480
// For now i dont have enough time to implement this
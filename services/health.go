package services

import (
	"fmt"

	"github.com/style77/appwrite-for-go"
)

type Health struct {
	client *appwrite.Client
}

func NewHealth(client *appwrite.Client) *Health {
	return &Health{client: client}
}

// Get HTTP
func (h *Health) Get() (interface{}, error) {
	url := fmt.Sprintf("/health")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Anti Virus
func (h *Health) GetAntiVirus() (interface{}, error) {
	url := fmt.Sprintf("/health/anti-virus")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Cache
func (h *Health) GetCache() (interface{}, error) {
	url := fmt.Sprintf("/health/cache")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get DB
func (h *Health) GetDB() (interface{}, error) {
	url := fmt.Sprintf("/health/db")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Certificate Queue
func (h *Health) GetQueueCertificates() (interface{}, error) {
	url := fmt.Sprintf("/health/queue/certificates")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Functions Queue
func (h *Health) GetQueueFunctions() (interface{}, error) {
	url := fmt.Sprintf("/health/queue/functions")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Logs Queue
func (h *Health) GetQueueLogs() (interface{}, error) {
	url := fmt.Sprintf("/health/queue/logs")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Webhooks Queue
func (h *Health) GetQueueWebhooks() (interface{}, error) {
	url := fmt.Sprintf("/health/queue/webhooks")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Storage local
func (h *Health) GetStorageLocal() (interface{}, error) {
	url := fmt.Sprintf("/health/storage/local")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Time
func (h *Health) GetTime() (interface{}, error) {
	url := fmt.Sprintf("/health/time")

	res, err := h.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}
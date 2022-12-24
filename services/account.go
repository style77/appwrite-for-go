package services

import (
	"errors"
	"fmt"

	"github.com/style77/appwrite-for-go"
)

type Account struct {
	client *appwrite.Client
}

func NewAccount(client *appwrite.Client) *Account {
	return &Account{client: client}
}

func (a *Account) Get() (interface{}, error) {
	url := fmt.Sprintf("/account")

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdateEmail(email string, password string) (interface{}, error) {
	if email == "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}
	if password == "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	url := fmt.Sprintf("/account/email")

	payload := make(map[string]interface{})

	payload["email"] = email
	payload["password"] = password

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) ListLogs(queries []string) (interface{}, error) {
	url := fmt.Sprintf("/account/logs")

	payload := make(map[string]interface{})

	if len(queries) > 0 {
		payload["queries"] = queries
	}

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdateName(name string) (interface{}, error) {
	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}

	url := fmt.Sprintf("/account/name")

	payload := make(map[string]interface{})

	payload["name"] = name

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdatePassword(password string, oldPassword string) (interface{}, error) {
	if password == "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}
	if oldPassword == "" {
		return nil, errors.New("Missing required parameter: \"oldPassword\"")
	}

	url := fmt.Sprintf("/account/password")

	payload := make(map[string]interface{})

	payload["password"] = password
	payload["old-password"] = oldPassword

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdatePhone(phone string, password string) (interface{}, error) {
	if phone == "" {
		return nil, errors.New("Missing required parameter: \"phone\"")
	}
	if password == "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	url := fmt.Sprintf("/account/phone")

	payload := make(map[string]interface{})

	payload["phone"] = phone
	payload["password"] = password

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) GetPrefs() (interface{}, error) {
	url := fmt.Sprintf("/account/prefs")

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdatePrefs(prefs map[string]interface{}) (interface{}, error) {
	if prefs == nil {
		return nil, errors.New("Missing required parameter: \"prefs\"")
	}

	url := fmt.Sprintf("/account/prefs")

	payload := make(map[string]interface{})

	payload["prefs"] = prefs

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) CreateRecovery(email string, url string) (interface{}, error) {
	if email == "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}
	if url == "" {
		return nil, errors.New("Missing required parameter: \"url\"")
	}

	url = fmt.Sprintf("/account/recovery")

	payload := make(map[string]interface{})

	payload["email"] = email
	payload["url"] = url

	res, err := a.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdateRecovery(userId string, secret string, password string, passwordAgain string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}
	if secret == "" {
		return nil, errors.New("Missing required parameter: \"secret\"")
	}
	if password == "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}
	if passwordAgain == "" {
		return nil, errors.New("Missing required parameter: \"passwordAgain\"")
	}

	url := fmt.Sprintf("/account/recovery")
	payload := make(map[string]interface{})

	payload["userId"] = userId
	payload["secret"] = secret
	payload["password"] = password
	payload["passwordAgain"] = passwordAgain

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) ListSessions() (interface{}, error) {
	url := fmt.Sprintf("/account/sessions")

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) DeleteSessions() (interface{}, error) {
	url := fmt.Sprintf("/account/sessions")

	res, err := a.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) GetSession(sessionId string) (interface{}, error) {
	if sessionId == "" {
		return nil, errors.New("Missing required parameter: \"sessionId\"")
	}

	url := fmt.Sprintf("/account/sessions/%s", sessionId)

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdateSession(sessionId string) (interface{}, error) {
	if sessionId == "" {
		return nil, errors.New("Missing required parameter: \"sessionId\"")
	}

	url := fmt.Sprintf("/account/sessions/%s", sessionId)

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) DeleteSession(sessionId string) (interface{}, error) {
	if sessionId == "" {
		return nil, errors.New("Missing required parameter: \"sessionId\"")
	}

	url := fmt.Sprintf("/account/sessions/%s", sessionId)

	res, err := a.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdateStatus() (interface{}, error) {
	url := fmt.Sprintf("/account/status")

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) CreateVerification(url string) (interface{}, error) {
	if url == "" {
		return nil, errors.New("Missing required parameter: \"url\"")
	}

	url = fmt.Sprintf("/account/verification")

	payload := make(map[string]interface{})

	payload["url"] = url

	res, err := a.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdateVerification(userId string, secret string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}
	if secret == "" {
		return nil, errors.New("Missing required parameter: \"secret\"")
	}

	url := fmt.Sprintf("/account/verification")
	payload := make(map[string]interface{})
	payload["userId"] = userId
	payload["secret"] = secret

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) CreatePhoneVerification() (interface{}, error) {
	url := fmt.Sprintf("/account/verification/phone")

	res, err := a.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Account) UpdatePhoneVerification(userId string, secret string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}
	if secret == "" {
		return nil, errors.New("Missing required parameter: \"secret\"")
	}

	url := fmt.Sprintf("/account/verification/phone")
	payload := make(map[string]interface{})

	payload["userId"] = userId
	payload["secret"] = secret

	res, err := a.client.Call("patch", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}
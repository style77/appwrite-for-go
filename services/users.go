package services

import (
	"errors"
	"fmt"

	"github.com/style77/sdk-for-go"
)

type Users struct {
	client *appwrite.Client
}

func NewUsers(client *appwrite.Client) *Users {
	return &Users{client: client}
}

func (u *Users) List(queries []string, search string) (interface{}, error) {
	url := fmt.Sprintf("/users")
	payload := make(map[string]interface{})

	if len(queries) > 0 {
		payload["queries"] = queries
	}
	if search != "" {
		payload["search"] = search
	}

	res, err := u.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Users) Create(userId string, email string, phone string, password string, name string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}

	url := fmt.Sprintf("/users")
	payload := make(map[string]interface{})

	payload["userId"] = userId

	if email != "" {
		payload["email"] = email
	}
	if phone != "" {
		payload["phone"] = phone
	}
	if password != "" {
		payload["password"] = password
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := u.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Users) CreateArgon2User(userId string, email string, password string, name string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}

	if email != "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}

	if password != "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	url := fmt.Sprintf("/users/argon2")
	payload := make(map[string]interface{})

	payload["userId"] = userId

	if email != "" {
		payload["email"] = email
	}
	if password != "" {
		payload["password"] = password
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := u.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Users) CreateBcryptUser(userId string, email string, password string, name string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}

	if email != "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}

	if password != "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	url := fmt.Sprintf("/users/bcrypt")
	payload := make(map[string]interface{})

	payload["userId"] = userId

	if email != "" {
		payload["email"] = email
	}
	if password != "" {
		payload["password"] = password
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := u.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Users) CreateMD5User(userId string, email string, password string, name string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}

	if email != "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}

	if password != "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	url := fmt.Sprintf("/users/md5")
	payload := make(map[string]interface{})

	payload["userId"] = userId

	if email != "" {
		payload["email"] = email
	}
	if password != "" {
		payload["password"] = password
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := u.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Users) CreatePHPassUser(userId string, email string, password string, name string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}

	if email != "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}

	if password != "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	url := fmt.Sprintf("/users/phpass")
	payload := make(map[string]interface{})

	payload["userId"] = userId

	if email != "" {
		payload["email"] = email
	}
	if password != "" {
		payload["password"] = password
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := u.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Users) CreateSCryptUser(userId string, email string, password string, passwordSalt string, passwordCpu int, passwordMemory int, passwordParallel int, passwordLength int, name string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}

	if email != "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}

	if password != "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	if passwordSalt != "" {
		return nil, errors.New("Missing required parameter: \"passwordSalt\"")
	}

	if passwordCpu != 0 {
		return nil, errors.New("Missing required parameter: \"passwordCpu\"")
	}

	if passwordMemory != 0 {
		return nil, errors.New("Missing required parameter: \"passwordMemory\"")
	}

	if passwordParallel != 0 {
		return nil, errors.New("Missing required parameter: \"passwordParallel\"")
	}

	if passwordLength != 0 {
		return nil, errors.New("Missing required parameter: \"passwordLength\"")
	}

	url := fmt.Sprintf("/users/scrypt")
	payload := make(map[string]interface{})

	payload["userId"] = userId

	if email != "" {
		payload["email"] = email
	}
	if password != "" {
		payload["password"] = password
	}
	if passwordSalt != "" {
		payload["passwordSalt"] = passwordSalt
	}
	if passwordCpu != 0 {
		payload["passwordCpu"] = passwordCpu
	}
	if passwordMemory != 0 {
		payload["passwordMemory"] = passwordMemory
	}
	if passwordParallel != 0 {
		payload["passwordParallel"] = passwordParallel
	}
	if passwordLength != 0 {
		payload["passwordLength"] = passwordLength
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := u.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Users) createScryptModifiedUser(userId string, email string, password string, passwordSalt string, passwordSaltSeparator string, passwordSignerKey string, name string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}

	if email != "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}

	if password != "" {
		return nil, errors.New("Missing required parameter: \"password\"")
	}

	if passwordSalt != "" {
		return nil, errors.New("Missing required parameter: \"passwordSalt\"")
	}

	if passwordSaltSeparator != "" {
		return nil, errors.New("Missing required parameter: \"passwordSaltSeparator\"")
	}

	if passwordSignerKey != "" {
		return nil, errors.New("Missing required parameter: \"passwordSignerKey\"")
	}

	url := fmt.Sprintf("/users/scrypt-modified")
	payload := make(map[string]interface{})

	payload["userId"] = userId

	if email != "" {
		payload["email"] = email
	}
	if password != "" {
		payload["password"] = password
	}
	if passwordSalt != "" {
		payload["passwordSalt"] = passwordSalt
	}
	if passwordSaltSeparator != "" {
		payload["passwordSaltSeparator"] = passwordSaltSeparator
	}
	if passwordSignerKey != "" {
		payload["passwordSignerKey"] = passwordSignerKey
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := u.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}
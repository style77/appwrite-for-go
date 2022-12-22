package services

import (
	"errors"
	"fmt"

	"github.com/style77/sdk-for-go"
)

type Teams struct {
	client *appwrite.Client
}

func NewTeams(client *appwrite.Client) *Teams {
	return &Teams{client: client}
}

func (t *Teams) List(queries []string, search string) (interface{}, error) {
	url := fmt.Sprintf("/teams")
	payload := make(map[string]interface{})

	if len(queries) > 0 {
		payload["queries"] = queries
	}
	if search != "" {
		payload["search"] = search
	}

	res, err := t.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) Create(teamId string, name string, roles []string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}
	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}
	if len(roles) == 0 {
		return nil, errors.New("Missing required parameter: \"roles\"")
	}

	url := fmt.Sprintf("/teams")
	payload := make(map[string]interface{})

	payload["teamId"] = teamId
	payload["name"] = name
	payload["roles"] = roles

	res, err := t.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) Get(teamId string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}

	url := fmt.Sprintf("/teams/%s", teamId)

	res, err := t.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) Update(teamId string, name string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}
	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}

	url := fmt.Sprintf("/teams/%s", teamId)
	payload := make(map[string]interface{})

	payload["name"] = name

	res, err := t.client.Call("put", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) Delete(teamId string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}

	url := fmt.Sprintf("/teams/%s", teamId)

	res, err := t.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) ListMemberships(teamId string, queries []string, search string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}

	url := fmt.Sprintf("/teams/%s/memberships", teamId)
	payload := make(map[string]interface{})

	if len(queries) > 0 {
		payload["queries"] = queries
	}
	if search != "" {
		payload["search"] = search
	}

	res, err := t.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "array")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) CreateMembership(teamId string, email string, roles []string, path string, name string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}
	if email == "" {
		return nil, errors.New("Missing required parameter: \"email\"")
	}
	if len(roles) == 0 {
		return nil, errors.New("Missing required parameter: \"roles\"")
	}

	url := fmt.Sprintf("/teams/%s/memberships", teamId)
	payload := make(map[string]interface{})

	payload["email"] = email
	payload["roles"] = roles
	if path != "" {
		payload["url"] = path
	}
	if name != "" {
		payload["name"] = name
	}

	res, err := t.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) GetMembership(teamId string, membershipId string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}
	if membershipId == "" {
		return nil, errors.New("Missing required parameter: \"membershipId\"")
	}

	url := fmt.Sprintf("/teams/%s/memberships/%s", teamId, membershipId)

	res, err := t.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) UpdateMembership(teamId string, membershipId string, roles []string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}
	if membershipId == "" {
		return nil, errors.New("Missing required parameter: \"membershipId\"")
	}
	if len(roles) == 0 {
		return nil, errors.New("Missing required parameter: \"roles\"")
	}

	url := fmt.Sprintf("/teams/%s/memberships/%s", teamId, membershipId)
	payload := make(map[string]interface{})

	payload["roles"] = roles

	res, err := t.client.Call("put", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) DeleteMembership(teamId string, membershipId string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}
	if membershipId == "" {
		return nil, errors.New("Missing required parameter: \"membershipId\"")
	}

	url := fmt.Sprintf("/teams/%s/memberships/%s", teamId, membershipId)

	res, err := t.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Teams) UpdateMembershipStatus(teamId string, membershipId string, userId string, secret string) (interface{}, error) {
	if teamId == "" {
		return nil, errors.New("Missing required parameter: \"teamId\"")
	}
	if membershipId == "" {
		return nil, errors.New("Missing required parameter: \"membershipId\"")
	}
	if userId == "" {
		return nil, errors.New("Missing required parameter: \"userId\"")
	}
	if secret == "" {
		return nil, errors.New("Missing required parameter: \"secret\"")
	}

	url := fmt.Sprintf("/teams/%s/memberships/%s", teamId, membershipId)
	payload := make(map[string]interface{})

	payload["userId"] = userId
	payload["secret"] = secret

	res, err := t.client.Call("put", url, map[string]string{
		"content-type": "application/json",
	}, payload, "object")
	if err != nil {
		return nil, err
	}

	return res, nil
}
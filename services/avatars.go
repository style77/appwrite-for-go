package services

import (
	"errors"
	"fmt"

	// "io/ioutil"
	// "net/http"
	"strconv"

	"github.com/style77/appwrite-for-go"
)

type Avatars struct {
	client *appwrite.Client
}

func NewAvatars(client *appwrite.Client) *Avatars {
	return &Avatars{client: client}
}

func (a *Avatars) GetBrowser(code string, width, height, quality int) (interface{}, error) {
	if code == "" {
		return nil, errors.New("Missing required parameter: \"code\"")
	}

	url := fmt.Sprintf("/avatars/browsers/%s", code)
	payload := make(map[string]interface{})

	if width != 0 {
		payload["width"] = strconv.Itoa(width)
	}
	if height != 0 {
		payload["height"] = strconv.Itoa(height)
	}
	if quality != 0 {
		payload["quality"] = strconv.Itoa(quality)
	}

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Avatars) GetCreditCard(code string, width int, height int, quality int) (interface{}, error) {
	if code == "" {
		return nil, errors.New("Missing required parameter: \"code\"")
	}

	url := fmt.Sprintf("/avatars/credit-cards/%s", code)
	payload := make(map[string]interface{})

	if width != 0 {
		payload["width"] = strconv.Itoa(width)
	}
	if height != 0 {
		payload["height"] = strconv.Itoa(height)
	}
	if quality != 0 {
		payload["quality"] = strconv.Itoa(quality)
	}

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Avatars) GetFlag(code string, width int, height int, quality int) (interface{}, error) {
	if code == "" {
		return nil, errors.New("Missing required parameter: \"code\"")
	}

	url := fmt.Sprintf("/avatars/flags/%s", code)
	payload := make(map[string]interface{})

	if width != 0 {
		payload["width"] = strconv.Itoa(width)
	}
	if height != 0 {
		payload["height"] = strconv.Itoa(height)
	}
	if quality != 0 {
		payload["quality"] = strconv.Itoa(quality)
	}

	res, err := a.client.Call("get", url, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Avatars) GetFavicon(url string) (interface{}, error) {
	if url == "" {
		return nil, errors.New("Missing required parameter: \"url\"")
	}

	path := "/avatars/favicon"
	payload := make(map[string]interface{})

	payload["url"] = url

	res, err := a.client.Call("get", path, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Avatars) GetImage(url string, width int, height int, quality int) (interface{}, error) {
	if url == "" {
		return nil, errors.New("Missing required parameter: \"url\"")
	}

	path := "/avatars/image"
	payload := make(map[string]interface{})

	payload["url"] = url

	if width != 0 {
		payload["width"] = strconv.Itoa(width)
	}
	if height != 0 {
		payload["height"] = strconv.Itoa(height)
	}
	if quality != 0 {
		payload["quality"] = strconv.Itoa(quality)
	}

	res, err := a.client.Call("get", path, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Avatars) GetInitials(name string, width int, height int, background string) (interface{}, error) {
	path := "/avatars/initials"
	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	}
	if width != 0 {
		payload["width"] = strconv.Itoa(width)
	}
	if height != 0 {
		payload["height"] = strconv.Itoa(height)
	}
	if background != "" {
		payload["background"] = background
	}

	res, err := a.client.Call("get", path, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Avatars) GetQR(text string, size int, margin int, download bool) (interface{}, error) {
	if text == "" {
		return nil, errors.New("Missing required parameter: \"text\"")
	}

	path := "/avatars/qr"
	payload := make(map[string]interface{})

	payload["text"] = text

	if size != 0 {
		payload["size"] = strconv.Itoa(size)
	}
	if margin != 0 {
		payload["margin"] = strconv.Itoa(margin)
	}
	if download != false {
		payload["download"] = strconv.FormatBool(download)
	}

	res, err := a.client.Call("get", path, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}
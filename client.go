package appwrite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const chunkSize = 5 * 1024 * 1024 // 5MB

type Client struct {
	endpoint       string
	headers        map[string]string
	selfSigned     bool
	responseFormat string
}

func NewClient() *Client {
	return &Client{
		endpoint: "https://HOSTNAME/v1",
		headers: map[string]string{
			"content-type":      "",
			"x-sdk-name":        "Go-SDK",
			"x-sdk-platform":    "server",
			"x-sdk-language":    "golang",
			"x-sdk-version":     "1.18",
			"X-Appwrite-Format": "1.0.0",
		},
		selfSigned: false,
	}
}

// SetProject sets the project ID.
func (c *Client) SetProject(project string) *Client {
	c.AddHeader("X-Appwrite-Project", project)
	return c
}

// SetKey sets the secret API key.
func (c *Client) SetKey(key string) *Client {
	c.AddHeader("X-Appwrite-Key", key)
	return c
}

// SetJWT sets the secret JSON Web Token.
func (c *Client) SetJWT(jwt string) *Client {
	c.AddHeader("X-Appwrite-JWT", jwt)
	return c
}

// SetLocale sets the locale.
func (c *Client) SetLocale(locale string) *Client {
	c.AddHeader("X-Appwrite-Locale", locale)
	return c
}

// SetSelfSigned allows self-signed requests.
func (c *Client) SetSelfSigned(status bool) *Client {
	c.selfSigned = status
	return c
}

// SetEndpoint sets the API endpoint.
func (c *Client) SetEndpoint(endpoint string) *Client {
	c.endpoint = endpoint
	return c
}

// AddHeader adds a custom header.
func (c *Client) AddHeader(key, value string) *Client {
	c.headers[key] = value
	return c
}

func (c *Client) Call(method, path string, headers map[string]string, params map[string]interface{}, responseType string) (interface{}, error) {
	if c.selfSigned {
		os.Setenv("NODE_TLS_REJECT_UNAUTHORIZED", "0") // idk if that should be node_tls_... or can be just tls_...
	}

	for key, value := range c.headers {
		headers[key] = value
	}

	contentType := strings.ToLower(headers["content-type"])

	var formData *bytes.Buffer
	if strings.HasPrefix(contentType, "multipart/form-data") {
		form := &bytes.Buffer{}
		writer := multipart.NewWriter(form)

		flatParams := flatten(params, "")

		for key, value := range flatParams {
			if value == nil {
				continue
			}

			switch value := value.(type) {
			case map[string]interface{}:
				if value["type"] == "file" {
					part, err := writer.CreateFormFile(key, value["filename"].(string))
					if err != nil {
						return nil, err
					}

					file, err := os.Open(value["file"].(string))
					if err != nil {
						return nil, err
					}
					defer file.Close()

					_, err = io.Copy(part, file)
					if err != nil {
						return nil, err
					}
				} else {
					_ = writer.WriteField(key, fmt.Sprint(value))
				}
			default:
				_ = writer.WriteField(key, value.(string))
			}
		}

		err := writer.Close()
		if err != nil {
			return nil, err
		}

		headers["Content-Type"] = writer.FormDataContentType()
		formData = form
	} else if strings.HasPrefix(contentType, "application/json") {
		formData = bytes.NewBufferString(stringify(params))
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		formData = bytes.NewBufferString(urlencode(params))
	}

	u, err := url.Parse(c.endpoint + path)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	for key, value := range params {
		value, ok := value.(string)
		if !ok {
			// handle the error
		}
		query.Set(key, value)
	}
	u.RawQuery = query.Encode()

	var req *http.Request

	method = strings.ToUpper(method)

	if formData != nil {
		req, err = http.NewRequestWithContext(context.Background(), method, u.String(), formData)
	} else {
		req, err = http.NewRequestWithContext(context.Background(), method, u.String(), http.NoBody)
	}

	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key,
			value,
		)
	}

	switch contentType {
	case "application/json":
		req.Header.Set("Content-Type", "application/json")
	case "application/x-www-form-urlencoded":
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case "multipart/form-data":
		req.Header.Set("Content-Type", "multipart/form-data")
	default:
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch responseType {
	case "file":
		return resp.Body, nil
	case "stream":
		return resp.Body, nil
	case "json":
		jsonData := make(map[string]interface{})
		err := json.NewDecoder(resp.Body).Decode(&jsonData)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	case "arraybuffer":
		return resp.Body, nil
	case "blob":
		return resp.Body, nil
	default:
		return nil, fmt.Errorf("Invalid response type")
	}
}

func stringify(data map[string]interface{}) string {
	output, _ := json.Marshal(data)
	return string(output)
}

func urlencode(data map[string]interface{}) string {
	output := url.Values{}
	for key, value := range data {
		output.Set(key, value.(string))
	}
	return output.Encode()
}

func flatten(data map[string]interface{}, prefix string) map[string]interface{} {
	output := make(map[string]interface{})

	for key, value := range data {
		finalKey := prefix
		if prefix != "" {
			finalKey = prefix + "[" + key + "]"
		} else {
			finalKey = key
		}

		if arr, ok := value.([]interface{}); ok {
			output = mergeMaps(output, flatten(convertSliceToMap(arr), finalKey))
		} else {
			output[finalKey] = value
		}
	}

	return output
}

func convertSliceToMap(slice []interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i, val := range slice {
		m[strconv.Itoa(i)] = val
	}
	return m
}

func mergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

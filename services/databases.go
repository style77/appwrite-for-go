package services

import (
	"errors"
	"fmt"

	"github.com/style77/sdk-for-go"
)

type Databases struct {
	client *appwrite.Client
}

func NewDatabases(client *appwrite.Client) *Databases {
	return &Databases{client: client}
}

// List all the project databases.
func (d *Databases) List() (interface{}, error) {
	url := fmt.Sprintf("/databases")
	payload := make(map[string]interface{})

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "multipart/form-data",
	}, payload, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new database.
func (d *Databases) Create(databaseId string, name string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}

	url := fmt.Sprintf("/databases")
	payload := make(map[string]interface{})
	payload["databaseId"] = databaseId
	payload["name"] = name

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get a database by its unique ID.
func (d *Databases) Get(databaseId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	url := fmt.Sprintf("/databases/%s", databaseId)

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "multipart/form-data",
	}, nil, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update a database by its unique ID.
func (d *Databases) Update(databaseId string, name string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}

	url := fmt.Sprintf("/databases/%s", databaseId)
	payload := make(map[string]interface{})
	payload["name"] = name

	res, err := d.client.Call("put", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete a database by its unique ID.
func (d *Databases) Delete(databaseId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	url := fmt.Sprintf("/databases/%s", databaseId)

	res, err := d.client.Call("delete", url, map[string]string{
		"content-type": "multipart/form-data",
	}, nil, "arraybuffer")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// List all the collections for a given database.
func (d *Databases) ListCollections(databaseId string, queries []string, search string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections", databaseId)
	payload := make(map[string]interface{})

	if len(queries) > 0 {
		payload["queries"] = queries
	}

	if search != "" {
		payload["search"] = search
	}

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new collection.
func (d *Databases) CreateCollection(databaseId string, collectionId string, name string, permissions []string, documentSecurity *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if name == "" {

		return nil, errors.New("Missing required parameter: \"name\"")
	}

	url := fmt.Sprintf("/databases/%s/collections", databaseId)
	payload := make(map[string]interface{})
	payload["collectionId"] = collectionId
	payload["name"] = name

	if len(permissions) > 0 {
		payload["permissions"] = permissions
	}

	if documentSecurity != nil {
		payload["documentSecurity"] = documentSecurity
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get a collection by its unique ID.
func (d *Databases) GetCollection(databaseId string, collectionId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s", databaseId, collectionId)

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update a collection by its unique ID.
func (d *Databases) UpdateCollection(databaseId string, collectionId string, name string, permissions []string, documentSecurity *bool, enabled *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if name == "" {
		return nil, errors.New("Missing required parameter: \"name\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["name"] = name

	if len(permissions) > 0 {
		payload["permissions"] = permissions
	}

	if documentSecurity != nil {
		payload["documentSecurity"] = documentSecurity
	}

	if enabled != nil {
		payload["enabled"] = enabled
	}

	res, err := d.client.Call("put", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete a collection by its unique ID.
func (d *Databases) DeleteCollection(databaseId string, collectionId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s", databaseId, collectionId)

	res, err := d.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// List all attributes for a given collection.
func (d *Databases) ListAttributes(databaseId string, collectionId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes", databaseId, collectionId)

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new boolean attribute.
func (d *Databases) CreateBooleanAttribute(databaseId string, collectionId string, key string, required *bool, xdefault *bool, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/boolean", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new datetime attribute.
func (d *Databases) CreateDatetimeAttribute(databaseId string, collectionId string, key string, required *bool, xdefault *string, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/datetime", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new Email attribute.
func (d *Databases) CreateEmailAttribute(databaseId string, collectionId string, key string, required *bool, xdefault *string, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/email", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new Enum attribute.
func (d *Databases) CreateEnumAttribute(databaseId string, collectionId string, key string, values []string, required *bool, xdefault *string, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if values == nil {
		return nil, errors.New("Missing required parameter: \"values\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/enum", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key
	payload["values"] = values

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new Float attribute.
func (d *Databases) CreateFloatAttribute(databaseId string, collectionId string, key string, required *bool, min *int, max *int, xdefault *float64, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/float", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	if min != nil {
		payload["min"] = min
	}

	if max != nil {
		payload["max"] = max
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new Integer attribute.
func (d *Databases) CreateIntegerAttribute(databaseId string, collectionId string, key string, required *bool, min *int, max *int, xdefault *int, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/integer", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	if min != nil {
		payload["min"] = min
	}

	if max != nil {
		payload["max"] = max
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new IP attribute.
func (d *Databases) CreateIpAttribute(databaseId string, collectionId string, key string, required *bool, xdefault *string, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/ip", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new String attribute.
func (d *Databases) CreateStringAttribute(databaseId string, collectionId string, key string, size int, required *bool, xdefault *string, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	if size == 0 {
		return nil, errors.New("Missing required parameter: \"size\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/string", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key
	payload["size"] = size

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new URL attribute.
func (d *Databases) CreateUrlAttribute(databaseId string, collectionId string, key string, required *bool, xdefault *string, array *bool) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if key == "" {
		return nil, errors.New("Missing required parameter: \"key\"")
	}

	if required == nil {
		return nil, errors.New("Missing required parameter: \"required\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/url", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["key"] = key

	if xdefault != nil {
		payload["default"] = xdefault
	}

	if array != nil {
		payload["array"] = array
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Attribute
func (d *Databases) GetAttribute(databaseId string, collectionId string, attributeId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if attributeId == "" {
		return nil, errors.New("Missing required parameter: \"attributeId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/%s", databaseId, collectionId, attributeId)

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete Attribute
func (d *Databases) DeleteAttribute(databaseId string, collectionId string, attributeId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if attributeId == "" {
		return nil, errors.New("Missing required parameter: \"attributeId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/attributes/%s", databaseId, collectionId, attributeId)

	res, err := d.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// List documents
func (d *Databases) ListDocuments(databaseId string, collectionId string, queries []string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/documents", databaseId, collectionId)
	payload := make(map[string]interface{})

	if queries != nil {
		payload["queries"] = queries
	}

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new document
func (d *Databases) CreateDocument(databaseId string, collectionId string, documentId string, data map[string]interface{}, permissions []string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if documentId == "" {
		return nil, errors.New("Missing required parameter: \"documentId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/documents", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["documentId"] = documentId
	payload["data"] = data

	if permissions != nil {
		payload["permissions"] = permissions
	}

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get document
func (d *Databases) GetDocument(databaseId string, collectionId string, documentId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if documentId == "" {
		return nil, errors.New("Missing required parameter: \"documentId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/documents/%s", databaseId, collectionId, documentId)

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update document
func (d *Databases) UpdateDocument(databaseId string, collectionId string, documentId string, data map[string]interface{}, permissions []string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if documentId == "" {
		return nil, errors.New("Missing required parameter: \"documentId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/documents/%s", databaseId, collectionId, documentId)
	payload := make(map[string]interface{})
	payload["data"] = data

	if permissions != nil {
		payload["permissions"] = permissions
	}

	res, err := d.client.Call("put", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete document
func (d *Databases) DeleteDocument(databaseId string, collectionId string, documentId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if documentId == "" {
		return nil, errors.New("Missing required parameter: \"documentId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/documents/%s", databaseId, collectionId, documentId)

	res, err := d.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// List indexes
func (d *Databases) ListIndexes(databaseId string, collectionId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/indexes", databaseId, collectionId)

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create Index
func (d *Databases) CreateIndex(databaseId string, collectionId string, indexId string, data map[string]interface{}) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if indexId == "" {
		return nil, errors.New("Missing required parameter: \"indexId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/indexes", databaseId, collectionId)
	payload := make(map[string]interface{})
	payload["indexId"] = indexId
	payload["data"] = data

	res, err := d.client.Call("post", url, map[string]string{
		"content-type": "application/json",
	}, payload, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get Index
func (d *Databases) GetIndex(databaseId string, collectionId string, indexId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if indexId == "" {
		return nil, errors.New("Missing required parameter: \"indexId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/indexes/%s", databaseId, collectionId, indexId)

	res, err := d.client.Call("get", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete Index
func (d *Databases) DeleteIndex(databaseId string, collectionId string, indexId string) (interface{}, error) {
	if databaseId == "" {
		return nil, errors.New("Missing required parameter: \"databaseId\"")
	}

	if collectionId == "" {
		return nil, errors.New("Missing required parameter: \"collectionId\"")
	}

	if indexId == "" {
		return nil, errors.New("Missing required parameter: \"indexId\"")
	}

	url := fmt.Sprintf("/databases/%s/collections/%s/indexes/%s", databaseId, collectionId, indexId)

	res, err := d.client.Call("delete", url, map[string]string{
		"content-type": "application/json",
	}, nil, "json")
	if err != nil {
		return nil, err
	}

	return res, nil
}
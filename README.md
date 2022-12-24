# Appwrite Golang SDK
SDK for Appwrite in Golang (Go), a simple and secure API for your database, storage, functions, and more.

## Getting Started

### Installation

```bash
go get github.com/Style77/sdk-for-go
```

### Usage

Simple usage for listing all the documents in a collection:
```go
package main

import (
    "fmt"
    "github.com/Style77/sdk-for-go"
)

func main() {
    client := appwrite.NewClient()
    client.SetEndpoint("http://localhost/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    // Call any of the Appwrite services: for example, database
    response, err := client.Database.ListDocuments("databaseId", "collectionId", nil) // List all the documents in the collection

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(response)
}
```
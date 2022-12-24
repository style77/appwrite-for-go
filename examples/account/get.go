package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/style77/appwrite-for-go"          // used for the appwrite.Client
	"github.com/style77/appwrite-for-go/services" // used for the services.Account
)

func main() {
	godotenv.Load(".env")

	fmt.Println(os.Getenv("APPWRITE_PROJECT_ID"))

	client := appwrite.NewClient()
	client.SetEndpoint(os.Getenv("APPWRITE_ENDPOINT")) // Your API Endpoint
	client.SetProject(os.Getenv("APPWRITE_PROJECT_ID")) // Your project ID
	client.SetKey(os.Getenv("APPWRITE_API_KEY")) // Your secret API key

	account := services.NewAccount(client)

	response, err := account.Get()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
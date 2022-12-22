package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/style77/sdk-for-go"          // used for the appwrite.Client
	"github.com/style77/sdk-for-go/services" // used for the services.Avatars
)

func main() {
	godotenv.Load("C:\\Users\\style\appwrite\\examples\\.env.local")

	client := appwrite.NewClient()
	client.SetEndpoint(os.Getenv("APPWRITE_ENDPOINT")) // Your API Endpoint
	client.SetProject(os.Getenv("APPWRITE_PROJECT_ID")) // Your project ID
	client.SetKey(os.Getenv("APPWRITE_API_KEY")) // Your secret API key

	avatars := services.NewAvatars(client)

	response, err := avatars.GetCreditCard("visa", 0, 0, 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
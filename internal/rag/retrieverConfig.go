package rag

import (
	"context"
	"log"
	"os"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
	"github.com/weaviate/weaviate-go-client/v5/weaviate/auth"
)

var WEAVIATE_URL = os.Getenv("WEAVIATE_RESTENDPOINT")
var WEAVIATE_APIKEY = os.Getenv("WEAVIATE_APIKEY")
var HUGGINGFACE_APIKEY = os.Getenv("HUGGINGFACE_APIKEY")
var GEMINI_APIKEY = os.Getenv("GEMINI_APIKEY")

const GEN_MODEL string = "gemini-2.5-flash"

func ConnectToClient() (*weaviate.Client, error) {
	cfg := weaviate.Config{
		Host:   WEAVIATE_URL,
		Scheme: "https",
		Headers: map[string]string{
			"X-HuggingFace-Api-Key": HUGGINGFACE_APIKEY,
		},
		AuthConfig: auth.ApiKey{Value: WEAVIATE_APIKEY},
	}

	client, err := weaviate.NewClient(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to client: %v", err)
	}

	ready, err := client.Misc().ReadyChecker().Do(context.Background())

	if err != nil {
		panic(err)
	}

	log.Printf("%v", ready)
	return client, nil
}

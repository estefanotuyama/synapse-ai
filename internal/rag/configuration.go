package rag

import (
	"context"
	"log"
	"os"

	"github.com/weaviate/weaviate/entities/models"
)

var WEAVIATE_URL = os.Getenv("WEAVIATE_RESTENDPOINT")
var WEAVIATE_APIKEY = os.Getenv("WEAVIATE_APIKEY")
var HUGGINGFACE_APIKEY = os.Getenv("HUGGINGFACE_APIKEY")
var GEMINI_APIKEY = os.Getenv("GEMINI_APIKEY")

const COLLECTION_NAME = "DocumentStore"
const GEN_MODEL string = "gemini-2.5-flash"

const SYSTEM_PROMPT = `You are a specialized agent named SynapseAI. A user with the name %s has created you,
	naming you %s and has provided the following as your purpose: %s .
	The user has uploaded documents that will be retrieved for you with RAG or you to answer
	user prompts. Here are the names of the documents, just so you know: %s .

	Important: Every user prompt will come in this format:
	{UserPrompt: 'user prompt wil be inside these quotes'}
	{Documents: 'chunks of documents retrieved by the retriever will be inside the quotes'}
`

func SetupWeaviateCollection() error {
	ctx := context.Background()
	client, err := ConnectToVectorDB()

	if err != nil {
		return err
	}

	exists, err := client.Schema().ClassExistenceChecker().WithClassName(COLLECTION_NAME).Do(ctx)

	if err != nil {
		return err
	} else if exists {
		log.Println("Found collection. Starting server.")
		return nil
	}

	log.Println("Couldn't find collection. Creating it now.")

	class := &models.Class{
		Class: COLLECTION_NAME,
		Properties: []*models.Property{
			{
				Name:     "content",
				DataType: []string{"text"},
			},
		},

		VectorConfig: map[string]models.VectorConfig{
			"content_vector": {
				VectorIndexType: "hnsw",
				Vectorizer: map[string]interface{}{
					"text2vec-huggingface": map[string]interface{}{
						"properties":         []string{"content"},
						"model":              "sentence-transformers/all-MiniLM-L6-v2",
						"vectorizeClassName": false,
					},
				},
			},
		},

		MultiTenancyConfig: &models.MultiTenancyConfig{
			Enabled:              true,
			AutoTenantActivation: true,
			AutoTenantCreation:   true,
		},
	}

	err = client.Schema().ClassCreator().WithClass(class).Do(ctx)

	if err != nil {
		return err
	}

	log.Println("Created collection. Starting server.")
	return nil
}

package rag

import (
	"context"
	"fmt"
	"time"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

// defines how many embedding retries are allowed
const MAX_RETRIES int = 3

type Collection struct {
	collectionName string
	docs           []map[string]string
}

func (c *Collection) Create(client *weaviate.Client) error {
	ctx := context.Background()

	collection := &models.Class{
		Class: c.collectionName,

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
	}

	err := client.Schema().ClassCreator().WithClass(collection).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Collection) AddDocuments(client *weaviate.Client) error {

	ctx := context.Background()
	batcher := client.Batch().ObjectsBatcher()

	for _, doc := range c.docs {
		batcher.WithObjects(&models.Object{
			Class: c.collectionName,
			Properties: map[string]interface{}{
				"content": doc["content"],
			},
		})
	}

	var err error
	var batchRes []models.ObjectsGetResponse

	for i := 0; i < MAX_RETRIES; i++ {
		batchRes, err = batcher.Do(ctx)
		if err == nil {
			break
		}

		fmt.Printf("Attempt %d failed. Retrying in 2 seconds... Error: %v\n", i+1, err)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("batch import failed after %d retries: %w", MAX_RETRIES, err)
	}

	// handle object-level errors
	var allErrors []string
	for _, res := range batchRes {
		if res.Result != nil && res.Result.Errors != nil {
			for _, errItem := range res.Result.Errors.Error {
				if errItem != nil {
					allErrors = append(allErrors, errItem.Message)
				}
			}
		}
	}

	if len(allErrors) > 0 {
		return fmt.Errorf("encountered %d errors during object import: %v", len(allErrors), allErrors)
	}

	return nil
}

func (c *Collection) DeleteCollection(client *weaviate.Client) error {
	ctx := context.Background()

	if err := client.Schema().ClassDeleter().WithClassName(c.collectionName).Do(ctx); err != nil {
		return fmt.Errorf("Error while deleting collection: %v", err)
	}
	return nil
}

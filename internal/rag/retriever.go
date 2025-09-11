package rag

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
	"github.com/weaviate/weaviate-go-client/v5/weaviate/auth"
	"github.com/weaviate/weaviate/entities/models"
)

// defines how many embedding retries are allowed
const MAX_RETRIES int = 3

type AgentTenant struct {
	name string
	docs []map[string]string
}

func (t *AgentTenant) Create(client *weaviate.Client) error {
	ctx := context.Background()

	err := client.Schema().TenantsCreator().
		WithClassName(COLLECTION_NAME).
		WithTenants(models.Tenant{Name: t.name}).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (t *AgentTenant) AddDocuments(client *weaviate.Client) error {
	ctx := context.Background()

	const BATCH_SIZE = 200
	for start := 0; start < len(t.docs); start += BATCH_SIZE {
		end := start + BATCH_SIZE
		if end > len(t.docs) {
			end = len(t.docs)
		}

		batcher := client.Batch().ObjectsBatcher()

		for _, doc := range t.docs[start:end] {
			batcher.WithObjects(&models.Object{
				Class:  COLLECTION_NAME,
				Tenant: t.name,
				Properties: map[string]interface{}{
					"content": doc["content"],
				},
			})
		}

		var err error
		var batchRes []models.ObjectsGetResponse

		for i := range MAX_RETRIES {
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
	}

	return nil
}

func (t *AgentTenant) Delete(client *weaviate.Client) error {
	ctx := context.Background()

	err := client.Schema().TenantsDeleter().WithTenants(t.name).Do(ctx)

	if err != nil {
		return fmt.Errorf("Error while deleting tenant: %v", err)
	}
	return nil
}

func ConnectToVectorDB() (*weaviate.Client, error) {
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

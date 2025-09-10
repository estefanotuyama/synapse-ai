package rag

import (
	"testing"
)

func TestRagSystem(t *testing.T) {
	t.Log("Connecting to Weaviate client...")
	client, err := ConnectToVectorDB()
	if err != nil {
		t.Fatalf("Failed to connect to Weaviate client: %v", err)
	}
	t.Log("Weaviate client connection successful.")

	// define the collection
	col := Collection{
		collectionName: "SystemTest",
		docs: []map[string]string{
			{"content": "The secret keyword is 'SIC MUNDUS CREATUS EST'"},
			{"content": "second chunk"},
		},
	}

	t.Cleanup(func() {
		t.Log("Cleaning up: deleting collection...")
		err := col.Delete(client)
		if err != nil {
			t.Errorf("Failed to delete collection during cleanup: %v", err)
		} else {
			t.Log("Cleanup successful.")
		}
	})

	t.Run("creates the collection", func(t *testing.T) {
		err := col.Create(client)
		if err != nil {
			t.Fatalf("Failed to create collection: %v", err)
		}
		t.Log("Collection creation successful.")
	})

	t.Run("adds documents to the collection", func(t *testing.T) {
		err := col.AddDocuments(client)
		if err != nil {
			t.Fatalf("Failed to add documents: %v", err)
		}
		t.Log("Successfully added documents. Secret keyword is SIC MUNDUS CREATUS EST")
	})

	t.Run("tests Rag pipeline", func(t *testing.T) {
		history, err := testRagPipeline(col.collectionName)

		if err != nil {
			t.Fatalf("Rag pipeline failed: %v", err)
		}

		if len(history) != 2 {
			t.Fatalf("Rag system returned wrong message history length.")
		}

		t.Logf("Rag pipeline success. LLM Response to 'secret keyword':%v", history[len(history)-1].Content)
	})
}

func testRagPipeline(collectionName string) ([]*ChatMessage, error) {
	prompt := "What is the secret keyword? Only say the secret keyword."
	msgHistory := []*ChatMessage{}

	responseHistory, err := CallRagSystem(prompt, msgHistory, collectionName)

	return responseHistory, err
}

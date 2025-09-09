package rag

import (
	"testing"
)

func TestCollectionLifecycle(t *testing.T) {
	t.Log("Connecting to Weaviate client...")
	client, err := ConnectToClient()
	if err != nil {
		t.Fatalf("Failed to connect to Weaviate client: %v", err)
	}
	t.Log("Weaviate client connection successful.")

	// define the collection
	col := Collection{
		collectionName: "TestLifecycleCollection",
		docs: []map[string]string{
			{"content": "first chunk"},
			{"content": "second chunk"},
		},
	}

	t.Cleanup(func() {
		t.Log("Cleaning up: deleting collection...")
		err := col.DeleteCollection(client)
		if err != nil {
			t.Errorf("Failed to delete collection during cleanup: %v", err)
		} else {
			t.Log("Cleanup successful.")
		}
	})

	// create collection test
	t.Run("creates the collection", func(t *testing.T) {
		err := col.Create(client)
		if err != nil {
			t.Fatalf("Failed to create collection: %v", err)
		}
		t.Log("Collection creation successful.")
	})

	// add documents test
	t.Run("adds documents to the collection", func(t *testing.T) {
		err := col.AddDocuments(client)
		if err != nil {
			t.Fatalf("Failed to add documents: %v", err)
		}
		t.Log("Successfully added documents.")
	})

}

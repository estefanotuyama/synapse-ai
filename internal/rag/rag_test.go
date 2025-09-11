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

	// define the tenant
	ten := AgentTenant{
		name: "TenantTest",
		docs: []map[string]string{
			{"content": "The secret keyword is 'SIC MUNDUS CREATUS EST'"},
			{"content": "second chunk"},
		},
	}

	t.Cleanup(func() {
		t.Log("Cleaning up: deleting tenant...")
		err := ten.Delete(client)
		if err != nil {
			t.Errorf("Failed to delete tenant during cleanup: %v", err)
		} else {
			t.Log("Cleanup successful.")
		}
	})

	t.Run("creates tenant", func(t *testing.T) {
		err := ten.Create(client)
		if err != nil {
			t.Fatalf("Failed to create tenant: %v", err)
		}
		t.Log("Tenant creation successful.")
	})

	t.Run("adds documents to tenant", func(t *testing.T) {
		err := ten.AddDocuments(client)
		if err != nil {
			t.Fatalf("Failed to add documents: %v", err)
		}
		t.Log("Successfully added documents. Secret keyword is SIC MUNDUS CREATUS EST")
	})

	t.Run("tests Rag pipeline", func(t *testing.T) {
		history, err := testRagPipeline(ten.name)

		if err != nil {
			t.Fatalf("Rag pipeline failed: %v", err)
		}

		if len(history) != 2 {
			t.Fatalf("Rag system returned wrong message history length.")
		}

		t.Logf("Rag pipeline success. LLM Response to 'secret keyword':%v", history[len(history)-1].Content)
	})
}

func testRagPipeline(tenantName string) ([]*ChatMessage, error) {
	prompt := "What is the secret keyword? Only say the secret keyword."
	msgHistory := []*ChatMessage{}

	responseHistory, err := CallRagSystem(prompt, msgHistory, tenantName)

	return responseHistory, err
}

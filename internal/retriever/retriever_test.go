package retriever

import (
	"testing"
)

func TestConnection(t *testing.T) {
	_, err := ConnectToClient()

	if err == nil {
		t.Log("Connection successful.")
	} else {
		t.Errorf("Failed connection to Weaviate client: %v", err)
	}
}

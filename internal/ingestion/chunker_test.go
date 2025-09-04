package ingestion

import (
	"os"
	"testing"
)

func TestChunkingSmall(t *testing.T) {
	file, err := os.ReadFile("testdata/ChunkingTest.txt")
	if err != nil {
		t.Fatalf("failed to read test file: %v", err)
	}
	document := string(file)
	chunks := ChunkDocument(document)

	if len(chunks) == 0 {
		t.Error("ChunkDocument returned no chunks, expected at least one.")
	}

	/*
		t.Log("Generated chunks:")
		for i, chunk := range chunks {
		t.Logf("------------------------Chunk %d------------------------%v", i, chunk)
		}
	*/

	t.Log("First and last chunk:")
	t.Log("-------------------------FIRST CHUNK-------------------------")
	t.Log(chunks[0])
	t.Log("-------------------------LAST CHUNK-------------------------")
	t.Log(chunks[len(chunks)-1])
}

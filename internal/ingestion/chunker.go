package ingestion

const CHUNK_SIZE int = 1000
const OVERLAP_SIZE int = 100

func ChunkDocument(doc string) []map[string]string {

	chunkCount := (len(doc) + CHUNK_SIZE - 1) / CHUNK_SIZE
	chunkList := make([]map[string]string, 0, chunkCount)

	for i := 0; i < len(doc); i += CHUNK_SIZE {

		chunk_start := max(0, i-OVERLAP_SIZE)
		chunk_end := min(len(doc), i+CHUNK_SIZE+OVERLAP_SIZE)

		current := doc[chunk_start:chunk_end]
		currentDict := map[string]string{"content": current}

		chunkList = append(chunkList, currentDict)
	}
	return chunkList
}

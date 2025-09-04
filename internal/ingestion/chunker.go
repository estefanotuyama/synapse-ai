package ingestion

const CHUNK_SIZE int = 500
const OVERLAP_SIZE int = 50

func ChunkDocument(doc string) []string {

	chunkCount := (len(doc) + CHUNK_SIZE - 1) / CHUNK_SIZE
	chunkList := make([]string, 0, chunkCount)

	for i := 0; i < len(doc); i += CHUNK_SIZE {

		chunk_start := max(0, i-OVERLAP_SIZE)
		chunk_end := min(len(doc), i+CHUNK_SIZE+OVERLAP_SIZE)

		current := doc[chunk_start:chunk_end]
		chunkList = append(chunkList, current)
	}
	return chunkList
}

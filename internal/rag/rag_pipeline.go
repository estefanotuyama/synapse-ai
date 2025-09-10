package rag

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v5/weaviate/graphql"
)

func CallRagSystem(userPrompt string, msgHistory []*ChatMessage, collectionName string) ([]*ChatMessage, error) {
	weaviateClient, err := ConnectToVectorDB()

	if err != nil {
		return msgHistory, err
	}

	ctx := context.Background()

	//how many documents we are retrieving
	limit := 8

	q := weaviateClient.GraphQL().Get().
		WithClassName(collectionName).
		WithFields(graphql.Field{Name: "content"}).
		WithHybrid(weaviateClient.GraphQL().HybridArgumentBuilder().WithQuery(userPrompt)).
		WithLimit(limit)

	result, err := q.Do(ctx)

	if err != nil {
		return msgHistory, err
	}

	documents := result.Data

	llmPrompt := fmt.Sprintf(`{userPrompt: '%s'}
		{documents: '%s'}`, userPrompt, documents)

	responseHistory, err := CallWithContext(llmPrompt, msgHistory)

	if err != nil {
		return msgHistory, err
	}

	return responseHistory, nil
}

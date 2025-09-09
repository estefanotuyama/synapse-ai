package rag

import (
	"context"
	"strings"

	"google.golang.org/genai"
)

type ChatMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

// calls the LLM and returns both the response and the entire chat history.
func CallWithContext(prompt string, msgHistory []*ChatMessage) (string, []*ChatMessage, error) {
	ctx := context.Background()
	client, err := connectToLlm()
	if err != nil {
		return "Error. Could not connect to LLM.", nil, err
	}

	// add user message and convert to genai type for processing
	msgHistory = append(msgHistory, &ChatMessage{Content: prompt, Role: "user"})
	genaiHistory := toGenaiContent(msgHistory)

	res, err := client.Models.GenerateContent(ctx, GEN_MODEL, genaiHistory, nil)
	clear(genaiHistory)

	if err != nil {
		return "Error. Could not generate content.", nil, err
	}

	textResponse := extractText(res)

	fullHistory := append(msgHistory, &ChatMessage{Content: textResponse, Role: "model"})

	return textResponse, fullHistory, nil
}

func connectToLlm() (*genai.Client, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GEMINI_APIKEY,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

// turn a Gemini LLM response into a string.
func extractText(resp *genai.GenerateContentResponse) string {
	var textBuilder strings.Builder
	if resp == nil || len(resp.Candidates) == 0 {
		return ""
	}

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				textBuilder.WriteString(string(part.Text))
			}
		}
	}
	return textBuilder.String()
}

// converts our own message struct into the type accepted by Gemini.
func toGenaiContent(msgHistory []*ChatMessage) []*genai.Content {
	genaiHistory := []*genai.Content{}

	for _, msg := range msgHistory {
		genaiHistory = append(genaiHistory,
			genai.NewContentFromText(msg.Content, genai.Role(msg.Role)),
		)
	}
	return genaiHistory
}

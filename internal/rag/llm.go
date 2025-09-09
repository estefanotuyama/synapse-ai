package rag

import (
	"context"

	"google.golang.org/genai"
)

type ChatMessage struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

func CallWithContext(prompt string, msgHistory []*genai.Content) (any, []*genai.Content, error) {
	ctx := context.Background()

	if msgHistory == nil {
		msgHistory = []*genai.Content{}
	}

	userMessage := genai.NewContentFromText(prompt, genai.RoleUser)
	msgHistory = append(msgHistory, userMessage)

	client, err := connectToLlm()

	if err != nil {
		return nil, nil, err //TODO: need to handle this more gracefully
	}

	res, err := client.Models.GenerateContent(ctx, GEN_MODEL, msgHistory, nil)

	if err != nil {
		return nil, nil, err
	}

	//msgHistory = append(msgHistory, )

	return res, msgHistory, nil
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

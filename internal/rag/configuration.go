package rag

import (
	"os"
)

var WEAVIATE_URL = os.Getenv("WEAVIATE_RESTENDPOINT")
var WEAVIATE_APIKEY = os.Getenv("WEAVIATE_APIKEY")
var HUGGINGFACE_APIKEY = os.Getenv("HUGGINGFACE_APIKEY")
var GEMINI_APIKEY = os.Getenv("GEMINI_APIKEY")

const GEN_MODEL string = "gemini-2.5-flash"

const SYSTEM_PROMPT = `You are a specialized agent named SynapseAI. A user with the name %s has created you,
	naming you %s and has provided the following as your purpose: %s .
	The user has uploaded documents that will be retrieved for you with RAG or you to answer
	user prompts. Here are the names of the documents, just so you know: %s .

	Important: Every user prompt will come in this format:
	{UserPrompt: 'user prompt wil be inside these quotes'}
	{Documents: 'chunks of documents retrieved by the retriever will be inside the quotes'}
`

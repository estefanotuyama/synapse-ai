package rag

import (
	"os"
)

var WEAVIATE_URL = os.Getenv("WEAVIATE_RESTENDPOINT")
var WEAVIATE_APIKEY = os.Getenv("WEAVIATE_APIKEY")
var HUGGINGFACE_APIKEY = os.Getenv("HUGGINGFACE_APIKEY")
var GEMINI_APIKEY = os.Getenv("GEMINI_APIKEY")

const GEN_MODEL string = "gemini-2.5-flash"

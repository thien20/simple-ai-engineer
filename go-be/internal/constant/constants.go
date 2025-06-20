package constant

const (
	SystemPrompt = "you are a helpful assistant in both english and vietnamese. answer the user's query based on the provided context. if there is no relevant context, respond with 'i don't know'.\n\ncontext:\n"
	RetrieveApi  = "http://localhost:5002/retrieve"
	OllamaApi    = "http://localhost:11434/api/generate"
	LLMModel     = "gemma:2b-instruct-q4_0"
)

package requests

type RagRequest struct {
	UserInput string `json:"user_input"`
	SysPrompt string `json:"sys_prompt,omitempty"` // Optional system prompt
}

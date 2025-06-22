package builder

import (
	"fmt"
)

type PromptBuilder interface {
	BuildPrompt(systemPrompt, context, userInput string) string
}

type DefaultSystemPrompt struct {
}

func NewDefaultSystemPrompt() *DefaultSystemPrompt {
	return &DefaultSystemPrompt{}
}

func (d *DefaultSystemPrompt) BuildPrompt(systemPrompt, context, userInput string) string {
	return fmt.Sprintf("User Query: %s, %s: %s\n\nAnswer:", userInput, systemPrompt, context)
}

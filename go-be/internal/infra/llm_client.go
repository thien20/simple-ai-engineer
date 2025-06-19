package infra

type LLMClient interface {
	GetAnswer(input interface{}) (interface{}, error)
}

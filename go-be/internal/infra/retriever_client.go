package infra

type RetrieverClient interface {
	Retrieve(input interface{}) (interface{}, error)
}

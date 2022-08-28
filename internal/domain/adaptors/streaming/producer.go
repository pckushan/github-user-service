package streaming

type Producer interface {
	Produce(message interface{}) error
}

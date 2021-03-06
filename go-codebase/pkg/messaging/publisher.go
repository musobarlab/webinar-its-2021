package messaging

// Publisher interface, publisher interface abstraction
type Publisher interface {
	Publish(string, string, []byte) error
}

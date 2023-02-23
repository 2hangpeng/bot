package bot

// Sender Define an interface for a service that can greet someone.
type Sender interface {
	Send(message interface{}) error
}

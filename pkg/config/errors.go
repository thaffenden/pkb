package config

const (
	ErrUnmashallingJSON Error = iota
)

type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrUnmashallingJSON:
		return "error unmarshalling JSON config file"

	default:
		return "unknown error"
	}
}

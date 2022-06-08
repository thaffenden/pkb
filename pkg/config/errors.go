package config

const (
	ErrUnmashallingJSON Error = iota
	ErrConfigNotFound
)

type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrConfigNotFound:
		return "no config file found"

	case ErrUnmashallingJSON:
		return "error unmarshalling JSON config file"

	default:
		return "unknown error"
	}
}

package config

const (
	ErrConfigNotFound Error = iota
	ErrReadingConfigFile
	ErrUnmashallingJSON
)

type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrConfigNotFound:
		return "no config file found"

	case ErrReadingConfigFile:
		return "error reading config file"

	case ErrUnmashallingJSON:
		return "error unmarshalling JSON config file"

	default:
		return "unknown error"
	}
}

package converters

type DeepgramOptions struct {
	LineLength int
}

type DeepgramOption func(*DeepgramOptions)

func WithLineLength(length int) DeepgramOption {
	return func(o *DeepgramOptions) {
		o.LineLength = length
	}
}

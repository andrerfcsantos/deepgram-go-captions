package converters

type LineOptions struct {
	LineLength int
}

type LineOption func(*LineOptions)

func WithLineLength(length int) LineOption {
	return func(o *LineOptions) {
		o.LineLength = length
	}
}

type Converter interface {
	Lines(options ...LineOption) ([][]TimedWord, error)
}

type HeaderConverter interface {
	Converter
	Headers() []string
}

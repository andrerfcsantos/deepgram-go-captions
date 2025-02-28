package converters

type BasicWorderOption func(*BasicWorder)

func WithLines(lines [][]TimedWord) BasicWorderOption {
	return func(b *BasicWorder) {
		b.lines = lines
	}
}

func WithLinesFunc(linesFunc func() [][]TimedWord) BasicWorderOption {
	return func(b *BasicWorder) {
		b.linesFunc = linesFunc
	}
}

func WithHeaders(headers []string) BasicWorderOption {
	return func(b *BasicWorder) {
		b.headers = headers
	}
}

func WithHeadersFunc(headersFunc func() []string) BasicWorderOption {
	return func(b *BasicWorder) {
		b.headersFunc = headersFunc
	}
}

func NewBasicWorder(options ...BasicWorderOption) *BasicWorder {
	b := &BasicWorder{}

	for _, option := range options {
		option(b)
	}

	return b
}

type BasicWorder struct {
	lines       [][]TimedWord
	linesFunc   func() [][]TimedWord
	headers     []string
	headersFunc func() []string
}

func (w *BasicWorder) Lines() [][]TimedWord {
	if w.linesFunc != nil {
		return w.linesFunc()
	}

	return w.lines
}

func (w *BasicWorder) Headers() []string {
	if w.headersFunc != nil {
		return w.headersFunc()
	}

	return w.headers
}

type Worder interface {
	Lines() [][]TimedWord
}

type HeaderWorder interface {
	Worder
	Headers() []string
}

package converters

type Converter interface {
	Convert() (Worder, error)
}

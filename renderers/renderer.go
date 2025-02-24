package renderers

type Renderer interface {
	Render() (string, error)
}

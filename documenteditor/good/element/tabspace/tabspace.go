package tabspace

type TabSpaceElement struct {
	imagePath string
}

func Init() *TabSpaceElement {
	return &TabSpaceElement{}
}

func (el *TabSpaceElement) Render() string {
	return "\t"
}

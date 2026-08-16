package text

type TextElement struct {
	text string
}

func Init(text string) *TextElement {
	return &TextElement{text}
}

func (el *TextElement) Render() string {
	return el.text
}

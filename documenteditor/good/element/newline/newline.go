package newline

type NewLineElement struct {
}

func Init() *NewLineElement {
	return &NewLineElement{}
}

func (el *NewLineElement) Render() string {
	return "\n"
}

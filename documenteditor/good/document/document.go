package document

import "strings"

type DocumentElement interface {
	Render() string
}

type Document struct {
	documentElements []DocumentElement
}

func Init() *Document {
	return &Document{}
}

func (doc *Document) AddElement(element DocumentElement) {
	doc.documentElements = append(doc.documentElements, element)
}

func (doc *Document) Render() string {
	var res strings.Builder
	for _, element := range doc.documentElements {
		res.WriteString(element.Render())
	}
	return res.String()
}

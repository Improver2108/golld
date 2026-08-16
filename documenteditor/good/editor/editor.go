package editor

import (
	"github.com/improver2108/golld/documenteditor/good/document"
	"github.com/improver2108/golld/documenteditor/good/element/image"
	"github.com/improver2108/golld/documenteditor/good/element/newline"
	"github.com/improver2108/golld/documenteditor/good/element/tabspace"
	"github.com/improver2108/golld/documenteditor/good/element/text"
)

type Persistence interface {
	Save(data string)
}

type DocumentEditor struct {
	document         *document.Document
	persistence      Persistence
	renderedDocument string
}

func Init(document *document.Document, persistence Persistence) *DocumentEditor {
	return &DocumentEditor{document: document, persistence: persistence}
}

func (e *DocumentEditor) AddText(data string) {
	e.document.AddElement(text.Init(data))
}

func (e *DocumentEditor) AddImage(path string) {
	e.document.AddElement(image.Init(path))
}

func (e *DocumentEditor) AddNewLine() {
	e.document.AddElement(newline.Init())
}

func (e *DocumentEditor) AddTabSpace() {
	e.document.AddElement(tabspace.Init())
}

func (e *DocumentEditor) RenderDocument() string {
	e.renderedDocument = e.document.Render()
	return e.renderedDocument
}

func (e *DocumentEditor) SaveDocument() {
	e.persistence.Save(e.renderedDocument)
}

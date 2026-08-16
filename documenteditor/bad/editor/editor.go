package editor

import (
	"fmt"
	"os"
	"strings"
)

type DocumentEditor struct {
	documentElements []string
	renderedDocument string
}

func Init() *DocumentEditor {
	return &DocumentEditor{}
}

func (e *DocumentEditor) AddText(text string) {
	e.documentElements = append(e.documentElements, text)
}

func (e *DocumentEditor) AddImage(path string) {
	e.documentElements = append(e.documentElements, path)
}

func (e *DocumentEditor) RenderDocument() string {
	for _, element := range e.documentElements {
		if strings.HasSuffix(element, ".jpg") || strings.HasSuffix(element, ".png") {
			e.renderedDocument += fmt.Sprintf("[Image : %s]\n", element)
		} else {
			e.renderedDocument += fmt.Sprintf("%s\n", element)
		}
	}
	return e.renderedDocument
}

func (e *DocumentEditor) SaveToFile() {
	if err := os.WriteFile("./documenteditor/bad/document.txt", []byte(e.renderedDocument), 0644); err != nil {
		fmt.Println("Error: Unable to open file for writing.")
		return
	}
	fmt.Println("Document saved to document.txt")
}

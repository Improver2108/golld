package main

import (
	"fmt"

	"github.com/improver2108/golld/documenteditor/good/document"
	"github.com/improver2108/golld/documenteditor/good/editor"
	"github.com/improver2108/golld/documenteditor/good/persistance/file"
)

func main() {
	document := document.Init()
	persistence := file.Init()

	editor := editor.Init(document, persistence)
	editor.AddText("Hello, world!")
	editor.AddNewLine()
	editor.AddText("This is a real-world document editor example.")
	editor.AddNewLine()
	editor.AddTabSpace()
	editor.AddText("Indented text after a tab space.")
	editor.AddNewLine()
	editor.AddImage("picture.jpg")

	fmt.Println(editor.RenderDocument())

	editor.SaveDocument()
}

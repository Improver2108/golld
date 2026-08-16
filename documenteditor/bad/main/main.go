package main

import (
	"fmt"

	"github.com/improver2108/golld/documenteditor/bad/editor"
)

func main() {
	editor := editor.Init()
	editor.AddText("Hello World!")
	editor.AddImage("picture.jpg")
	editor.AddText("")
	editor.AddText("This is a document editor.")

	fmt.Println(editor.RenderDocument())

	editor.SaveToFile()
}

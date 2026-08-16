package image

import "fmt"

type ImageElement struct {
	imagePath string
}

func Init(path string) *ImageElement {
	return &ImageElement{path}
}

func (el *ImageElement) Render() string {
	return fmt.Sprintf("[Image : %s]\n", el.imagePath)
}

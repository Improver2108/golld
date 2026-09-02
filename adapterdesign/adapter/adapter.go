package adapter

import (
	"fmt"
	"strings"
)

type XMLDataProvider interface {
	GetXMLData(data string) string
}

type xmlDataProviderAdapter struct {
	xmlDataProvider XMLDataProvider
}

func NewXMLDataProviderAdapter(provider XMLDataProvider) *xmlDataProviderAdapter {
	return &xmlDataProviderAdapter{
		xmlDataProvider: provider,
	}
}

func (a *xmlDataProviderAdapter) GetJSONData(data string) string {
	xml := a.xmlDataProvider.GetXMLData(data)

	startName := strings.Index(xml, "<name>") + 6
	endName := strings.Index(xml, "</name>")
	name := xml[startName:endName]

	startId := strings.Index(xml, "<id>") + 4
	endId := strings.Index(xml, "</id>")
	id := xml[startId:endId]

	return fmt.Sprintf("{\"name\":\"%s\", \"id\":\"%s\"}", name, id)
}

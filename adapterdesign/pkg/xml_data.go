package pkg

import (
	"fmt"
	"strings"
)

type xmlDataProvider struct{}

func NewXMLDataProvider() *xmlDataProvider {
	return &xmlDataProvider{}
}

func (p *xmlDataProvider) GetXMLData(data string) string {
	dataSplit := strings.Split(data, ":")
	name, id := dataSplit[0], dataSplit[1]
	return fmt.Sprintf("<user> <name> %s </name> <id> %s </id> </user>", name, id)
}

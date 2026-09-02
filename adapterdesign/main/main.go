package main

import (
	"github.com/improver2108/golld/adapterdesign/adapter"
	"github.com/improver2108/golld/adapterdesign/client"
	"github.com/improver2108/golld/adapterdesign/pkg"
)

func main() {
	xmlProvider := pkg.NewXMLDataProvider()

	adptr := adapter.NewXMLDataProviderAdapter(xmlProvider)

	rawData := "Alice:42"

	client := client.NewClient()

	client.GetReport(adptr, rawData)
}

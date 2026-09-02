package client

import "fmt"

type Report interface {
	GetJSONData(data string) string
}

type client struct{}

func NewClient() *client {
	return &client{}
}

func (c *client) GetReport(report Report, rawData string) {
	fmt.Println("Processed JSON:", report.GetJSONData(rawData))
}

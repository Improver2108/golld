package sql

import "fmt"

type SQLDatabase struct {
}

func Init() *SQLDatabase {
	return &SQLDatabase{}
}

func (db *SQLDatabase) SaveToSQl(data string) {
	fmt.Println("Executing SQL Query: INSERT INTO users VALUES('" + data + "');")
}

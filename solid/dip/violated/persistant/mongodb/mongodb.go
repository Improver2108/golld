package mongodb

import (
	"fmt"
)

type MongoDBDatabase struct{}

func Init() *MongoDBDatabase {
	return &MongoDBDatabase{}
}

func (db *MongoDBDatabase) SaveToMongo(data string) {
	fmt.Println(
		"Executing MongoDB Function: db.users.insert({name: '" + data + "'})",
	)
}

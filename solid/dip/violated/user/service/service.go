package service

import (
	"github.com/improver2108/golld/solid/dip/violated/persistant/mongodb"
	"github.com/improver2108/golld/solid/dip/violated/persistant/sql"
)

type UserService struct {
	sqlDb   sql.SQLDatabase
	mongodb mongodb.MongoDBDatabase
}

func Init() *UserService {
	return &UserService{}
}

func (u *UserService) StoreUserToSQL(user string) {
	u.sqlDb.SaveToSQl(user)
}

func (u *UserService) StoreUserToMongo(user string) {
	u.mongodb.SaveToMongo(user)
}

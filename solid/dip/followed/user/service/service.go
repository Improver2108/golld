package service

type Database interface {
	Save(data string)
}

type UserService struct {
	db Database
}

func Init(db Database) *UserService {
	return &UserService{db: db}
}

func (u *UserService) StoreUser(user string) {
	u.db.Save(user)
}

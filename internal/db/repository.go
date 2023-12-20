package db

type Repository interface {
	CreateUser(user User) (uuid string, err error)
	TopUpBalance(userUuid string, reqAmount uint32) error
}

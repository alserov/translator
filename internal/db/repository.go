package db

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user User) (err error)
	TopUpRequests(ctx context.Context, userUuid string, reqAmount uint32) error
	DebitRequests(ctx context.Context, userUuid string, reqAmount uint32) error
}

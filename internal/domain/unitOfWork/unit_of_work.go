package unitOfWork

import "context"

type UnitOfWork interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Complete() error
	Repository(name string) interface{}
}

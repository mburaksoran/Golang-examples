package user

import "context"

type Service interface {
	Get(ctx context.Context, id uint) (*Model, error)
	Create(ctx context.Context, model Model) (uint, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, model Model) error
}

type service struct {
	repo Repository
}

var _ Service = service{}

func NewService(repo Repository) Service {
	return service{repo: repo}
}

func (serv service) Get(ctx context.Context, id uint) (*Model, error) {
	return serv.repo.Get(ctx, id)
}

func (serv service) Create(ctx context.Context, model Model) (uint, error) {
	return serv.repo.Create(ctx, model)
}

func (serv service) Delete(ctx context.Context, id uint) error {
	return serv.repo.Delete(ctx, id)
}

func (serv service) Update(ctx context.Context, id uint, model Model) error {
	return serv.repo.Update(ctx, id, model)
}

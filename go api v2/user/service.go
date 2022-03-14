package user

type Service interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Delete(id uint) error
	Update(id uint, model Model) error
}

type service struct {
	repo Repository
}

var _ Service = service{}

func NewService(repo Repository) Service {
	return service{repo: repo}
}

func (serv service) Get(id uint) (*Model, error) {
	return serv.repo.Get(id)
}

func (serv service) Create(model Model) (uint, error) {
	return serv.repo.Create(model)
}

func (serv service) Delete(id uint) error {
	return serv.repo.Delete(id)
}

func (serv service) Update(id uint, model Model) error {
	return serv.repo.Update(id, model)
}

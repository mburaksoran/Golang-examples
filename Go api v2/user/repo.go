package user

import "gorm.io/gorm"

type Repository interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Delete(id uint) error
	Update(id uint, model Model) error
	Migration() error
}

type repository struct {
	db *gorm.DB
}

var _ Repository = repository{}

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}
}

func (repo repository) Get(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.First(model).Error
	// select * from users where id= id
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo repository) Create(model Model) (uint, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (repo repository) Delete(id uint) error {
	model := &Model{ID: id}
	err := repo.db.Delete(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo repository) Update(id uint, newModel Model) error {
	model := &Model{ID: id}
	err := repo.db.First(&model).Error
	if err != nil {
		return err
	}
	model.Email = newModel.Email
	model.Name = newModel.Name
	repo.db.Save(&model)
	return nil
}

func (repo repository) Migration() error {
	return repo.db.AutoMigrate(&Model{})
}

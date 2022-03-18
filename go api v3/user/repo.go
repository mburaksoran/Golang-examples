package user

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Get(ctx context.Context, id uint) (*Model, error)
	Create(ctx context.Context, model Model) (uint, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, model Model) error
	Migration() error
}

type repository struct {
	db *gorm.DB
}

var _ Repository = repository{}

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}
}

func (repo repository) Get(ctx context.Context, id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.WithContext(ctx).First(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo repository) Create(ctx context.Context, model Model) (uint, error) {
	err := repo.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (repo repository) Delete(ctx context.Context, id uint) error {
	model := &Model{ID: id}
	err := repo.db.WithContext(ctx).Delete(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo repository) Update(ctx context.Context, id uint, newModel Model) error {
	model := &Model{ID: id}
	err := repo.db.WithContext(ctx).First(&model).Error
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

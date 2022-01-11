package repositorytask

import (
	"github.com/arfan21/golang-kanbanboard/entity"
	"gorm.io/gorm"
)

type RepositoryTask interface {
	Create(task entity.Task) (entity.Task, error)
	IsCategoryExist(categoryID uint) error
	Gets() ([]entity.Task, error)
	Update(task entity.Task) (entity.Task, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(task entity.Task) (entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) IsCategoryExist(categoryID uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Gets() ([]entity.Task, error) {
	//TODO implement m e
	panic("implement me")
}

func (r *repository) Update(task entity.Task) (entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func New(db *gorm.DB) RepositoryTask {
	return &repository{db: db}
}

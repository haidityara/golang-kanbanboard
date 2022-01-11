package repositorytask

import (
	"errors"
	"github.com/arfan21/golang-kanbanboard/constant"
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
	err := r.db.Create(&task).Error
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *repository) IsCategoryExist(categoryID uint) error {
	category := new(entity.Category)
	err := r.db.Where("id = ?", categoryID).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrorCategoryDoesNotExists
		}
		return err
	}
	return nil
}

func (r *repository) Gets() ([]entity.Task, error) {
	var tasks []entity.Task
	err := r.db.Preload("User").Preload("Category").Find(&tasks).Error
	if err != nil {
		return []entity.Task{}, err
	}
	return tasks, nil
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

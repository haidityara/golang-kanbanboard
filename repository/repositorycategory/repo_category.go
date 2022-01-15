package repositorycategory

import (
	"github.com/arfan21/golang-kanbanboard/entity"
	"gorm.io/gorm"
	"log"
)

type RepositoryCategory interface {
	Create(category entity.Category) (entity.Category, error)
	Gets() ([]entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	Delete(ID uint) error
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (r *Repository) Gets() ([]entity.Category, error) {
	var categories []entity.Category
	//err := r.db.Preload("task").First(&categories).Error
	err := r.db.Preload("Task").Find(&categories).Error
	if err != nil {
		return []entity.Category{}, err
	}
	log.Println(categories)
	return categories, nil
}

func (r *Repository) Update(category entity.Category) (entity.Category, error) {
	err := r.db.Where("id = ?", category.ID).Updates(&category).Error
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (r *Repository) Delete(ID uint) error {
	category := entity.Category{}
	category.ID = ID
	err := r.db.First(&category).Where("id = ?", category.ID).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func New(db *gorm.DB) RepositoryCategory {
	return &Repository{db: db}
}

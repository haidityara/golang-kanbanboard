package servicecategory

import (
	"github.com/arfan21/golang-kanbanboard/entity"
	"github.com/arfan21/golang-kanbanboard/model/modelcategory"
	"github.com/arfan21/golang-kanbanboard/repository/repositorycategory"
	"github.com/arfan21/golang-kanbanboard/validation"
	"github.com/jinzhu/copier"
)

type ServiceCategory interface {
	Create(request modelcategory.Request) (modelcategory.Response, error)
	Gets() ([]modelcategory.ResponseGet, error)
}

type Service struct {
	repo repositorycategory.RepositoryCategory
}

func (s *Service) Gets() ([]modelcategory.ResponseGet, error) {
	var resp []modelcategory.ResponseGet
	gets, err := s.repo.Gets()
	if err != nil {
		return []modelcategory.ResponseGet{}, err
	}
	copier.Copy(&resp, &gets)
	return resp, nil
}

func (s *Service) Create(request modelcategory.Request) (modelcategory.Response, error) {
	err := validation.ValidateCategoryStore(request)
	if err != nil {
		return modelcategory.Response{}, err
	}
	entityCategory := new(entity.Category)
	copier.Copy(entityCategory, &request)
	create, err := s.repo.Create(*entityCategory)
	if err != nil {
		return modelcategory.Response{}, err
	}
	resp := new(modelcategory.Response)
	copier.Copy(resp, &create)
	return *resp, nil
}

func New(repo repositorycategory.RepositoryCategory) ServiceCategory {
	return &Service{repo: repo}
}

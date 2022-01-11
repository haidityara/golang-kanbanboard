package servicetask

import (
	"github.com/arfan21/golang-kanbanboard/entity"
	"github.com/arfan21/golang-kanbanboard/model/modeltask"
	"github.com/arfan21/golang-kanbanboard/repository/repositorytask"
	"github.com/arfan21/golang-kanbanboard/validation"
	"github.com/jinzhu/copier"
)

type ServiceTask interface {
	Create(request modeltask.Request) (modeltask.ResponseStore, error)
	Get() (modeltask.ResponseGet, error)
}

type Service struct {
	repo repositorytask.RepositoryTask
}

func (s *Service) Create(request modeltask.Request) (modeltask.ResponseStore, error) {

	// validation
	err := validation.ValidateTaskCreate(request, s.repo)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}

	entityTask := new(entity.Task)
	copier.Copy(&entityTask, &request)

	entityTask.Status = false

	create, err := s.repo.Create(*entityTask)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}

	resp := new(modeltask.ResponseStore)
	copier.Copy(&resp, &create)
	return *resp, nil
}

func (s *Service) Get() (modeltask.ResponseGet, error) {
	//TODO implement me
	panic("implement me")
}

func New(repo repositorytask.RepositoryTask) ServiceTask {
	return &Service{repo: repo}
}

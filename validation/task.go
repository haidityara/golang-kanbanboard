package validation

import (
	"errors"
	"github.com/arfan21/golang-kanbanboard/model/modeltask"
	"github.com/arfan21/golang-kanbanboard/repository/repositorytask"
	validation "github.com/go-ozzo/ozzo-validation"
)

func isCategoryExists(repo repositorytask.RepositoryTask) validation.RuleFunc {
	return func(value interface{}) error {
		cid, ok := value.(uint)
		if !ok {
			return errors.New("invalid value")
		}

		return repo.IsCategoryExist(cid)
	}
}

func ValidateTaskCreate(data modeltask.Request, repo repositorytask.RepositoryTask) error {
	return validation.Errors{
		"title":       validation.Validate(data.Title, validation.Required),
		"description": validation.Validate(data.Description, validation.Required),
		"category_id": validation.Validate(data.CategoryID, validation.Required, validation.By(isCategoryExists(repo))),
	}.Filter()
}

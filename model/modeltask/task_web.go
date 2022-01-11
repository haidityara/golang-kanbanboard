package modeltask

import (
	"github.com/arfan21/golang-kanbanboard/model/modeluser"
	"time"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status,omitempty"`
	CategoryID  uint   `json:"category_id"`
	UserID      uint   `json:"user_id,omitempty"`
}

type ResponseStore struct {
	Title       string     `json:"title"`
	Status      bool       `json:"status"`
	Description string     `json:"description"`
	UserID      uint       `json:"user_id"`
	CategoryID  uint       `json:"category_id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type ResponseGet struct {
	ResponseStore
	User modeluser.Response `json:"user"`
}

type ExampleResponseDelete struct {
	Message string `json:"message" example:"Task has been deleted"`
}

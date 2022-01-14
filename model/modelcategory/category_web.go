package modelcategory

import (
	"time"
)

type Request struct {
	Type string `json:"type"`
}

type Response struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ResponseGet struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
	Task []struct {
		ID          uint64     `json:"id"`
		Title       string     `json:"title"`
		Status      bool       `json:"status"`
		Description string     `json:"description"`
		UserID      uint       `json:"user_id"`
		CategoryID  uint       `json:"category_id"`
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	} `json:"tasks"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

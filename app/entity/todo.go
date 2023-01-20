package entity

import "time"

type Todo struct {
	Id          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	File        string    `json:"file,omitempty"`
	IsParent    bool      `json:"is_parent,omitempty"`
	ParentId    int       `json:"parent_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type TodoWithChild struct {
	Id          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	File        string    `json:"file,omitempty"`
	IsParent    bool      `json:"is_parent,omitempty"`
	ParentId    int       `json:"parent_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
	ChildTodo   []Todo    `json:"child_todo"`
}

type TodoRequest struct {
	Title       string `json:"title" validate:"required,gte=0,lte=100"`
	Description string `json:"description" validate:"required,gte=0,lte=1000"`
	File        string `json:"-"`
}

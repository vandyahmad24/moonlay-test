package repository

import (
	"fmt"
	"github.com/vandyahmad24/moonlay-test/app/entity"
	"gorm.io/gorm"
)

type Repository interface {
	GetParent(query entity.Query) ([]entity.Todo, error)
	GetChild(parentId int) ([]entity.Todo, error)
	GetParentById(id int) (entity.Todo, error)
	GetAllById(id int) (entity.Todo, error)
	Create(req entity.Todo) (entity.Todo, error)
	CountAllParent() int64
	DeleteById(id int) error
	Update(req entity.Todo) (entity.Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetParent(query entity.Query) ([]entity.Todo, error) {
	var allTodo []entity.Todo
	q := r.db.Limit(query.Limit).Offset(query.Skip).Where("is_parent = ? ", 1)
	if query.Search != "" {
		q = q.Where("title LIKE ?", fmt.Sprint("%", query.Search, "%")).Where("description LIKE ?", fmt.Sprint("%", query.Search, "%"))
	}

	err := q.Find(&allTodo).Debug().Error
	if err != nil {
		return allTodo, err
	}
	return allTodo, nil
}

func (r *repository) CountAllParent() (hasil int64) {
	var total int64
	r.db.Table("todos").Where("is_parent = ? ", 1).Select("count(id)").Count(&total)

	return total
}

func (r *repository) GetChild(parentId int) ([]entity.Todo, error) {
	var allTodo []entity.Todo
	q := r.db.Where("parent_id = ? ", parentId)

	err := q.Find(&allTodo).Debug().Error
	if err != nil {
		return allTodo, err
	}
	return allTodo, nil
}

func (r *repository) GetParentById(id int) (entity.Todo, error) {
	var allTodo entity.Todo
	q := r.db.Where("id = ? ", id)

	err := q.First(&allTodo).Debug().Error
	if err != nil {
		return allTodo, err
	}
	return allTodo, nil
}

func (r *repository) Create(req entity.Todo) (entity.Todo, error) {

	err := r.db.Create(&req).Debug().Error
	if err != nil {
		return req, err
	}
	return req, nil
}

func (r *repository) Update(req entity.Todo) (entity.Todo, error) {

	err := r.db.Debug().Table("todos").Where("id = ?", req.Id).Updates(map[string]interface{}{"title": req.Title, "description": req.Description, "file": req.File}).Error
	if err != nil {
		return req, err
	}
	return req, nil
}

func (r *repository) GetAllById(id int) (entity.Todo, error) {
	var allTodo entity.Todo
	q := r.db.Where("id = ? ", id)

	err := q.First(&allTodo).Debug().Error
	if err != nil {
		return allTodo, err
	}
	return allTodo, nil
}

func (r *repository) DeleteById(id int) error {
	var todo entity.Todo
	err := r.db.Where("id =?", id).Debug().Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

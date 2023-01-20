package usecase

import (
	"github.com/vandyahmad24/moonlay-test/app/entity"
	"github.com/vandyahmad24/moonlay-test/app/repository"
)

type TodoUsecase interface {
	GetParent(query entity.Query) ([]entity.Todo, error)
	GetParentWithChild(query entity.Query) ([]entity.TodoWithChild, error)
	DetailParent(id int) (entity.TodoWithChild, error)
	Create(req entity.TodoRequest) (entity.Todo, error)
	CreateSubTask(req entity.TodoRequest, parentInt int) (entity.Todo, error)
	Update(req entity.TodoRequest, id int) (entity.Todo, error)
	DeleteById(id int) error
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(repository repository.Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

func (s *usecase) GetParent(query entity.Query) ([]entity.Todo, error) {
	todo, err := s.repository.GetParent(query)
	if err != nil {
		return todo, err
	}
	return todo, err
}

func (s *usecase) GetParentWithChild(query entity.Query) ([]entity.TodoWithChild, error) {
	var result []entity.TodoWithChild
	todo, err := s.repository.GetParent(query)
	if err != nil {
		return result, err
	}

	for _, v := range todo {
		child, err := s.repository.GetChild(v.Id)
		if err != nil {
			continue
		}
		temp := entity.TodoWithChild{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			File:        v.File,
			IsParent:    v.IsParent,
			ParentId:    v.ParentId,
			ChildTodo:   child,
		}
		result = append(result, temp)
	}

	return result, err
}

func (s *usecase) DetailParent(id int) (entity.TodoWithChild, error) {
	var result entity.TodoWithChild
	todo, err := s.repository.GetParentById(id)
	if err != nil {
		return result, err
	}
	result.Id = todo.Id
	result.Title = todo.Title
	result.Description = todo.Description
	result.File = todo.File
	result.IsParent = todo.IsParent
	result.ParentId = todo.ParentId

	child, _ := s.repository.GetChild(todo.Id)

	result.ChildTodo = child

	return result, err
}

func (s *usecase) Create(req entity.TodoRequest) (entity.Todo, error) {
	model := entity.Todo{
		Title:       req.Title,
		Description: req.Description,
		File:        req.File,
		IsParent:    true,
		ParentId:    0,
	}
	result, err := s.repository.Create(model)
	if err != nil {
		return model, err
	}
	return result, nil
}

func (s *usecase) CreateSubTask(req entity.TodoRequest, parentInt int) (entity.Todo, error) {
	model := entity.Todo{
		Title:       req.Title,
		Description: req.Description,
		File:        req.File,
		IsParent:    false,
		ParentId:    parentInt,
	}
	result, err := s.repository.Create(model)
	if err != nil {
		return model, err
	}
	return result, nil
}

func (s *usecase) Update(req entity.TodoRequest, id int) (entity.Todo, error) {
	// get first
	var model entity.Todo
	data, err := s.repository.GetAllById(id)
	if err != nil {
		return model, err
	}

	data.File = req.File
	data.Description = req.Description
	data.Title = req.Title

	result, err := s.repository.Update(data)
	if err != nil {
		return model, err
	}
	return result, nil
}

func (s *usecase) DeleteById(id int) error {
	// get first

	data, err := s.repository.GetAllById(id)
	if err != nil {
		return err
	}

	err = s.repository.DeleteById(data.Id)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/vandyahmad24/moonlay-test/app/entity"
)

type MockCakeRepository struct {
	mock.Mock
}

func (m *MockCakeRepository) GetParent(query entity.Query) ([]entity.Todo, error) {
	args := m.Called(query)
	return args.Get(0).([]entity.Todo), args.Error(1)
}

func (m *MockCakeRepository) GetChild(parentId int) ([]entity.Todo, error) {
	args := m.Called(parentId)
	return args.Get(0).([]entity.Todo), args.Error(1)
}

func (m *MockCakeRepository) GetParentById(id int) (entity.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Todo), args.Error(1)
}

func (m *MockCakeRepository) GetAllById(id int) (entity.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Todo), args.Error(1)
}

func (m *MockCakeRepository) Create(req entity.Todo) (entity.Todo, error) {
	args := m.Called(req)
	return args.Get(0).(entity.Todo), args.Error(1)
}

func (m *MockCakeRepository) CountAllParent() int64 {

	return 1
}

func (m *MockCakeRepository) DeleteById(id int) error {
	args := m.Called(id)
	if id == 0 {
		return args.Error(0)
	}
	return nil
}

func (m *MockCakeRepository) Update(req entity.Todo) (entity.Todo, error) {
	args := m.Called(req)
	return args.Get(0).(entity.Todo), args.Error(1)
}

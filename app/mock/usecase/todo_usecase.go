package usecase

import (
	"github.com/stretchr/testify/mock"
	"github.com/vandyahmad24/moonlay-test/app/entity"
)

type MockCakeUsecase struct {
	mock.Mock
}

func (m *MockCakeUsecase) GetParent(query entity.Query) ([]entity.Todo, error) {
	args := m.Called(query)
	return args.Get(0).([]entity.Todo), args.Error(1)
}

func (m *MockCakeUsecase) GetParentWithChild(query entity.Query) ([]entity.TodoWithChild, error) {
	args := m.Called(query)
	return args.Get(0).([]entity.TodoWithChild), args.Error(1)
}

func (m *MockCakeUsecase) DetailParent(id int) (entity.TodoWithChild, error) {
	args := m.Called(id)
	return args.Get(0).(entity.TodoWithChild), args.Error(1)
}

func (m *MockCakeUsecase) Create(req entity.TodoRequest) (entity.Todo, error) {
	args := m.Called(req)
	return args.Get(0).(entity.Todo), args.Error(1)
}

func (m *MockCakeUsecase) CreateSubTask(req entity.TodoRequest, parentInt int) (entity.Todo, error) {
	args := m.Called(req)
	return args.Get(0).(entity.Todo), args.Error(1)
}
func (m *MockCakeUsecase) Update(req entity.TodoRequest, id int) (entity.Todo, error) {
	args := m.Called(req, id)
	return args.Get(0).(entity.Todo), args.Error(1)
}

func (m *MockCakeUsecase) DeleteById(id int) error {
	args := m.Called(id)
	return args.Error(1)
}

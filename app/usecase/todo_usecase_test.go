package usecase

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vandyahmad24/moonlay-test/app/entity"
	"github.com/vandyahmad24/moonlay-test/app/mock/repository"
	"testing"
)

func TestTodoUsecase(t *testing.T) {
	mockRepo := repository.MockCakeRepository{}
	Convey("Unit Test TodoInteractor", t, func() {
		Convey("Scenario GetParent", func() {
			Convey("Positive Scenario GetParent", func() {

				var query entity.Query
				var response []entity.Todo
				mockRepo.On("GetParent", query).Return(response, nil).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.GetParent(query)
				So(err, ShouldBeNil)
				So(len(res), ShouldEqual, 0)
			})
			Convey("Negative Scenario GetParent", func() {

				var query entity.Query
				var response []entity.Todo
				mockRepo.On("GetParent", query).Return(response, fmt.Errorf("Error")).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.GetParent(query)
				So(err, ShouldNotBeNil)
				So(len(res), ShouldEqual, 0)
			})
		})
		Convey("Scenario GetParentWithChild", func() {
			Convey("Positive Scenario GetParentWithChild", func() {
				var query entity.Query
				var response []entity.Todo
				mockRepo.On("GetParent", query).Return(response, nil).Once()
				mockRepo.On("GetParentWithChild", query).Return(response, nil).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.GetParentWithChild(query)
				So(err, ShouldBeNil)
				So(len(res), ShouldEqual, 0)
			})
			Convey("Negative Scenario GetParentWithChild", func() {

				var query entity.Query
				var response []entity.Todo
				mockRepo.On("GetParent", query).Return(response, fmt.Errorf("Error")).Once()
				mockRepo.On("GetParentWithChild", query).Return(response, fmt.Errorf("Error")).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.GetParentWithChild(query)
				So(err, ShouldNotBeNil)
				So(len(res), ShouldEqual, 0)
			})
		})
		Convey("Scenario DetailParent", func() {
			Convey("Negative Scenario DetailParent", func() {
				var response entity.Todo
				mockRepo.On("GetParentById", 0).Return(response, fmt.Errorf("Error")).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.DetailParent(0)
				So(err, ShouldNotBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Positive Scenario DetailParent", func() {
				var response entity.Todo
				var resp []entity.Todo
				mockRepo.On("GetParentById", 0).Return(response, nil).Once()
				mockRepo.On("GetChild", 0).Return(resp, nil).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.DetailParent(0)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})

		Convey("Scenario Create", func() {
			Convey("Negative Scenario Create", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    true,
					ParentId:    0,
				}
				mockRepo.On("Create", model).Return(model, fmt.Errorf("Error")).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.Create(req)
				So(err, ShouldNotBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Postive Scenario Create", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    true,
					ParentId:    0,
				}
				mockRepo.On("Create", model).Return(model, nil).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.Create(req)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Scenario Create Subtask", func() {
			Convey("Negative Scenario Subtask", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("Create", model).Return(model, fmt.Errorf("Error")).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.CreateSubTask(req, 0)
				So(err, ShouldNotBeNil)
				So(res, ShouldNotBeNil)
			})

			Convey("Positive Scenario Subtask", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("Create", model).Return(model, nil).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.CreateSubTask(req, 0)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Scenario Update", func() {
			Convey("Negative Scenario Update Failed Get By Id", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("GetAllById", 0).Return(model, fmt.Errorf("error")).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.Update(req, 0)
				So(err, ShouldNotBeNil)
				So(res, ShouldNotBeNil)
			})

			Convey("Negative Scenario Update ", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("GetAllById", 0).Return(model, nil).Once()
				mockRepo.On("Update", model).Return(model, fmt.Errorf("error")).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.Update(req, 0)
				So(err, ShouldNotBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Positive Scenario Update ", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("GetAllById", 0).Return(model, nil).Once()
				mockRepo.On("Update", model).Return(model, nil).Once()
				uc := NewUsecase(&mockRepo)
				res, err := uc.Update(req, 0)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})

		Convey("Scenario Delete", func() {
			Convey("Negative Scenario Delete Failed Get By Id", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("GetAllById", 0).Return(model, fmt.Errorf("error")).Once()
				uc := NewUsecase(&mockRepo)
				err := uc.DeleteById(0)
				So(err, ShouldNotBeNil)
			})

			Convey("Negative Scenario Delete ", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("GetAllById", 0).Return(model, nil).Once()
				mockRepo.On("DeleteById", 0).Return(fmt.Errorf("error")).Once()
				uc := NewUsecase(&mockRepo)
				err := uc.DeleteById(0)
				So(err, ShouldNotBeNil)

			})
			Convey("Positive Scenario Update ", func() {
				var req entity.TodoRequest
				model := entity.Todo{
					Title:       req.Title,
					Description: req.Description,
					File:        req.File,
					IsParent:    false,
					ParentId:    0,
				}
				mockRepo.On("GetAllById", 0).Return(model, nil).Once()
				mockRepo.On("DeleteById", 0).Return(nil).Once()
				uc := NewUsecase(&mockRepo)
				err := uc.DeleteById(0)
				So(err, ShouldBeNil)
			})
		})

	})
}

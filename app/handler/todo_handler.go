package handler

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/vandyahmad24/moonlay-test/app/entity"
	"github.com/vandyahmad24/moonlay-test/app/helper"
	"github.com/vandyahmad24/moonlay-test/app/usecase"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHanler(todoUsecase usecase.TodoUsecase) *todoHandler {
	return &todoHandler{
		todoUsecase: todoUsecase,
	}
}

func (t *todoHandler) GetAllParent(c echo.Context) error {
	// Get team and member from the query string
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")
	search := c.QueryParam("search")
	if limit == "" {
		limit = strconv.Itoa(10)
	}
	if skip == "" {
		skip = strconv.Itoa(0)
	}
	lim, _ := strconv.Atoi(limit)
	skp, _ := strconv.Atoi(skip)

	query := entity.Query{
		Limit:  lim,
		Skip:   skp,
		Search: search,
	}

	data, err := t.todoUsecase.GetParent(query)
	if err != nil {
		response := helper.ApiResponse(false, "Todo Not found", nil)
		return c.JSON(http.StatusOK, response)
	}
	response := helper.ApiResponse(true, "Get parent todo", data)
	return c.JSON(http.StatusOK, response)

}

func (t *todoHandler) GetAllWithChild(c echo.Context) error {
	// Get team and member from the query string
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")
	search := c.QueryParam("search")
	if limit == "" {
		limit = strconv.Itoa(10)
	}
	if skip == "" {
		skip = strconv.Itoa(0)
	}
	lim, _ := strconv.Atoi(limit)
	skp, _ := strconv.Atoi(skip)

	query := entity.Query{
		Limit:  lim,
		Skip:   skp,
		Search: search,
	}

	data, err := t.todoUsecase.GetParentWithChild(query)
	if err != nil {
		response := helper.ApiResponse(false, "Todo Not found", nil)
		return c.JSON(http.StatusOK, response)
	}
	response := helper.ApiResponse(true, "Get parent todo", data)
	return c.JSON(http.StatusOK, response)

}

func (t *todoHandler) GetDetailTodo(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := t.todoUsecase.DetailParent(idInt)
	if err != nil {
		response := helper.ApiResponse(false, "Todo Not found", nil)
		return c.JSON(http.StatusOK, response)
	}
	response := helper.ApiResponse(true, "Get parent todo", data)
	return c.JSON(http.StatusOK, response)

}

func (t *todoHandler) CreateTodo(c echo.Context) error {

	title := c.FormValue("title")
	description := c.FormValue("description")

	req := entity.TodoRequest{
		Title:       title,
		Description: description,
	}
	var newFileName string

	if err := c.Validate(&req); err != nil {
		resultHelper := helper.ValidationError(err)
		response := helper.ApiResponse(false, "Bad request", resultHelper)
		return c.JSON(http.StatusBadRequest, response)
	}

	// handle upload file
	file, err := c.FormFile("file")
	if err == nil {
		// Save the file to the assets folder
		src, err := file.Open()
		if err != nil {
			response := helper.ApiResponse(false, "Bad request", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		defer src.Close()

		// Generate a random UUID for the file name
		u := uuid.NewV4()

		// Get the file extension
		ext := filepath.Ext(file.Filename)

		// Create new file name
		newFileName = u.String() + ext

		dst, err := os.Create("./assets/" + newFileName)
		if err != nil {
			response := helper.ApiResponse(false, "Bad request Create asset", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			response := helper.ApiResponse(false, "Bad request Copy", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	// create
	req.File = newFileName
	result, err := t.todoUsecase.Create(req)
	if err != nil {
		response := helper.ApiResponse(false, "Internal Error", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.ApiResponse(true, "Get parent todo", result)
	return c.JSON(http.StatusOK, response)

}

func (t *todoHandler) CreateSublist(c echo.Context) error {

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := t.todoUsecase.DetailParent(idInt)
	if err != nil {
		response := helper.ApiResponse(false, "Todo Not found", nil)
		return c.JSON(http.StatusOK, response)
	}

	title := c.FormValue("title")
	description := c.FormValue("description")

	req := entity.TodoRequest{
		Title:       title,
		Description: description,
	}
	var newFileName string

	if err := c.Validate(&req); err != nil {
		resultHelper := helper.ValidationError(err)
		response := helper.ApiResponse(false, "Bad request", resultHelper)
		return c.JSON(http.StatusBadRequest, response)
	}

	// handle upload file
	file, err := c.FormFile("file")
	if err == nil {
		// Save the file to the assets folder
		src, err := file.Open()
		if err != nil {
			response := helper.ApiResponse(false, "Bad request", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		defer src.Close()

		// Generate a random UUID for the file name
		u := uuid.NewV4()

		// Get the file extension
		ext := filepath.Ext(file.Filename)

		// Create new file name
		newFileName = u.String() + ext

		dst, err := os.Create("./assets/" + newFileName)
		if err != nil {
			response := helper.ApiResponse(false, "Bad request Create asset", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			response := helper.ApiResponse(false, "Bad request Copy", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	// create
	req.File = newFileName
	result, err := t.todoUsecase.CreateSubTask(req, data.Id)
	if err != nil {
		response := helper.ApiResponse(false, "Internal Error", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.ApiResponse(true, "Get parent todo", result)
	return c.JSON(http.StatusOK, response)

}

func (t *todoHandler) Update(c echo.Context) error {

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	title := c.FormValue("title")
	description := c.FormValue("description")

	req := entity.TodoRequest{
		Title:       title,
		Description: description,
	}
	var newFileName string

	if err := c.Validate(&req); err != nil {
		resultHelper := helper.ValidationError(err)
		response := helper.ApiResponse(false, "Bad request", resultHelper)
		return c.JSON(http.StatusBadRequest, response)
	}

	// handle upload file
	file, err := c.FormFile("file")
	if err == nil {
		// Save the file to the assets folder
		src, err := file.Open()
		if err != nil {
			response := helper.ApiResponse(false, "Bad request", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		defer src.Close()

		// Generate a random UUID for the file name
		u := uuid.NewV4()

		// Get the file extension
		ext := filepath.Ext(file.Filename)

		// Create new file name
		newFileName = u.String() + ext

		dst, err := os.Create("./assets/" + newFileName)
		if err != nil {
			response := helper.ApiResponse(false, "Bad request Create asset", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			response := helper.ApiResponse(false, "Bad request Copy", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	// create
	req.File = newFileName

	resultUpdate, err := t.todoUsecase.Update(req, idInt)
	if err != nil {
		response := helper.ApiResponse(false, "Internal server error", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.ApiResponse(true, "Success Update", resultUpdate)
	return c.JSON(http.StatusOK, response)

}

func (t *todoHandler) Delete(c echo.Context) error {

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := t.todoUsecase.DeleteById(idInt)
	if err != nil {
		response := helper.ApiResponse(false, "Internal server error", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.ApiWithOutData(true, "Success Delete")
	return c.JSON(http.StatusOK, response)

}

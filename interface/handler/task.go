package handler

import (
	"github.com/labstack/echo/v4"
	"go-ddd-sample/usecase"
	"net/http"
	"strconv"
)

type TaskHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

type requestTask struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTask struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (th *taskHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTask, err := th.taskUsecase.Create(req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      createdTask.ID,
			Title:   createdTask.Title,
			Content: createdTask.Content,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (th *taskHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTask, err := th.taskUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      foundTask.ID,
			Title:   foundTask.Title,
			Content: foundTask.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTask, err := th.taskUsecase.Update(id, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      updatedTask.ID,
			Title:   updatedTask.Title,
			Content: updatedTask.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.taskUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}

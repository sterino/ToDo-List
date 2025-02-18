package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/internal/domain/task"
	interfaces "todo-list/internal/service/interface"
	"todo-list/pkg/response"
)

type TaskHandler struct {
	taskService interfaces.TaskService
}

func NewTaskHandler(service interfaces.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}

// CreateTask godoc
// @Summary Create a new payment
// @Description Create a new payment with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param payment body task.Request true "Task Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /tasks [post]
func (th *TaskHandler) CreateTask(c *gin.Context) {
	req := task.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := th.taskService.CreateTask(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, task.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields must be unique", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		if errors.Is(err, task.ErrorInvalidDate) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to create payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "the payment was successfully created", res, nil)
	c.JSON(http.StatusCreated, successRes)
}

// ListTasks godoc
// @Summary List all tasks
// @Description Get a list of tasks with optional status filter
// @Tags tasks
// @Accept json
// @Produce json
// @Param status query string false "Task Status"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /tasks [get]
func (th *TaskHandler) ListTasks(c *gin.Context) {
	status := c.Query("status")
	res, err := th.taskService.ListTasks(c.Request.Context(), status)
	if err != nil {
		if errors.Is(err, task.ErrorInvalidStatus) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		if errors.Is(err, task.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no tasks found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to list tasks", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the tasks list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetTask godoc
// @Summary Get a payment by ID
// @Description Get details of a payment by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /tasks/{id} [get]
func (th *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	res, err := th.taskService.GetTask(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, task.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusOK, "no tasks found", "", nil)
			c.JSON(http.StatusOK, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to get payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the payment details", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// UpdateTask godoc
// @Summary Update a payment by ID
// @Description Update details of a payment by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param payment body task.Request true "Task Request"
// @Success 204 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /tasks/{id} [put]
func (th *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	req := task.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := th.taskService.UpdateTask(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, task.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "payment not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return

		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to update payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusNoContent, "the payment was successfully updated", nil, nil)
	c.JSON(http.StatusNoContent, successRes)
}

// DeleteTask godoc
// @Summary Delete a payment by ID
// @Description Delete a payment by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 204 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /tasks/{id} [delete]
func (th *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := th.taskService.DeleteTask(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, task.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "payment not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete payment", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusNoContent, "the payment was successfully deleted", nil, nil)
	c.JSON(http.StatusNoContent, successRes)
}

// DoneTask godoc
// @Summary Mark a payment as done
// @Description Mark a payment as done by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 204 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /tasks/{id}/done [put]
func (th *TaskHandler) DoneTask(c *gin.Context) {
	id := c.Param("id")
	err := th.taskService.DoneTask(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, task.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "payment not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to mark payment as done", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusNoContent, "the payment was successfully marked as done", nil, nil)
	c.JSON(http.StatusNoContent, successRes)
}

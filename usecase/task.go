package usecase

import (
	"fmt"
	"go-ddd-sample/domain/model"
	"go-ddd-sample/domain/repository"
)

type TaskUsecase interface {
	Create(title, content string) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(id int, title, content string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (tu *taskUsecase) Create(title, content string) (*model.Task, error) {
	task, err := model.NewTask(title, content)
	if err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}

	createdTask, err := tu.taskRepo.Create(task)
	if err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}
	return createdTask, nil
}

func (tu *taskUsecase) FindByID(id int) (*model.Task, error) {
	foundTask, err := tu.taskRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("find error: %w", err)
	}

	return foundTask, nil
}

func (tu *taskUsecase) Update(id int, title, content string) (*model.Task, error) {
	targetTask, err := tu.taskRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("find error: %w", err)
	}
	err = targetTask.Set(title, content)
	if err != nil {
		return nil, fmt.Errorf("task set error: %w", err)
	}
	updatedTask, err := tu.taskRepo.Update(targetTask)
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}
	return updatedTask, nil
}

func (tu *taskUsecase) Delete(id int) error {
	task, err := tu.taskRepo.FindById(id)
	if err != nil {
		return fmt.Errorf("find error: %w", err)
	}

	err = tu.taskRepo.Delete(task)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}
	return nil
}

package infra

import (
	"fmt"
	"go-ddd-sample/domain/model"
	"go-ddd-sample/domain/repository"

	"gorm.io/gorm"
)

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Create(task).Error; err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}

	return task, nil
}

func (tr *TaskRepository) FindById(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	if err := tr.Conn.First(task).Error; err != nil {
		return nil, fmt.Errorf("find error: %w", err)
	}

	return task, nil
}

func (tr *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Model(task).Updates(task).Error; err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}
	return task, nil
}

func (tr *TaskRepository) Delete(task *model.Task) error {
	if err := tr.Conn.Delete(task).Error; err != nil {
		return fmt.Errorf("delete error: %w", err)
	}
	return nil
}

package repository

import "go-ddd-sample/domain/model"

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	FindById(id int) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(task *model.Task) error
}

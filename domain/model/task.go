package model

import (
	"errors"
)

type Task struct {
	ID      int
	Title   string
	Content string
}

func NewTask(title, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("title is empty")
	}

	task := &Task{
		Title:   title,
		Content: content,
	}

	return task, nil
}

func (t *Task) Set(title, content string) error {
	if title == "" {
		return errors.New("title is empty")
	}

	t.Title = title
	t.Content = content

	return nil
}

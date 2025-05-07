package model

import (
	"time"

	"github.com/google/uuid"
)

type TaskID uuid.UUID
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID       TaskID     `json:"id" db:"id"`
	Title    string     `json:"string" db:"title"`
	Status   TaskStatus `json:"status" db:"status"`
	Deadline time.Time  `json:"deadline" db:"deadline"`
	Created  time.Time  `json:"created" db:"created"`
	Modified time.Time  `json:"modified" db:"modified"`
}

type Tasks []*Task

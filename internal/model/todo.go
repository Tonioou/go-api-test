package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created-at"`
	FinishedAt time.Time `json:"finished-at"`
	Finished   bool      `json:"finished"`
	Active     bool      `json:"active"`
}

func NewTodo(name string) *Todo {
	return &Todo{
		Id:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		Finished:  false,
		Active:    true,
	}
}

func (t *Todo) Finish() {
	t.FinishedAt = time.Now()
	t.Finished = true
}

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
	Enabled    bool      `json:"enabled"`
}

func NewTodo(name string) *Todo {
	return &Todo{
		Id:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		Finished:  false,
		Enabled:   true,
	}
}

func (t *Todo) Finish() {
	t.FinishedAt = time.Now()
	t.Finished = true
}

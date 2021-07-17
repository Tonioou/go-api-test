package model

import "github.com/google/uuid"

type Person struct {
	Id    uuid.UUID `json:"id" `
	Email string    `json:"email" binding:"required"`
	Name  string    `json:"name" binding:"required"`
	Age   int       `json:"age" binding:"required"`
}

func NewPerson(age int, email string, name string) *Person {
	return &Person{
		Id:    uuid.New(),
		Email: email,
		Age:   age,
		Name:  name,
	}
}

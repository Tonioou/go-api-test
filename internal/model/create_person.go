package model

type CreatePerson struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}

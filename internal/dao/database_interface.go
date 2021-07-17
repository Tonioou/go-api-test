package dao

import (
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type Database interface {
	Add(tableName string, obj interface{}) *errorx.Error
	Get(tableName string) ([]interface{}, *errorx.Error)
	GetById(tableName string, id uuid.UUID) (interface{}, *errorx.Error)
}

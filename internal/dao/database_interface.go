package dao

import (
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type Database interface {
	Add(tableName string, obj interface{}) *errorx.Error
	Get() (interface{}, *errorx.Error)
	GetById(id uuid.UUID) (interface{}, *errorx.Error)
}

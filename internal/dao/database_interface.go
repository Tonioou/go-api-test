package dao

import (
	"github.com/joomcode/errorx"
)

type Database interface {
	Add(tableName string, obj interface{}) *errorx.Error
	Get(tableName string) ([]interface{}, *errorx.Error)
	GetById(tableName string, id string) (interface{}, *errorx.Error)
	Delete(tableName string, id string) (interface{}, *errorx.Error)
}

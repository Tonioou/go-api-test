package dao

import "github.com/joomcode/errorx"

type Database interface {
	Add(tableName string, obj interface{}) *errorx.Error
	Get() (interface{}, *errorx.Error)
}

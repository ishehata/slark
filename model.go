package slark

import (
	"reflect"
)

type Model struct {
	Name          string
	TableName     string
	Fields        []Field
	OriginalModel interface{}
	Reflection    reflect.Type
}

type Field struct {
	Name   string
	DBType string
}

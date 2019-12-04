package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
)

// OperationLogModel is operation log model structure.
type OperationLogModel struct {
	Base

	Id        int64
	UserId    int64
	Path      string
	Method    string
	Ip        string
	Input     string
	CreatedAt string
	UpdatedAt string
}

// OperationLog return a default operation log model.
func OperationLog() OperationLogModel {
	return OperationLogModel{Base: Base{TableName: "goadmin_operation_log"}}
}

// Find return a default operation log model of given id.
func (t OperationLogModel) Find(id interface{}) OperationLogModel {
	item, _ := t.Table(t.TableName).Find(id)
	return t.MapToModel(item)
}

func (t OperationLogModel) SetConn(con db.Connection) OperationLogModel {
	t.Conn = con
	return t
}

// New create a new operation log model.
func (t OperationLogModel) New(userId int64, path, method, ip, input string) OperationLogModel {

	id, _ := t.Table(t.TableName).Insert(dialect.H{
		"user_id": userId,
		"path":    path,
		"method":  method,
		"ip":      ip,
		"input":   input,
	})

	t.Id = id
	t.UserId = userId
	t.Path = path
	t.Method = method
	t.Ip = ip
	t.Input = input

	return t
}

// MapToModel get the operation log model from given map.
func (t OperationLogModel) MapToModel(m map[string]interface{}) OperationLogModel {
	t.Id = m["id"].(int64)
	t.UserId = m["user_id"].(int64)
	t.Path, _ = m["path"].(string)
	t.Method, _ = m["method"].(string)
	t.Ip, _ = m["ip"].(string)
	t.Input, _ = m["input"].(string)
	t.CreatedAt, _ = m["created_at"].(string)
	t.UpdatedAt, _ = m["updated_at"].(string)
	return t
}

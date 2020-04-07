package models

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/collection"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
)

// SiteModel is role model structure.
type SiteModel struct {
	Base

	Id    int64
	Key   string
	Value string
	Desc  string
	State int64

	CreatedAt string
	UpdatedAt string
}

const (
	SiteItemOpenState = 1
	SiteItemOffState  = 0
)

// Site return a default role model.
func Site() SiteModel {
	return SiteModel{Base: Base{TableName: "goadmin_site"}}
}

func (t SiteModel) SetConn(con db.Connection) SiteModel {
	t.Conn = con
	return t
}

func (t SiteModel) WithTx(tx *sql.Tx) SiteModel {
	t.Tx = tx
	return t
}

func (t SiteModel) Init(cfg map[string]string) {
	items, err := t.Table(t.TableName).All()
	if db.CheckError(err, db.QUERY) {
		panic(err)
	}
	itemsCol := collection.Collection(items)
	for key, value := range cfg {
		row := itemsCol.Where("key", "=", key)
		if row.Length() == 0 {
			_, err := t.Table(t.TableName).Insert(dialect.H{
				"key":   key,
				"value": value,
				"state": SiteItemOpenState,
			})
			if db.CheckError(err, db.INSERT) {
				panic(err)
			}
		} else {
			if value != "" {
				_, err := t.Table(t.TableName).
					Where("key", "=", key).Update(dialect.H{
					"value": value,
				})
				if db.CheckError(err, db.UPDATE) {
					panic(err)
				}
			}
		}
	}
}

func (t SiteModel) AllToMap() map[string]string {

	var m = make(map[string]string, 0)

	items, err := t.Table(t.TableName).Where("state", "=", SiteItemOpenState).All()
	if db.CheckError(err, db.QUERY) {
		return m
	}

	for _, item := range items {
		m[item["key"].(string)] = item["value"].(string)
	}

	return m
}

func (t SiteModel) AllToMapInterface() map[string]interface{} {

	var m = make(map[string]interface{}, 0)

	items, err := t.Table(t.TableName).Where("state", "=", SiteItemOpenState).All()
	if db.CheckError(err, db.QUERY) {
		return m
	}

	for _, item := range items {
		m[item["key"].(string)] = item["value"]
	}

	m["id"] = "1"

	return m
}

func (t SiteModel) Update(v form.Values) error {
	for key, vv := range v {
		if len(vv) > 0 && vv[0] != "" {
			_, err := t.Table(t.TableName).Where("key", "=", key).Update(dialect.H{
				"value": vv[0],
			})
			if db.CheckError(err, db.UPDATE) {
				return err
			}
		}
	}
	return nil
}

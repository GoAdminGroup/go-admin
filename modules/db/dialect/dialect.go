package dialect

import (
	"github.com/chenhg5/go-admin/modules/config"
	"strings"
)

type Dialect interface {
	// GetName get dialect's name
	GetName() string

	// ShowColumns show columns of specified table
	ShowColumns(table string) string

	// ShowTables show tables of database
	ShowTables() string

	// Insert
	Insert(comp *SqlComponent) string

	// Delete
	Delete(comp *SqlComponent) string

	// Update
	Update(comp *SqlComponent) string

	// Select
	Select(comp *SqlComponent) string
}

func GetDialect() Dialect {
	return GetDialectByDriver(config.Get().Databases.GetDefault().Driver)
}

func GetDialectByDriver(driver string) Dialect {
	switch driver {
	case "mysql":
		return mysql{}
	case "mssql":
		return mssql{}
	case "postgresql":
		return postgresql{}
	case "sqlite":
		return sqlite{}
	default:
		return commonDialect{}
	}
}

type H map[string]interface{}

type SqlComponent struct {
	Fields     []string
	TableName  string
	Wheres     []Where
	Leftjoins  []Join
	Args       []interface{}
	Order      string
	Offset     string
	Limit      string
	WhereRaws  string
	UpdateRaws []RawUpdate
	Statement  string
	Values     H
}

type Where struct {
	Operation string
	Field     string
	Qmark     string
}

type Join struct {
	Table     string
	FieldA    string
	Operation string
	FieldB    string
}

type RawUpdate struct {
	Expression string
	Args       []interface{}
}

// *******************************
// internal help function
// *******************************

func (sql *SqlComponent) getLimit() string {
	if sql.Limit == "" {
		return ""
	}
	return " limit " + sql.Limit + " "
}

func (sql *SqlComponent) getOffset() string {
	if sql.Offset == "" {
		return ""
	}
	return " offset " + sql.Offset + " "
}

func (sql *SqlComponent) getOrderBy() string {
	if sql.Order == "" {
		return ""
	}
	return " order by " + sql.Order + " "
}

func (sql *SqlComponent) getJoins() string {
	if len(sql.Leftjoins) == 0 {
		return ""
	}
	joins := ""
	for _, join := range sql.Leftjoins {
		joins += " left join " + join.Table + " on " + join.FieldA + " " + join.Operation + " " + join.FieldB + " "
	}
	return joins
}

func (sql *SqlComponent) getFields() string {
	if len(sql.Fields) == 0 {
		return "*"
	}
	if sql.Fields[0] == "count(*)" {
		return "count(*)"
	}
	fields := ""
	if len(sql.Leftjoins) == 0 {
		for _, field := range sql.Fields {
			fields += "`" + field + "`,"
		}
	} else {
		for _, field := range sql.Fields {
			arr := strings.Split(field, ".")
			if len(arr) > 1 {
				fields += arr[0] + ".`" + arr[1] + "`,"
			} else {
				fields += "`" + field + "`,"
			}
		}
	}
	return fields[:len(fields)-1]
}

func (sql *SqlComponent) getWheres() string {
	if len(sql.Wheres) == 0 {
		if sql.WhereRaws != "" {
			return " where " + sql.WhereRaws
		}
		return ""
	}
	wheres := " where "
	var arr []string
	for _, where := range sql.Wheres {
		arr = strings.Split(where.Field, ".")
		if len(arr) > 1 {
			wheres += arr[0] + ".`" + arr[1] + "` " + where.Operation + " " + where.Qmark + " and "
		} else {
			wheres += "`" + where.Field + "` " + where.Operation + " " + where.Qmark + " and "
		}
	}

	if sql.WhereRaws != "" {
		return wheres + sql.WhereRaws
	} else {
		return wheres[:len(wheres)-5]
	}
}

func (sql *SqlComponent) prepareUpdate() {
	fields := ""
	args := make([]interface{}, 0)

	if len(sql.Values) != 0 {

		for key, value := range sql.Values {
			fields += "`" + key + "` = ?, "
			args = append(args, value)
		}

		if len(sql.UpdateRaws) == 0 {
			fields = fields[:len(fields)-2]
		} else {
			for i := 0; i < len(sql.UpdateRaws); i++ {
				if i == len(sql.UpdateRaws)-1 {
					fields += sql.UpdateRaws[i].Expression + " "
				} else {
					fields += sql.UpdateRaws[i].Expression + ","
				}
				args = append(args, sql.UpdateRaws[i].Args...)
			}
		}

		sql.Args = append(args, sql.Args...)
	} else {
		if len(sql.UpdateRaws) == 0 {
			panic("prepareUpdate: wrong parameter")
		} else {
			for i := 0; i < len(sql.UpdateRaws); i++ {
				if i == len(sql.UpdateRaws)-1 {
					fields += sql.UpdateRaws[i].Expression + " "
				} else {
					fields += sql.UpdateRaws[i].Expression + ","
				}
				args = append(args, sql.UpdateRaws[i].Args...)
			}
		}
		sql.Args = append(args, sql.Args...)
	}

	sql.Statement = "update " + sql.TableName + " set " + fields + sql.getWheres()
}

func (sql *SqlComponent) prepareInsert() {
	fields := " ("
	quesMark := "("

	for key, value := range sql.Values {
		fields += "`" + key + "`,"
		quesMark += "?,"
		sql.Args = append(sql.Args, value)
	}
	fields = fields[:len(fields)-1] + ")"
	quesMark = quesMark[:len(quesMark)-1] + ")"

	sql.Statement = "insert into " + sql.TableName + fields + " values " + quesMark
}

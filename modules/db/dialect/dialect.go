package dialect

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
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

	GetDelimiter() string
}

func GetDialect() Dialect {
	return GetDialectByDriver(config.Get().Databases.GetDefault().Driver)
}

func GetDialectByDriver(driver string) Dialect {
	switch driver {
	case "mysql":
		return mysql{
			commonDialect: commonDialect{delimiter: "`"},
		}
	case "mssql":
		return mssql{
			commonDialect: commonDialect{delimiter: "`"},
		}
	case "postgresql":
		return postgresql{
			commonDialect: commonDialect{delimiter: `"`},
		}
	case "sqlite":
		return sqlite{
			commonDialect: commonDialect{delimiter: "`"},
		}
	default:
		return commonDialect{delimiter: "`"}
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

func (sql *SqlComponent) getJoins(delimiter string) string {
	if len(sql.Leftjoins) == 0 {
		return ""
	}
	joins := ""
	for _, join := range sql.Leftjoins {
		joins += " left join " + delimiter + join.Table + delimiter + " on " + join.FieldA + " " + join.Operation + " " + join.FieldB + " "
	}
	return joins
}

func (sql *SqlComponent) getFields(delimiter string) string {
	if len(sql.Fields) == 0 {
		return "*"
	}
	if sql.Fields[0] == "count(*)" {
		return "count(*)"
	}
	fields := ""
	if len(sql.Leftjoins) == 0 {
		for _, field := range sql.Fields {
			fields += delimiter + field + delimiter + ","
		}
	} else {
		for _, field := range sql.Fields {
			arr := strings.Split(field, ".")
			if len(arr) > 1 {
				fields += arr[0] + "." + delimiter + arr[1] + delimiter + ","
			} else {
				fields += delimiter + field + delimiter + ","
			}
		}
	}
	return fields[:len(fields)-1]
}

func (sql *SqlComponent) getWheres(delimiter string) string {
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
			wheres += arr[0] + "." + delimiter + arr[1] + delimiter + " " + where.Operation + " " + where.Qmark + " and "
		} else {
			wheres += delimiter + where.Field + delimiter + " " + where.Operation + " " + where.Qmark + " and "
		}
	}

	if sql.WhereRaws != "" {
		return wheres + sql.WhereRaws
	} else {
		return wheres[:len(wheres)-5]
	}
}

func (sql *SqlComponent) prepareUpdate(delimiter string) {
	fields := ""
	args := make([]interface{}, 0)

	if len(sql.Values) != 0 {

		for key, value := range sql.Values {
			fields += delimiter + key + delimiter + " = ?, "
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

	sql.Statement = "update " + sql.TableName + " set " + fields + sql.getWheres(delimiter)
}

func (sql *SqlComponent) prepareInsert(delimiter string) {
	fields := " ("
	quesMark := "("

	for key, value := range sql.Values {
		fields += delimiter + key + delimiter + ","
		quesMark += "?,"
		sql.Args = append(sql.Args, value)
	}
	fields = fields[:len(fields)-1] + ")"
	quesMark = quesMark[:len(quesMark)-1] + ")"

	sql.Statement = "insert into " + sql.TableName + fields + " values " + quesMark
}

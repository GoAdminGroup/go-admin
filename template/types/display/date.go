package display

import (
	"strconv"
	"time"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type Date struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("date", new(Date))
}

func (d *Date) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		format := args[0].(string)
		ts, _ := strconv.ParseInt(value.Value, 10, 64)
		tm := time.Unix(ts, 0)
		return tm.Format(format)
	}
}

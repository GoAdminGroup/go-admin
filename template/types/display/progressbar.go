package display

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type ProgressBar struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("progressbar", new(ProgressBar))
}

func (p *ProgressBar) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		param := args[0].([]types.FieldProgressBarData)
		style := "primary"
		size := "sm"
		max := 100
		if len(param) > 0 {
			if param[0].Style != "" {
				style = param[0].Style
			}
			if param[0].Size != "" {
				size = param[0].Size
			}
			if param[0].Max != 0 {
				max = param[0].Max
			}
		}
		base, _ := strconv.Atoi(value.Value)
		per := fmt.Sprintf("%.0f", float32(base)/float32(max)*100)
		return template.HTML(`
<div class="row" style="min-width: 100px;">
	<span class="col-sm-3" style="color:#777;width: 60px">` + per + `%</span>
	<div class="progress progress-` + size + ` col-sm-9" style="padding-left: 0;width: 100px;margin-left: -13px;">
		<div class="progress-bar progress-bar-` + style + `" role="progressbar" aria-valuenow="1" 
			aria-valuemin="0" aria-valuemax="` + strconv.Itoa(max) + `" style="width: ` + per + `%">
		</div>
	</div>
</div>`)
	}
}

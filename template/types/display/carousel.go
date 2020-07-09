package display

import (
	"html/template"
	"strconv"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type Carousel struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("carousel", new(Carousel))
}

func (c *Carousel) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		fn := args[0].(types.FieldGetImgArrFn)
		size := args[1].([]int)

		width := "300"
		height := "200"

		if len(size) > 0 {
			width = strconv.Itoa(size[0])
		}

		if len(size) > 1 {
			height = strconv.Itoa(size[1])
		}

		images := fn(value.Value)

		indicators := ""
		items := ""
		active := ""

		for i, img := range images {
			indicators += `<li data-target="#carousel-value-` + value.ID + `" data-slide-to="` +
				strconv.Itoa(i) + `" class=""></li>`
			if i == 0 {
				active = " active"
			} else {
				active = ""
			}
			items += `<div class="item` + active + `">
            <img src="` + img + `" alt="" 
style="max-width:` + width + `px;max-height:` + height + `px;display: block;margin-left: auto;margin-right: auto;" />
            <div class="carousel-caption"></div>
        </div>`
		}

		return template.HTML(`
<div id="carousel-value-` + value.ID + `" class="carousel slide" data-ride="carousel" width="` + width + `" height="` + height + `" 
style="padding: 5px;border: 1px solid #f4f4f4;background-color:white;width:` + width + `px;">
    <ol class="carousel-indicators">
		` + indicators + `
    </ol>
    <div class="carousel-inner">
       ` + items + `
    </div>
    <a class="left carousel-control" href="#carousel-value-` + value.ID + `" data-slide="prev">
        <span class="fa fa-angle-left"></span>
    </a>
    <a class="right carousel-control" href="#carousel-value-` + value.ID + `" data-slide="next">
        <span class="fa fa-angle-right"></span>
    </a>
</div>
`)
	}
}

package parameter

import (
	"fmt"
	"testing"
)

func TestGetParamFromUrl(t *testing.T) {
	fmt.Println(GetParamFromURL("/admin/info/user?__page=1&__pageSize=10&__sort=id&__sort_type=desc",
		1, "asc", "id"))
}

func TestParameters_PKs(t *testing.T) {
	pks := BaseParam().PKs()
	fmt.Println("pks", pks, "len", len(pks))
}

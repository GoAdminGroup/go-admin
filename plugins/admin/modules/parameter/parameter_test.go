package parameter

import (
	"fmt"
	"testing"
)

func TestGetParamFromUrl(t *testing.T) {
	fmt.Println(GetParamFromURL("/admin/info/user?__page=1&__pageSize=10&__sort=id&__sort_type=desc",
		true, 1, "id", "asc"))
}

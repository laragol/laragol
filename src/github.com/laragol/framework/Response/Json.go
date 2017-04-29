package Response

import (
	"app/Helpers"
	"fmt"

	"encoding/json"

	"github.com/qiangxue/fasthttp-routing"
)

func Json(c *routing.Context, v interface{}) (err error) {
	b, err := json.Marshal(v)

	if err != nil {
		Helpers.Error(c, err)
		return
	}

	fmt.Fprintf(c, string(b))

	return
}

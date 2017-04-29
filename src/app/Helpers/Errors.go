package Helpers

import (
	"fmt"
	"os"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Error(c *routing.Context, err error) {

	if os.Getenv("DEBUG") == "true" {
		fmt.Fprintf(c, err.Error())
	}

	fmt.Println(err.Error())

}

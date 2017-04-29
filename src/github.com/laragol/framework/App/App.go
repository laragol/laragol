package App

import (
	"app/Http"
	"database/migrations"
	"database/seeds"
	"fmt"
	"os"

	"github.com/laragol/framework/Database/Migration"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func Run() {
	var cmd string

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "listen"
	}
	switch cmd {
	case "migrate":
		Migrations.Migrate()
		Migration.Up()
	case "migrate:down":
		Migrations.Migrate()
		Migration.Down()
	case "db:seed":
		Seeds.Run()
	default:
		router := routing.New()
		Routes.Handle(router)
		fmt.Println("Listening in port 8080")
		panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
	}
}

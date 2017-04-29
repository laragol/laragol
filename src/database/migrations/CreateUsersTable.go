package Migrations

import (
	"app/Models"
	"fmt"

	"github.com/laragol/framework/Database/Connection"
)

type CreateUsersTable struct {
}

func (CreateUsersTable) Up() {
	db := Connection.Get()
	db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	db.DropTable(&Models.User{})
	db.CreateTable(&Models.User{})
	db.Model(&Models.User{}).AddForeignKey("country_id", "countries(id)", "RESTRICT", "RESTRICT")
}

func (CreateUsersTable) Down() {
	fmt.Println("CreateCountriesTable::down Not Implemented.")
}

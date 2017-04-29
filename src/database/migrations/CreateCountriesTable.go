package Migrations

import (
	"app/Models"
	"fmt"

	"github.com/laragol/framework/Database/Connection"
)

type CreateCountriesTable struct {
}

func (CreateCountriesTable) Up() {
	db := Connection.Get()
	db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	db.DropTable(&Models.Country{})
	db.CreateTable(&Models.Country{})
}

func (CreateCountriesTable) Down() {
	fmt.Println("CreateCountriesTable::down Not Implemented.")
}

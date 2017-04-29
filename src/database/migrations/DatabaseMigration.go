package Migrations

import . "github.com/laragol/framework/Database/Migration"

func Migrate() {
	Register(CreateCountriesTable{})
	Register(CreateUsersTable{})
}

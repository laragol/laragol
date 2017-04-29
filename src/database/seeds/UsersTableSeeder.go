package Seeds

import (
	"app/Models"

	"github.com/laragol/framework/Database/Connection"
)

type UsersTableSeeder struct {
}

func (UsersTableSeeder) Run() {
	db := Connection.Get()
	db.Create(&Models.User{Name: "Bob", Email: "bob@mail.com", CountryID: 1})
	db.Create(&Models.User{Name: "Alice", Email: "alice@mail.com", CountryID: 2})
}

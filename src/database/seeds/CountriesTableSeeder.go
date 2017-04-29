package Seeds

import (
	"app/Models"

	"github.com/laragol/framework/Database/Connection"
)

type CountriesTableSeeder struct {
}

func (CountriesTableSeeder) Run() {
	db := Connection.Get()
	db.Create(&Models.Country{Name: "Argentina"})
	db.Create(&Models.Country{Name: "Chile"})
	db.Create(&Models.Country{Name: "Uruguay"})
	db.Create(&Models.Country{Name: "Brazil"})
	db.Create(&Models.Country{Name: "Peru"})
}

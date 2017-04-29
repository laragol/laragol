package Seeds

import . "github.com/laragol/framework/Database/Seeder"

func Run() {
	Call(CountriesTableSeeder{})
	Call(UsersTableSeeder{})
}

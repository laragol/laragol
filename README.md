# Laravel clone in go

This is an ongoing experiment to help me learn go by trying to clone the developer experience of Laravel or more accurate Lumen.

I don't know if this is actually a good idea to build applications in this manner. But maybe is a way to present go to PHP or Rails developers.

### Features

* Live update on file change (makefile sugar) + [(cespare/reflex)](https://github.com/cespare/reflex)
* Dependencies folder vendor (GOPATH to /vendor)
* Http Server [(valyala/fasthttp)](https://github.com/valyala/fasthttp)
* Routing [(qiangxue/fasthttp-routing)](https://github.com/qiangxue/fasthttp-routing)
* ORM [(jinzhu/gorm)](https://github.com/jinzhu/gorm)
* Migration (only folder like atm, looking for libraries)
* Seeding (Basic)
* Examples (User & Country CRUD)

### Roadmap

* Templating (valyala/quicktemplate)
* Migration library
* More Examples
* Even More Examples
* Auth
* Config
* Commands
* Package Management [(Masterminds/glide)](https://github.com/Masterminds/glide)

### Folder structure

```bash
$ tree src
.
├── app
│   ├── Helpers
│   │   └── DB.go
│   ├── Http
│   │   ├── Controllers
│   │   │   ├── CountryController.go
│   │   │   └── UserController.go
│   │   └── routes.go
│   └── Models
│       ├── Country.go
│       └── User.go
├── database
│   ├── migrations
│   │   ├── 000001_create_countries_table.go
│   │   ├── 000002_create_users_table.go
│   │   └── migrations.go
│   └── seeds
│       ├── CountriesTableSeeder.go
│       ├── UsersTableSeeder.go
│       └── seeds.go
└── main.go

8 directories, 13 files
```

## Docker

```bash
docker-compose up
```

Server is listening on http://127.0.0.1:8080/

### Run migrations

```bash
docker-compose run api /code/bin/main migrate
```

### Run seeding

```bash
docker-compose run api /code/bin/main db:seed
```

Now you can see it in acion at:

http://127.0.0.1:8080/api/countries

http://127.0.0.1:8080/api/users

## Without Docker

Without docker yo need to install go and have MySQL up and running. Create a database "go_api" and root user with password "123456" or modify the file src/app/Helpers/DB.go with your MySQL credentials.

```bash
make install
make build-dev
./bin/main migrate
./bin/main db:seed
make dev #linux
make dev-mac #mac
```
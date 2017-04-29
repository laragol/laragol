package Migration

import (
	"fmt"
	"reflect"
)

type Migration interface {
	Up()
	Down()
}

var migrations []Migration

func Register(object Migration) {
	migrations = append(migrations, object)
}

func Up() {
	for _, migration := range migrations {
		fmt.Println(reflect.TypeOf(migration).Name())
		migration.Up()
	}
}

func Down() {
	for i := range migrations {
		fmt.Println(reflect.TypeOf(migrations[len(migrations)-1-i]).Name())
		migrations[i].Down()
	}
}

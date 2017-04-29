package Seeder

import (
	"fmt"
	"reflect"
)

type Seeder interface {
	Run()
}

func Call(class Seeder) {
	fmt.Println(reflect.TypeOf(class).Name())
	class.Run()
}

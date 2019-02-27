package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {

	var a MyStruct
	b := new(MyStruct)

	a.name = "liuxuewen"
	b.name = "yejianfeng"
	val := reflect.ValueOf(a).FieldByName("name")

	fmt.Println(val)
	valB := reflect.ValueOf(b).Elem().FieldByName("name")
	fmt.Println("--------------")
	fmt.Println(valB)
	fmt.Println("--------------////////////")
	//fmt.Println(reflect.ValueOf(a).FieldByName("name").CanSet())
	//fmt.Println(reflect.ValueOf(&(a.name)).Elem().CanSet())

}

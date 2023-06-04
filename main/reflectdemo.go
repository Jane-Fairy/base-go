package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

func reflectDemo() {
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}
	//使用 reflect.TypeOf() 函数可以获得任意值的类型对象
	typeof := reflect.TypeOf(cat{})
	fmt.Println(typeof.Name(), typeof.Kind())

	//获取常量的反射类型
	typeofZero := reflect.TypeOf(Zero)
	fmt.Println(typeofZero.Name(), typeofZero.Kind())

	//指针与指针指向的元素
	catPtr := &cat{}
	typeofPtr := reflect.TypeOf(catPtr)
	fmt.Printf("name:'%v' kind:'%v'\n", typeofPtr.Name(), typeofPtr.Kind())
	typeofPtr = typeofPtr.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeofPtr.Name(), typeofPtr.Kind())

	ins := cat{Name: "mimi", Type: 1}
	catins := reflect.TypeOf(ins)
	for i := 0; i < catins.NumField(); i++ {
		fieldType := catins.Field(i)
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)

		fmt.Println("==============================================================")
		if catType, ok := catins.FieldByName("Type"); ok {
			// 从tag中取出需要的tag
			fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name            string
	Age             int
	FavouriteColors []string // non-comparable field 不可比较
}

func main() {
	p0 := &Person{}
	p1 := &Person{
		Name:            "John",
		Age:             45,
		FavouriteColors: []string{"red", "blue", "green"},
	}

	// 千万注意要解引用后与类型零值做比较
	fmt.Println(reflect.DeepEqual(*p0, Person{})) // true
	fmt.Println(reflect.ValueOf(*p0).IsZero())    // false

	fmt.Println(reflect.DeepEqual(*p1, Person{})) // false
	// 【推荐这种方式判断】 reflect.ValueOf(x).IsZero()
	fmt.Println(reflect.ValueOf(*p1).IsZero()) // false
}

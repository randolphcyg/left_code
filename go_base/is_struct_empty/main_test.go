package main

import (
	"reflect"
	"testing"
)

var (
	p0 = &Person{}
	p1 = &Person{
		Name:            "John",
		Age:             45,
		FavouriteColors: []string{"red", "blue", "green"},
	}
)

func BenchmarkReflectDeepEqual1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(p0, Person{})
	}
}

func BenchmarkReflectDeepEqual2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(p1, Person{})
	}
}

func BenchmarkReflectValueOfIsZero1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(p0).IsZero()
	}
}

func BenchmarkReflectValueOfIsZero2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(p1).IsZero()
	}
}

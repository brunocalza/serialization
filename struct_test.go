package main

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

type MyStruct struct {
	a int
	b string
}

func TestStruct(t *testing.T) {
	_, err := os.OpenFile("file", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	myStruct := MyStruct{1, "hellodajlkasjlkdjsalkdjsalkasjdksajsldkjdsalkjdsalkjdlksajdaskljdaslkdaskljdaskljdaslkjdsalkajadadsadsadsadasdasasdasdasdasdsasadasdsasasdasdsasadasasasdadsadasdsasdldalkj"}
	fmt.Println(unsafe.Sizeof(myStruct))

}

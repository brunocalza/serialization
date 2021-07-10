package main

import (
	"fmt"
	"os"
	"testing"
)

func TestInteger(t *testing.T) {
	file, err := os.OpenFile("file", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	myNumber := uint32(100)
	fmt.Printf("%08b\n", myNumber)

	myNumberInBytes := IntToByteArray(myNumber)

	writtenBytes, err := file.WriteAt(myNumberInBytes, 0)
	fmt.Println(writtenBytes)
	if err != nil {
		panic(err)
	}
	file.Close()

	file, err = os.OpenFile("file", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	myNumberInBytes = make([]byte, 4)

	_, err = file.ReadAt(myNumberInBytes, 0)
	if err != nil {
		panic(err)
	}

	myNumber = ByteArrayToInt(myNumberInBytes)
	fmt.Println(myNumber)

	file.Close()
}

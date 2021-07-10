package main

import (
	"fmt"
	"os"

	"capnproto.org/go/capnp/v3"
	"github.com/brunocalza/serialization/books"
	"github.com/edsrzf/mmap-go"
)

func main() {
	file, err := os.OpenFile("file", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	mmap, _ := mmap.Map(file, mmap.RDWR, 0)
	defer mmap.Unmap()

	msg2, err := capnp.Unmarshal(mmap)
	if err != nil {
		panic(err)
	}

	book2, err := books.ReadRootBook(msg2)
	if err != nil {
		panic(err)
	}
	fmt.Println(book2.Title())
	fmt.Println(book2.PageCount())

	book2.SetTitle("Changed title")
}

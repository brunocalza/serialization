package capnprot

import (
	"os"
	"testing"

	"capnproto.org/go/capnp/v3"
	"github.com/brunocalza/serialization/books"
)

func TestBook(t *testing.T) {

	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.
	book, err := books.NewRootBook(seg)
	if err != nil {
		panic(err)
	}
	book.SetTitle("War and Peace")
	book.SetPageCount(1440)

	file, err := os.OpenFile("file", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	err = capnp.NewEncoder(file).Encode(msg)
	if err != nil {
		panic(err)
	}

}

using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("books");
$Go.import("github.com/brunocalza/serialization/books");

struct Book {
	title @0 :Text;
	pageCount @1 :Int32;
}
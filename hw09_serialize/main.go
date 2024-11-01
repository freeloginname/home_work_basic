package book

import (
	"encoding/json"

	"github.com/freeloginname/home_work_basic/hw09_serialize/book"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

func (book *Book) Marshaller() ([]byte, error) {
	return json.Marshal(book)
}

func (book *Book) Unmarshaller(data []byte) error {
	err := json.Unmarshal(data, book)
	if err != nil {
		return err
	}
	return nil
}

func SliceMarshaller(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func SliceUnmarshaller(data []byte) ([]Book, error) {
	var result []Book
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type BookProto struct {
	*book.Message
}

func (book *BookProto) Marshaller() ([]byte, error) {
	return json.Marshal(book)
}

// func main() {
// 	// Place your code here.
// 	book1 := Book{1, "title", "author", 1, 1, 1.0}

// 	bytes, _ := book1.Marshaller()
// 	fmt.Println(string(bytes))
// 	j := []byte(`{"Id":2,"Title":"title2","Author":"author2","Year":2,"Size":2,"Rate":2.0}`)
// 	var book2 Book
// 	_ = book2.Unmarshaller(j)
// 	fmt.Println(book2)

// 	result, _ := SliceMarshaller([]Book{{1, "title", "author", 1, 1, 1.0}, {2, "title2", "author2", 2, 2, 2.0}})
// 	fmt.Println(string(result))

// 	result2, _ := SliceUnmarshaller([]byte(`[{"Id":1,"Title":"title","Author":"author","Year":1,"Size":1,"Rate":1.0},
// 	{"Id":2,"Title":"title2","Author":"author2","Year":2,"Size":2,"Rate":2.0}]`))
// 	fmt.Println(result2)

// 	bookProto := BookProto{Message: book.Message{Id: 1, Title: "title", Author: "author", Year: 1, Size: 1, Rate: 1.0}}
// 	bytes, _ = bookProto.Marshaller()
// 	fmt.Println(string(bytes))

// 	pb := &book.Message{Id: 1, Title: "title", Author: "author", Year: 1, Size: 1, Rate: 1.0}
// 	bytes, _ = json.Marshal(pb)
// 	fmt.Println(string(bytes))
// }

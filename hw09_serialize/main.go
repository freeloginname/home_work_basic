package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Id     int
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
	var result []byte
	for book := range books {
		currentBook, err := books[book].Marshaller()
		if err != nil {
			return nil, err
		}
		result = append(result, currentBook...)
	}
	return result, nil
}

func SliceUnmarshaller(data []byte) ([]Book, error) {
	var result []Book
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	// Place your code here.
	var book Book = Book{1, "title", "author", 1, 1, 1.0}

	bytes, _ := book.Marshaller()
	fmt.Println(bytes)
	j := []byte(`{"Id":2,"Title":"title2","Author":"author2","Year":2,"Size":2,"Rate":2.0}`)
	var book2 Book
	_ = book2.Unmarshaller(j)
	fmt.Println(book2)

	result, _ := SliceMarshaller([]Book{{1, "title", "author", 1, 1, 1.0}, {2, "title2", "author2", 2, 2, 2.0}})
	fmt.Println(result)

	result2, _ := SliceUnmarshaller([]byte(`[{"Id":1,"Title":"title","Author":"author","Year":1,"Size":1,"Rate":1.0},{"Id":2,"Title":"title2","Author":"author2","Year":2,"Size":2,"Rate":2.0}]`))
	fmt.Println(result2)

}

package main

import (
	"fmt"
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (book *Book) GetBookID() int {
	return book.id
}

func (book *Book) GetBookTitle() string {
	return book.title
}

func (book *Book) GetBookAuthor() string {
	return book.author
}

func (book *Book) GetBookYear() int {
	return book.year
}

func (book *Book) GetBookSize() int {
	return book.size
}

func (book *Book) GetBookRate() float32 {
	return book.rate
}

func (book *Book) SetBookID(id int) {
	book.id = id
}

func (book *Book) SetBookTitle(title string) {
	book.title = title
}

func (book *Book) SetBookAuthor(author string) {
	book.author = author
}

func (book *Book) SetBookYear(year int) {
	book.year = year
}

func (book *Book) SetBookSize(size int) {
	book.size = size
}

func (book *Book) SetBookRate(rate float32) {
	book.rate = rate
}

type BookComparer struct {
	// firstBook Book
	// secondBook Book
	compareBy ComparisonType
}

type ComparisonType int

const (
	Year ComparisonType = iota
	Size ComparisonType = iota
	Rate ComparisonType = iota
)

func NewBookComparer(compareBy ComparisonType) *BookComparer {
	return &BookComparer{
		// firstBook: firstBook,
		// secondBook: secondBook,
		compareBy: compareBy,
	}
}

func (comparer *BookComparer) CompareBooks(firstBook Book, secondBook Book) bool {
	switch comparer.compareBy {
	case Year:
		if firstBook.year > secondBook.year {
			return true
		}
	case Size:
		if firstBook.size > secondBook.size {
			return true
		}
	case Rate:
		if firstBook.rate > secondBook.rate {
			return true
		}
	}
	return false
}

// ID, Title, Author, Year, Size, Rate

func main() {
	fmt.Println("run some tests")
	var firstBook Book
	var secondBook Book
	firstBook.SetBookID(1)
	firstBook.SetBookTitle("aaa")
	firstBook.SetBookAuthor("aaaaa")
	firstBook.SetBookYear(1990)
	firstBook.SetBookSize(10)
	firstBook.SetBookRate(2.0)
	secondBook.SetBookID(2)
	secondBook.SetBookTitle("bb")
	secondBook.SetBookAuthor("bbb")
	secondBook.SetBookYear(1991)
	secondBook.SetBookSize(20)
	secondBook.SetBookRate(5.0)

	yearComparer := NewBookComparer(Year)
	compareByYear := yearComparer.CompareBooks(firstBook, secondBook)
	fmt.Printf("Year %t \n", compareByYear)

	sizeComparer := NewBookComparer(Size)
	compareBySize := sizeComparer.CompareBooks(firstBook, secondBook)
	fmt.Printf("Size %t \n", compareBySize)

	RateComparer := NewBookComparer(Rate)
	compareByRate := RateComparer.CompareBooks(firstBook, secondBook)
	fmt.Printf("Size %t \n", compareByRate)
}

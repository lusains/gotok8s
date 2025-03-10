package iterator

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
	Current() interface{}
	Reset()
}

// Container 容器接口
type Container interface {
	CreateIterator() Iterator
}

// BookShelf 书架（具体容器）
type BookShelf struct {
	books []*Book
}

func NewBookShelf() *BookShelf {
	return &BookShelf{
		books: make([]*Book, 0),
	}
}

func (b *BookShelf) AddBook(book *Book) {
	b.books = append(b.books, book)
}

func (b *BookShelf) CreateIterator() Iterator {
	return NewBookShelfIterator(b)
}

// Book 书籍
type Book struct {
	Name   string
	Author string
}

func NewBook(name, author string) *Book {
	return &Book{
		Name:   name,
		Author: author,
	}
}

// BookShelfIterator 书架迭代器（具体迭代器）
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

func (i *BookShelfIterator) HasNext() bool {
	return i.index < len(i.bookShelf.books)
}

func (i *BookShelfIterator) Next() interface{} {
	if i.HasNext() {
		book := i.bookShelf.books[i.index]
		i.index++
		return book
	}
	return nil
}

func (i *BookShelfIterator) Current() interface{} {
	if i.index < len(i.bookShelf.books) {
		return i.bookShelf.books[i.index]
	}
	return nil
}

func (i *BookShelfIterator) Reset() {
	i.index = 0
}

// ReverseIterator 反向迭代器接口
type ReverseIterator interface {
	HasPrevious() bool
	Previous() interface{}
	Current() interface{}
	Reset()
}

// BookShelfReverseIterator 书架反向迭代器
type BookShelfReverseIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfReverseIterator(bookShelf *BookShelf) *BookShelfReverseIterator {
	return &BookShelfReverseIterator{
		bookShelf: bookShelf,
		index:     len(bookShelf.books) - 1,
	}
}

func (i *BookShelfReverseIterator) HasPrevious() bool {
	return i.index >= 0
}

func (i *BookShelfReverseIterator) Previous() interface{} {
	if i.HasPrevious() {
		book := i.bookShelf.books[i.index]
		i.index--
		return book
	}
	return nil
}

func (i *BookShelfReverseIterator) Current() interface{} {
	if i.index >= 0 {
		return i.bookShelf.books[i.index]
	}
	return nil
}

func (i *BookShelfReverseIterator) Reset() {
	i.index = len(i.bookShelf.books) - 1
}

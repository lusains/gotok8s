package iterator

import "testing"

func TestIterator(t *testing.T) {
	// 创建书架并添加书籍
	bookShelf := NewBookShelf()
	bookShelf.AddBook(NewBook("Go语言设计模式", "作者1"))
	bookShelf.AddBook(NewBook("Go语言编程", "作者2"))
	bookShelf.AddBook(NewBook("Go Web编程", "作者3"))

	// 测试正向迭代器
	iterator := bookShelf.CreateIterator()
	count := 0
	for iterator.HasNext() {
		book := iterator.Next().(*Book)
		if book == nil {
			t.Error("迭代器返回了nil")
		}
		count++
	}
	if count != 3 {
		t.Errorf("迭代器遍历数量错误，期望3，得到%d", count)
	}

	// 测试迭代器重置
	iterator.Reset()
	if !iterator.HasNext() {
		t.Error("重置后迭代器应该可以继续遍历")
	}

	// 测试反向迭代器
	reverseIterator := NewBookShelfReverseIterator(bookShelf)
	books := make([]string, 0)
	for reverseIterator.HasPrevious() {
		book := reverseIterator.Previous().(*Book)
		books = append(books, book.Name)
	}

	// 验证反向遍历顺序
	if books[0] != "Go Web编程" || books[2] != "Go语言设计模式" {
		t.Error("反向迭代器顺序错误")
	}

	// 测试当前元素
	reverseIterator.Reset()
	currentBook := reverseIterator.Current().(*Book)
	if currentBook.Name != "Go Web编程" {
		t.Error("Current()方法返回错误")
	}
}

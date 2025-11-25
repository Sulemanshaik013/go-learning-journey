package main

import (
	"fmt"
	"strconv"
)

type Borrowable interface {
	Borrow(memberId string) bool
	Return() bool
	IsAvailable() bool
}

type Displayable interface {
	Display() string
}

type Book struct {
	isbn       string
	title      string
	author     string
	available  bool
	borrowedBy string
}

func (b Book) Details() string {
	status := "Available"
	if !b.available {
		status = "borrowed by: " + b.borrowedBy
	}
	return b.title + " written by " + b.author + " is " + status
}

func (b Book) Display() string {
	return b.title + " written by " + b.author
}

func (b Book) IsAvailable() bool {
	return b.available
}

func (b *Book) Borrow(memberId string) bool {
	if !b.available {
		return false
	}
	b.available = false
	b.borrowedBy = memberId
	return true
}

func (b *Book) Return() bool {
	if b.available {
		return false
	}
	b.available = true
	b.borrowedBy = ""
	return true
}

type Member struct {
	id            string
	name          string
	email         string
	borrowedBooks []string
}

func (m Member) Info() string {
	return m.name + " has registered with email " + m.email + ", and borrowed " + strconv.FormatInt(int64(m.BorrowedCount()), 10) + "Books"
}

func (m Member) Display() string {
	return m.name + " has registered with email " + m.email
}

func (m Member) BorrowedCount() int {
	return len(m.borrowedBooks)
}

func (m Member) CanBorrow() bool {
	return m.BorrowedCount() < 3
}

func (m *Member) AddBorrowedBook(isbn string) {
	m.borrowedBooks = append(m.borrowedBooks, isbn)
}

func (m *Member) RemoveBorrowedBook(isbn string) bool {
	for idx, val := range m.borrowedBooks {
		if val == isbn {
			m.borrowedBooks = append(m.borrowedBooks[:idx], m.borrowedBooks[idx+1:]...)
			return true
		}
	}
	return false
}

type Library struct {
	name    string
	books   []Book
	members []Member
}

func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

func (l *Library) AddMember(member Member) {
	l.members = append(l.members, member)
}

func (l *Library) FindBook(isbn string) *Book {
	for idx := range l.books {
		if l.books[idx].isbn == isbn {
			return &l.books[idx]
		}
	}
	return nil
}

func (l Library) FindMember(id string) *Member {
	for idx := range l.members {
		if l.members[idx].id == id {
			return &l.members[idx]
		}
	}
	return nil
}

func (l Library) BorrowBook(isbn, memberId string) {
	book := l.FindBook(isbn)
	if book == nil {
		fmt.Println("Book Not found")
		return
	}
	member := l.FindMember(memberId)
	if member == nil {
		fmt.Println("Member Not found")
		return
	}
	if member.CanBorrow() && book.IsAvailable() && book.Borrow(memberId) {
		member.AddBorrowedBook(isbn)
		fmt.Println("Success BorrowBook")
		return
	}
	fmt.Println("Failed BorrowBook")
}

func (l Library) ReturnBook(isbn, memberId string) {
	book := l.FindBook(isbn)
	if book == nil {
		fmt.Println("Book Not found")
		return
	}
	member := l.FindMember(memberId)
	if member == nil {
		fmt.Println("Member Not found")
		return
	}
	if book.borrowedBy != memberId {
		fmt.Println("Failed ReturnBook")
		return
	}

	if book.Return() && member.RemoveBorrowedBook(isbn) {
		fmt.Println("Success ReturnBook")
		return
	}
	fmt.Println("Failed ReturnBook")

}

func (l Library) ListAvailableBooks() {
	fmt.Println("Available Books are : ")
	for _, book := range l.books {
		fmt.Println(book.Details())
	}
}

func (l Library) ListMemberBooks(memberId string) {
	m := l.FindMember(memberId)
	if len(m.borrowedBooks) == 0 {
		fmt.Println("No books Borrowed by ", m.name)
		return
	}
	for _, isbn := range m.borrowedBooks {
		book := l.FindBook(isbn)
		if book != nil {
			fmt.Println(book.Details())
		}
	}
}

func (l Library) Display() {
	for _, b := range l.books {
		ShowDetails(b)
	}
	for _, m := range l.members {
		ShowDetails(m)
	}
}

func ShowDetails(d Displayable) {
	fmt.Println(d.Display())
}

func main() {
	library := &Library{name: "Central Library"}

	library.AddBook(Book{isbn: "123", title: "Hey Java", author: "JAVA", available: true})
	library.AddBook(Book{isbn: "234", title: "Hey Go", author: "GO", available: true})
	library.AddBook(Book{isbn: "345", title: "Hey HTML", author: "HTML", available: true})
	library.AddBook(Book{isbn: "456", title: "Hey JavaScript", author: "JS", available: true})
	library.AddBook(Book{isbn: "567", title: "Hey Ruby", author: "RUBY", available: true})

	library.AddMember(Member{id: "M1", name: "Vishal", email: "vishal@gmail.com"})
	library.AddMember(Member{id: "M2", name: "John", email: "john@gmail.com"})
	library.AddMember(Member{id: "M3", name: "Karna", email: "karna@gmail.com"})
	library.AddMember(Member{id: "M4", name: "Samantha", email: "samantha@gmail.com"})

	library.Display()

	fmt.Println("Borrow operations")
	library.BorrowBook("123", "M1")
	library.BorrowBook("234", "M1")
	library.BorrowBook("456", "M1")

	//Invalid
	library.BorrowBook("345", "M1")
	library.BorrowBook("234", "M1")

	library.ListAvailableBooks()

	library.ReturnBook("123", "M1")
	library.ReturnBook("234", "M1")

	//invalid
	library.ReturnBook("345", "M1")

	library.ListMemberBooks("M1")

	library.ListAvailableBooks()

	var borrowables []Borrowable

	book1 := library.FindBook("234")
	if book1 != nil {
		borrowables = append(borrowables, book1)
	}
	book2 := library.FindBook("456")
	if book2 != nil {
		borrowables = append(borrowables, book2)
	}

	for i, item := range borrowables {
		fmt.Printf("Item %d Status %v\n", i+1, item.IsAvailable())
		switch item.(type) {
		case *Book:
			fmt.Println("Type : Book")
		default:
			fmt.Println("Type : Unknown")
		}
	}
}

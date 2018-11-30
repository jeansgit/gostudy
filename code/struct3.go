package main

import "fmt"
type Books struct{
	title string
	author string
	subject string
	bookid int
}
func main() {
	var book1 Books
	var book2 Books

	book1.title = "test"
	book1.author = "admin"
	book1.subject = "jean"
	book1.bookid = 100

	book2.title = "i"
	book2.author = "am"
	book2.subject = "bob"
	book2.bookid = 200

	printbook(&book1)
	printbook(&book2)
}
func printbook(book *Books) {
	fmt.Printf("%s-%s-%s-%d\n",book.title,book.author,book.subject,book.bookid)
}


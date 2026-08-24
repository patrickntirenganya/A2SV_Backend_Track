package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	Library services.LibraryManager
	Scanner *bufio.Scanner
}

func NewLibraryController(library services.LibraryManager) *LibraryController {
	return &LibraryController{
		Library: library,
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *LibraryController) Start() {
	for {
		c.printMenu()

		choice := c.readInt("Choose an option: ")

		switch choice {
		case 1:
			c.addBook()
		case 2:
			c.removeBook()
		case 3:
			c.borrowBook()
		case 4:
			c.returnBook()
		case 5:
			c.listAvailableBooks()
		case 6:
			c.listBorrowedBooks()
		case 7:
			c.addMember()
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}

		fmt.Println()
	}
}

func (c *LibraryController) printMenu() {
	fmt.Println("=================================")
	fmt.Println("   LIBRARY MANAGEMENT SYSTEM")
	fmt.Println("=================================")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Member's Borrowed Books")
	fmt.Println("7. Register Member")
	fmt.Println("0. Exit")
	fmt.Println("=================================")
}

func (c *LibraryController) addBook() {
	fmt.Println("\n--- Add Book ---")

	id := c.readInt("Book ID: ")
	title := c.readString("Title: ")
	author := c.readString("Author: ")

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	c.Library.AddBook(book)

	fmt.Println("Book added successfully.")
}

func (c *LibraryController) removeBook() {
	fmt.Println("\n--- Remove Book ---")

	id := c.readInt("Book ID: ")

	c.Library.RemoveBook(id)

	fmt.Println("Book removed successfully.")
}

func (c *LibraryController) borrowBook() {
	fmt.Println("\n--- Borrow Book ---")

	bookID := c.readInt("Book ID: ")
	memberID := c.readInt("Member ID: ")

	err := c.Library.BorrowBook(bookID, memberID)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book borrowed successfully.")
}

func (c *LibraryController) returnBook() {
	fmt.Println("\n--- Return Book ---")

	bookID := c.readInt("Book ID: ")
	memberID := c.readInt("Member ID: ")

	err := c.Library.ReturnBook(bookID, memberID)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book returned successfully.")
}

func (c *LibraryController) listAvailableBooks() {
	fmt.Println("\n--- Available Books ---")

	books := c.Library.ListAvailableBooks()

	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}

	for _, book := range books {
		fmt.Printf(
			"ID: %d | Title: %s | Author: %s | Status: %s\n",
			book.ID,
			book.Title,
			book.Author,
			book.Status,
		)
	}
}

func (c *LibraryController) listBorrowedBooks() {
	fmt.Println("\n--- Borrowed Books ---")

	memberID := c.readInt("Member ID: ")

	books := c.Library.ListBorrowedBooks(memberID)

	if len(books) == 0 {
		fmt.Println("This member has no borrowed books.")
		return
	}

	for _, book := range books {
		fmt.Printf(
			"ID: %d | Title: %s | Author: %s | Status: %s\n",
			book.ID,
			book.Title,
			book.Author,
			book.Status,
		)
	}
}

func (c *LibraryController) addMember() {
	fmt.Println("\n--- Register Member ---")

	id := c.readInt("Member ID: ")
	name := c.readString("Member Name: ")

	member := models.Member{
		ID:   id,
		Name: name,
	}

	if library, ok := c.Library.(*services.Library); ok {
		library.AddMember(member)
		fmt.Println("Member registered successfully.")
	} else {
		fmt.Println("Unable to register member.")
	}
}

func (c *LibraryController) readString(prompt string) string {
	fmt.Print(prompt)
	c.Scanner.Scan()

	return strings.TrimSpace(c.Scanner.Text())
}

func (c *LibraryController) readInt(prompt string) int {
	for {
		value := c.readString(prompt)

		number, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		return number
	}
}

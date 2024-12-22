package main

//Exercise 1
import (
  "fmt"
)

// Book struct to define book properties
type Book struct {
  ID         string
  Title      string
  Author     string
  IsBorrowed bool
}

// Library struct to manage book collection
type Library struct {
  books map[string]Book
}

// NewLibrary function to initialize a new Library instance
func NewLibrary() *Library {
  return &Library{books: make(map[string]Book)}
}

// AddBook method to add a new book to the library collection
func (l *Library) AddBook(book Book) {
  l.books[book.ID] = book
}

// BorrowBook method to borrow a book by updating IsBorrowed to true
func (l *Library) BorrowBook(id string) error {
  book, ok := l.books[id]
  if !ok {
    return fmt.Errorf("book with ID %s not found", id)
  }
  if book.IsBorrowed {
    return fmt.Errorf("book with ID %s is already borrowed", id)
  }
  book.IsBorrowed = true
  l.books[id] = book
  return nil
}

// ReturnBook method to return a borrowed book by updating IsBorrowed to false
func (l *Library) ReturnBook(id string) error {
  book, ok := l.books[id]
  if !ok {
    return fmt.Errorf("book with ID %s not found", id)
  }
  if !book.IsBorrowed {
    return fmt.Errorf("book with ID %s is already returned", id)
  }
  book.IsBorrowed = false
  l.books[id] = book
  return nil
}

// ListBooks method to print all books in the library
func (l *Library) ListBooks() {
  fmt.Println("** Library Book Collection**")
  fmt.Println("ID\tTitle\t\tAuthor\tBorrowed")
  fmt.Println("-----------------------------------------")
  for _, book := range l.books {
    fmt.Printf("%s\t%s\t\t%s\t%t\n", book.ID, book.Title, book.Author, book.IsBorrowed)
  }
}

func main() {
  // Create a new library instance
  library := NewLibrary()

  // Add some books to the library
  library.AddBook(Book{ID: "1", Title: "Волоколомское Шоссе", Author: "Бауыржан Момышулы", IsBorrowed: false})
  library.AddBook(Book{ID: "2", Title: "Путь Абая", Author: "Мухтар Ауезов", IsBorrowed: false})
  library.AddBook(Book{ID: "3", Title: "Война и Мир", Author: "Лев Толстой", IsBorrowed: false})

  // Simulate library usage
  var choice int
  for {
    fmt.Println("\nMenu:")
    fmt.Println("1. Add Book")
    fmt.Println("2. Borrow Book")
    fmt.Println("3. Return Book")
    fmt.Println("4. List Books")
    fmt.Println("5. Exit")
    fmt.Print("Enter your choice: ")
    fmt.Scanf("%d", &choice)

    switch choice {
    case 1:
      var id, title, author string
      fmt.Println("Enter book details:")
      fmt.Print("ID: ")
      fmt.Scanf("%s", &id)
      fmt.Print("Title: ")
      fmt.Scanf("%s", &title)
      fmt.Print("Author: ")
      fmt.Scanf("%s", &author)
      library.AddBook(Book{ID: id, Title: title, Author: author, IsBorrowed: false})
      fmt.Println("Book added successfully!")
    case 2:
      var id string
      fmt.Print("Enter book ID to borrow: ")
      fmt.Scanf("%s", &id)
      err := library.BorrowBook(id)
      if err != nil {
        fmt.Println(err.Error())
      } else {
        fmt.Println("Book borrowed successfully!")
      }
    case 3:
      var id string
      fmt.Print("Enter book ID to return: ")
      fmt.Scanf("%s", &id)
      err := library.ReturnBook(id)
      if err != nil {
        fmt.Println(err.Error())
      } else {
        fmt.Println("Book returned successfully!")
      }
    case 4:
      library.ListBooks()
    case 5:
      fmt.Println("Exiting the program...")
      return
    default:
      fmt.Println("Invalid choice. Please try again.")
    }
  }
}

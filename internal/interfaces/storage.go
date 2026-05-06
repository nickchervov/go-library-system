package interfaces

import (
	"time"
)

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Year      string `json:"year"`
	Genre     string `json:"genre"`
	Pages     int    `json:"pages"`
	Available bool   `json:"available"`
}

type LoanInfo struct {
	BookID     int        `json:"book_id"`
	MemberID   string     `json:"member_id"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
}

type Library struct {
	Books   map[int]*Book
	Members map[string][]string // [Член библиотеки][]id_книг
	Loans   map[string]LoanInfo
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*Book),
		Members: make(map[string][]string),
		Loans:   make(map[string]LoanInfo),
	}
}

func (l *Library) AddBook(b Book) {

}

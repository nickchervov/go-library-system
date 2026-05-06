package models

import (
	"errors"
	"time"

	"github.com/nickchervov/go-library-system/internal/utils"
)

type Book struct {
	ID        int    `json:"id" validate:"omitempty"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required,alpha"`
	ISBN      string `json:"isbn" validate:"required,numeric,min=13,max=13"`
	Year      string `json:"year" validate:"required,numeric"`
	Genre     string `json:"genre" validate:"required,alpha"`
	Pages     int    `json:"pages" validate:"required,numeric"`
	Available bool   `json:"available" validate:"required,alpha"`
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

var (
	nextId = 1
	Lib    = NewLibrary()
)

func (l *Library) AddBook(b Book) error {
	if err := utils.Validator.Struct(b); err != nil {
		return errors.New("ошибка при валидации данных")
	}
	if b.ID == 0 {
		b.ID = nextId
		nextId++
	}
	l.Books[b.ID] = &b
	return nil
}

func (l *Library) ListBook(genre string) map[int]*Book {
	listByGenreBooks := make(map[int]*Book)
	if genre == "" {
		return l.Books
	}
	for k, v := range l.Books {
		if v.Genre == genre {
			listByGenreBooks[k] = v
		}
	}
	return listByGenreBooks
}

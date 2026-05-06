package models

import (
	"fmt"
	"time"

	"github.com/nickchervov/go-library-system/internal/utils"
)

type Book struct {
	ID        int    `json:"id" validate:"omitempty"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required,alpha"`
	ISBN      string `json:"isbn" validate:"required,numeric,min=13,max=13"`
	Year      int    `json:"year" validate:"required,numeric"`
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
	Members map[string][]int // [Член библиотеки][]id_книг
	Loans   map[string]LoanInfo
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*Book),
		Members: make(map[string][]int),
		Loans:   make(map[string]LoanInfo),
	}
}

var (
	nextId = 1
	Lib    = NewLibrary()
)

func (l *Library) AddBook(b Book) error {
	if err := utils.Validator.Struct(b); err != nil {
		return fmt.Errorf("ошибка при валидации данных: %v", err)
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

func (l *Library) GetMostPopularGenre() (string, int) {
	genreCount := make(map[string]int)
	for _, v := range l.Books {
		genreCount[v.Genre]++
	}

	var mostPopularGenre string
	var biggestCount int
	for k, v := range genreCount {
		if v >= biggestCount {
			biggestCount = v
			mostPopularGenre = k
		}
	}

	return mostPopularGenre, biggestCount
}

func (l *Library) CalculateReadTime(memberId string) time.Duration {
	var readTime float64
	for _, bookIDs := range l.Members[memberId] {
		if book, exist := l.Books[bookIDs]; exist {
			//1 страница = 30 минут чтения
			readTime += float64(book.Pages) * 0.5 / 24.0
		}
	}
	return time.Duration(readTime * 24)
}

func (l *Library) FindOverdueLoans() []LoanInfo {
	var overdueLoans []LoanInfo
	for _, v := range l.Loans {
		if v.ReturnDate != nil && v.ReturnDate.Before(time.Now()) {
			overdueLoans = append(overdueLoans, v)
		}
	}
	return overdueLoans
}

func (l *Library) GenerateMonthlyReport(start, end time.Time) (map[string]interface{}, error) {
	report := make(map[string]interface{})
	report["totalBooks"] = len(l.Books)
	report["activeMembers"] = len(l.Members)
	report["overdueLoans"] = len(l.FindOverdueLoans())

	return report, nil
}

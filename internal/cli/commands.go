package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nickchervov/go-library-system/internal/models"
)

var scanner = bufio.NewScanner(os.Stdin)

func Input(title string) string {
	fmt.Print(title)

	scanner.Scan()

	return scanner.Text()
}

func Add() {
	var book models.Book

	var id int
	var err error
	for {
		inpId := Input("Введите id книги: ")
		if inpId != "" {
			id, err = strconv.Atoi(inpId)
			if err != nil {
				fmt.Printf("ошибка при конвертации id: %v\n", err)
				continue
			}
		}
		break
	}
	book.ID = id

	var title string
	for {
		title = Input("Введите название книги: ")
		if title == "" {
			fmt.Println("Название не может быть пустым")
			continue
		}
		break
	}
	book.Title = title

	var author string
	for {
		author = Input("Введите автора книги: ")
		if author == "" {
			fmt.Println("Поле автор не может быть пустым")
			continue
		}
		break
	}
	book.Author = author

	var ISBN string
	for {
		ISBN = Input("Введите ISBN книги: ")
		if ISBN == "" {
			fmt.Println("ISBN книги не может быть пустым")
			continue
		}
		if len(ISBN) != 13 {
			fmt.Println("ISBN книги не может быть меньше или больше 13 символов")
			continue
		}
		break
	}
	book.ISBN = ISBN

	var year int
	for {
		year, err = strconv.Atoi(Input("Введите год выпуска книги: "))
		if err != nil {
			fmt.Printf("ошибка при конвертации года: %v\n", err)
			continue
		}
		break
	}
	book.Year = year

	var genre string
	for {
		genre = Input("Введите жанр книги: ")
		if genre == "" {
			fmt.Println("Поле жанра не может быть пустым")
			continue
		}
		break
	}
	book.Genre = genre

	var pages int
	for {
		pages, err = strconv.Atoi(Input("Введите количество страниц книги: "))
		if err != nil {
			fmt.Printf("ошибка при конвертации страниц: %v\n", err)
			continue
		}
		break
	}
	book.Pages = pages

	if err := models.Lib.AddBook(book); err != nil {
		fmt.Printf("ошибка при добавлении книги: %v\n", err)
		return
	}
	fmt.Println("Книга успешно добавлена!")
}

func List() {
	genre := Input("Введите жанр или оставьте поле пустым если хотите вывести все книги: ")
	list := models.Lib.ListBook(genre)
	if len(list) == 0 {
		fmt.Println("Книг с таким жанром нет в наличии.")
		return
	}
	for _, v := range list {
		fmt.Printf("ID: %d. Название: %s. Автор: %s. ISBN: %s. Year: %d. Genre: %s. Pages: %d. Available: %t\n",
			v.ID, v.Title, v.Author, v.ISBN, v.Year, v.Genre, v.Pages, v.Available)
	}
}

func Borrow() {
	var member string
	var bookId int
	var returnDate string
	var err error

	for {
		member = Input("Введите ФИО бронирующего человека: ")
		if member == "" {
			fmt.Println("ФИО не может быть пустым")
			continue
		}
		if strings.ContainsAny(member, "0123456789") {
			fmt.Println("ФИО не может содержать цифры")
			continue
		}
		bookId, err = strconv.Atoi(Input("Введите id книги: "))
		if err != nil {
			fmt.Println("ID должно быть числом")
			continue
		}
		returnDate = Input("Введите дату возвращения книги в формате (ДД.ММ.ГГГГ): ")
		break
	}

	if err := models.Lib.BorrowBook(member, bookId, returnDate); err != nil {
		fmt.Printf("ошибка при создании брони: %v", err)
		return
	}
	fmt.Println("Бронь успешно создана!")
}

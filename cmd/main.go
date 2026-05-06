package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nickchervov/go-library-system/internal/cli"
	"github.com/nickchervov/go-library-system/internal/models"
	"github.com/nickchervov/go-library-system/internal/storage"
)

func RunProgram() bool {
	command := cli.Input("Введите команду: ")
	switch strings.TrimSpace(command) {
	case "add":
		fmt.Println()
		cli.Add()
		fmt.Println()
	case "list":
		fmt.Println()
		cli.List()
		fmt.Println()
	case "borrow":
		fmt.Println()
		cli.Borrow()
		fmt.Println()
	case "return":
		if err := storage.SaveToFile("library.json"); err != nil {
			log.Fatalf("ошибка при записи данных в файл: %v", err)
		}
		fmt.Println("До свидания!")
		return false
	default:
		fmt.Printf("команды с названием %s не найдено\n", command)
		return false
	}
	return true
}

func main() {
	fmt.Println("Это CLI приложение для управления библиотекой.")
	fmt.Println()

	var err error
	models.Lib, err = storage.LoadFromFile("library.json")
	if err != nil {
		log.Fatalf("ошибка при чтении данных из файла %v", err)
	}

	for RunProgram() {
	}
}

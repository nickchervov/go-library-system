package main

import (
	"fmt"
	"strings"

	"github.com/nickchervov/go-library-system/internal/cli"
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

	for RunProgram() {
	}
}

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Бесконечный цикл для чтения команд пользователя.
	for {
		fmt.Print("> ") // Вывод приглашения к вводу команды.
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			if err == io.EOF {
				os.Exit(0) // Завершаем программу при достижении конца ввода.
			}
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input) // Удаление лишних пробелов.

		args := strings.Fields(input) // Разделение строки на аргументы.
		if len(args) == 0 {
			continue // Игнорируем пустые строки.
		}

		// Обработка встроенных команд.
		switch args[0] {
		case "cd":
			if len(args) > 1 {
				err := os.Chdir(args[1]) // Меняем директорию на указанную.
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			} else {
				home, err := os.UserHomeDir() // Переходим в домашнюю директорию.
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				err = os.Chdir(home)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		case "pwd":
			dir, err := os.Getwd() // Получаем текущую директорию.
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(dir) // Выводим текущую директорию.
			}
		case "echo":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " ")) // Выводим аргументы команды echo.
			}
		case "kill":
			if len(args) > 1 {
				pid := args[1]
				cmd := exec.Command("kill", pid) // Выполняем команду kill для указанного PID.
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		case "ps":
			cmd := exec.Command("ps", "aux") // Выполняем команду ps aux.
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		default:
			cmd := exec.Command(args[0], args[1:]...) // Выполняем внешнюю команду с аргументами.
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}


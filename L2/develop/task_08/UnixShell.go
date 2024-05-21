package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input)

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) > 1 {
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			} else {
				home, err := os.UserHomeDir()
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
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(dir)
			}
		case "echo":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			}
		case "kill":
			if len(args) > 1 {
				pid := args[1]
				cmd := exec.Command("kill", pid)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		case "ps":
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
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

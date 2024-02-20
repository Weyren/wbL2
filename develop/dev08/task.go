package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	switch parts[0] {

	case "cd":
		if len(parts) < 2 {
			return errors.New("cd: no directory")
		}
		err := os.Chdir(parts[1])
		return err

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, dir)

	case "echo":
		cmd := exec.Command(parts[0], parts[1:]...)
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(output))

	case "kill":
		if len(parts) < 2 {
			return errors.New("kill: no pid")
		}
		cmd := exec.Command("kill", parts[1])
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(output))
	case "ps":
		cmd := exec.Command("ps", parts[1:]...)
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(output))
	default:

	}
	return nil
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigChan
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

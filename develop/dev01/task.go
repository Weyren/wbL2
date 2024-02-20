package dev01

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку .
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

// PrintTimeNTP выводит точное время используя NTP
func PrintTimeNTP() {
	timeNTP, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_ = fmt.Errorf(err.Error(), stderr)
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(stdout, "NTP:", timeNTP)

}

// PrintTimeLocal выводит локальное время
func PrintTimeLocal() {
	localTime := time.Now()
	_, _ = fmt.Fprintln(stdout, "Local:", localTime)
}

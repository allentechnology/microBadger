package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

func getUsername() {
	fmt.Print("Username: ")
	reader := bufio.NewReader(os.Stdin)
	providedUsername, err := reader.ReadBytes('\n')
	if err != nil {
		return
	}
	fmt.Print("Password: ")
	providedPassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}
	*username = strings.Trim(string(providedUsername), "\n")
	*password = strings.Trim(string(providedPassword), "\n")
	fmt.Println()
}

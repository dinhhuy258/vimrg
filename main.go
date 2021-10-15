package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var commandArgs = []string{
	"-F",
	"--column",
	"--line-number",
	"--no-heading",
	"--color=always",
	"--smart-case",
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	query := os.Args[1]
	elements := strings.SplitN(query, ">", 2)

	if len(elements) == 1 {
		commandArgs = append(commandArgs, "--", elements[0])
	} else {
		globs := strings.Fields(elements[1])

		for _, glob := range globs {
			commandArgs = append(commandArgs, "--iglob", glob)
		}

		commandArgs = append(commandArgs, "--", strings.Trim(elements[0], " "))
	}

	output, _ := exec.Command("rg", commandArgs...).Output()

	fmt.Printf(string(output))
}

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// ReadInput read user text input with prompt
func ReadInput(label string) string {
	fmt.Print(label)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// PromptEnter ask the user to press Enter before continuing
func PromptEnter() {
	fmt.Print("\nTekan ENTER untuk kembali...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// ClearScreen clears the terminal screen
func ClearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code
}

// PrintTable display data in table format
func PrintTable(headers []string, rows [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, strings.Join(headers, "\t"))
	for _, row := range rows {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}

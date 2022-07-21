package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func buildBallon(lines []string, maxwidth int) string {
	count := len(lines)
	var ret []string

	borders := []string{"/", "\\", "\\", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)

	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[4], lines[0], borders[5])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1

		for ; i < count-1; i++ {
			s = fmt.Sprintf("%s %s %s", borders[3], lines[i], borders[3])
			ret = append(ret, s)
		}

		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[0])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}
	return w
}

func normalizeStringLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func readFromArgs() []string {
	args := os.Args[1]
	splitArgs := strings.Split(args, "\n")

	return splitArgs
}

func readFromStdIn() []string {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("Command must take a string argument or input")
		fmt.Println("Usage: echo 'hello world' | gosay")
		fmt.Println("Usage: gosay 'hello world'")
		os.Exit(1)
	}

	var output []string

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, string(input))
	}
	return output
}

func main() {

	var output []string

	if len(os.Args) > 1 {
		output = readFromArgs()
	} else {
		output = readFromStdIn()
	}

	var cow = `   \  ^__^
    \ (oo)\_______
      (__)\       )\/\
          ||----w |
          ||     ||
  `
	lines := tabsToSpaces(output)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringLength(lines, maxwidth)
	ballon := buildBallon(messages, maxwidth)
	fmt.Println(ballon)
	fmt.Println(cow)
	fmt.Println()
}

package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"github.com/fatih/color"
)

// ScanFile scans a given file path for pattern matches and optionally redacts output
func ScanFile(path string, redact bool) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", path, err)
		return 0
	}
	defer file.Close()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	matchCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		for name, re := range Patterns {
			if re.MatchString(line) {
				matchCount++

				displayLine := line
				if redact {
					displayLine = redactLine(line)
				}

				fmt.Fprintf(writer, "%s\t%s\tline %d\t|\t%s\n",
					red("["+name+"]"),
					cyan(path),
					lineNum,
					yellow(displayLine),
				)
			}
		}
		lineNum++
	}

	writer.Flush()
	return matchCount
}

// redactLine replaces secret values with ***REDACTED***
func redactLine(line string) string {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) == 2 {
		return parts[0] + "=***REDACTED***"
	}
	return "***REDACTED***"
}

package scanner

import (
	"text/tabwriter"
	"os"
	"fmt"
	"github.com/fatih/color"
)


var writer = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

func PrintFinding(rule string, file string, lineNum int, line string) {
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Fprintf(writer,
		"%s\t%s\tline %d\t|\t%s\n",
		red("[" + rule + "]"),
		cyan(file),
		lineNum,
		yellow(line),
	)
}


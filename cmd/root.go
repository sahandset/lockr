package cmd

import (
	"fmt"
	"os"
	"github.com/sahandset/lockr/internal"
	"github.com/spf13/cobra"
)

var (
	rulesetPath string
	redact      bool
	testMode    bool
)

// rootCmd defines the CLI behavior
var rootCmd = &cobra.Command{
	Use:   "lockr [files...]",
	Short: "Scan env/config files for secrets and misconfigurations",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Load custom rules if provided
		if rulesetPath != "" {
			if err := scanner.LoadCustomRules(rulesetPath); err != nil {
				fmt.Println("Error loading rules:", err)
				os.Exit(1)
			}
		}

		totalMatches := 0

		for _, file := range args {
			fmt.Println("Scanning:", file)
			matches := scanner.ScanFile(file, redact)
			totalMatches += matches
		}

		if totalMatches == 0 {
			fmt.Println("No issues found.")
		} else {
			fmt.Printf("%d issues found.\n", totalMatches)
			if testMode {
				os.Exit(1)
			}
		}
	},
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// init defines CLI flags
func init() {
	rootCmd.PersistentFlags().StringVar(&rulesetPath, "ruleset", "", "Path to custom rules JSON or YAML file")
	rootCmd.PersistentFlags().BoolVar(&redact, "redact", false, "Redact secret values in output")
	rootCmd.PersistentFlags().BoolVar(&testMode, "test", false, "Exit with code 1 if issues are found (CI mode)")
}

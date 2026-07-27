package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
	verbose      bool
)

var rootCmd = &cobra.Command{
	Use:   "wcag-scanner",
	Short: "WCAG accessibility scanner for HTML files",
	Long: `wcag-scanner is a fast, concurrent CLI tool that scans HTML files
against WCAG 2.2 standards, identifies accessibility issues, and provides
recommended fixes.

Scan a single file:
  wcag-scanner scan index.html

Scan a directory recursively:
  wcag-scanner scan ./src --output markdown

Output to JSON for CI integration:
  wcag-scanner scan ./src --output json > report.json`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show help if no subcommand is provided
		cmd.Help()
	},
}

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan HTML files for WCAG issues",
	Long:  "Recursively scan a directory or single file for WCAG accessibility violations.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		fmt.Printf("🔍 Scanning: %s\n", path)
		fmt.Printf("📋 Output format: %s\n", outputFormat)

		if verbose {
			fmt.Println("📝 Verbose mode enabled")
		}

		fmt.Println("\n✅ Scan complete (placeholder)")

		return nil
	},
}

func init() {
	// Add scan command to root
	rootCmd.AddCommand(scanCmd)

	// Persistent flags (available to all commands)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Local flags (only for scan command)
	scanCmd.Flags().StringVarP(&outputFormat, "output", "o", "terminal", "Output format: terminal|markdown|json")
	scanCmd.Flags().Bool("fail-on", false, "Exit with error on severity level: error|warning")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

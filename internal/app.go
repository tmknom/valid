package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// AppName is the cli name (set by main.go)
var AppName string

// AppVersion is the current version (set by main.go)
var AppVersion string

type App struct {
	*IO
	rootCmd *cobra.Command
}

func NewApp(io *IO) *App {
	return &App{
		IO: io,
		rootCmd: &cobra.Command{
			Use:           AppName,
			Version:       AppVersion,
			Short:         "Validates that input values meet specified rules",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
	}
}

func (a *App) Run(ctx context.Context, args []string) error {
	a.rootCmd.SetContext(ctx)

	// override default settings
	a.rootCmd.SetArgs(args)
	a.rootCmd.SetIn(a.IO.InReader)
	a.rootCmd.SetOut(a.IO.OutWriter)
	a.rootCmd.SetErr(a.IO.ErrWriter)

	// setup log
	cobra.OnInitialize(func() { a.setupLog(args) })

	// setup version option
	a.rootCmd.SetVersionTemplate(AppVersion)

	// setup flags
	orchestrator := newOrchestrator()
	a.rootCmd.Flags().StringVar(&orchestrator.Value.raw, "value", "", "the value to validate against the specified rules")
	a.rootCmd.Flags().StringVar(&orchestrator.Value.name, "value-name", "", "the name of the value to include in error messages")
	a.rootCmd.Flags().BoolVar(&orchestrator.Value.mask, "mask-value", false, "masks the value in error messages to protect sensitive data")
	a.rootCmd.Flags().StringVar(&orchestrator.Formatter.format, "format", "default", "specifies the output format (default, github-actions)")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.min, "min", "", "validates that the value is greater than or equal to the specified minimum")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.max, "max", "", "validates that the value is less than or equal to the specified maximum")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.exactLength, "exact-length", "", "validates that the length of value is exactly the specified number")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.minLength, "min-length", "", "validates that the length of value is greater than or equal to the specified minimum")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.maxLength, "max-length", "", "validates that the length of value is less than or equal to the specified maximum")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.notEmpty, "not-empty", false, "validates that the value is not empty")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.digit, "digit", false, "validates that the value contains only digits (0-9)")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.alpha, "alpha", false, "validates that the value contains only English letters (a-zA-Z)")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.alphanumeric, "alphanumeric", false, "validates that the value contains only English letters and digits (a-zA-Z0-9)")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.ascii, "ascii", false, "validates that the value contains only ASCII characters")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.printableASCII, "printable-ascii", false, "validates that the value contains only printable ASCII characters")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.lowerCase, "lower-case", false, "validates that the value contains only lowercase Unicode letters")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.upperCase, "upper-case", false, "validates that the value contains only uppercase Unicode letters")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.int, "int", false, "validates that the value is an integer")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.float, "float", false, "validates that the value is a floating-point number")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.url, "url", false, "validates that the value is a valid URL")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.domain, "domain", false, "validates that the value is a valid domain")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.email, "email", false, "validates that the value is a valid email address")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.semver, "semver", false, "validates that the value is a valid semantic version")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.uuid, "uuid", false, "validates that the value is a valid UUID")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.base64, "base64", false, "validates that the value is a valid Base64 string")
	a.rootCmd.Flags().BoolVar(&orchestrator.Validator.json, "json", false, "validates that the value is a valid JSON string")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.pattern, "pattern", "", "validates that the value matches the specified regular expression")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.enum, "enum", "", "validates that the value matches one of the specified enumerations (comma-separated list)")
	a.rootCmd.Flags().StringVar(&orchestrator.Validator.timestamp, "timestamp", "", "validates that the value matches the timestamp format specified in the timestamp input (rfc3339, datetime, date, or time)")

	a.rootCmd.RunE = func(cmd *cobra.Command, args []string) error { return orchestrator.Orchestrate() }
	return a.rootCmd.Execute()
}

func (a *App) setupLog(args []string) {
	log.SetOutput(io.Discard)
	if a.isDebug() {
		log.SetOutput(a.IO.ErrWriter)
	}
	log.SetPrefix(fmt.Sprintf("[%s] ", AppName))
	log.Printf("Start: args: %v", args)
}

func (a *App) isDebug() bool {
	switch os.Getenv("VALID_DEBUG") {
	case "true", "1", "yes":
		return true
	default:
		return false
	}
}

type IO struct {
	InReader  io.Reader
	OutWriter io.Writer
	ErrWriter io.Writer
}

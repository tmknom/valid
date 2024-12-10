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
			Use:          AppName,
			Version:      AppVersion,
			Short:        "Validates that input values meet specified rules",
			SilenceUsage: true,
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
	validator := newValidator()
	a.rootCmd.Flags().StringVar(&validator.value, "value", "", "value for validation")
	a.rootCmd.Flags().StringVar(&validator.min, "min", "", "checks if the value is greater than or equal to the specified minimum")
	a.rootCmd.Flags().StringVar(&validator.max, "max", "", "checks if the value is less than or equal to the specified maximum")
	a.rootCmd.Flags().StringVar(&validator.exactlyLength, "exactly-length", "", "checks if the length matches exactly")
	a.rootCmd.Flags().StringVar(&validator.minLength, "min-length", "", "checks if the length is greater than or equal to the specified minimum")
	a.rootCmd.Flags().StringVar(&validator.maxLength, "max-length", "", "checks if the length is less than or equal to the specified maximum")
	a.rootCmd.Flags().BoolVar(&validator.notEmpty, "not-empty", false, "checks if the value is not empty")
	a.rootCmd.Flags().BoolVar(&validator.digit, "digit", false, "checks if the value contains only digits (0-9)")
	a.rootCmd.Flags().BoolVar(&validator.alpha, "alpha", false, "checks if the value contains only English letters (a-zA-Z)")
	a.rootCmd.Flags().BoolVar(&validator.alphanumeric, "alphanumeric", false, "checks if the value contains only English letters and digits (a-zA-Z0-9)")
	a.rootCmd.Flags().BoolVar(&validator.ascii, "ascii", false, "checks if the value contains only ASCII characters")
	a.rootCmd.Flags().BoolVar(&validator.printableASCII, "printable-ascii", false, "checks if the value contains only printable ASCII characters")
	a.rootCmd.Flags().BoolVar(&validator.int, "int", false, "checks if the value is an integer")
	a.rootCmd.Flags().BoolVar(&validator.float, "float", false, "checks if the value is a floating-point number")
	a.rootCmd.Flags().BoolVar(&validator.url, "url", false, "checks if the value is a valid URL")
	a.rootCmd.Flags().BoolVar(&validator.email, "email", false, "checks if the value is a valid email address")
	a.rootCmd.Flags().BoolVar(&validator.semver, "semver", false, "checks if the value is a valid semantic version")
	a.rootCmd.Flags().BoolVar(&validator.base64, "base64", false, "checks if the value is encoded in Base64")
	a.rootCmd.Flags().BoolVar(&validator.json, "json", false, "checks if the value in valid JSON format")
	a.rootCmd.Flags().StringVar(&validator.pattern, "pattern", "", "checks if the value matches the specified regular expression")
	a.rootCmd.Flags().StringVar(&validator.enum, "enum", "", "checks if the value can be found in the given enumerations")
	a.rootCmd.Flags().StringVar(&validator.timestamp, "timestamp", "", "checks if the value is a timestamp whose format is specified by the layout [rfc3339,datetime,date,time]")

	a.rootCmd.RunE = func(cmd *cobra.Command, args []string) error { return validator.validate() }
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

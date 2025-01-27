package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tmknom/valid/internal"
)

// Specify explicitly in ldflags
// For full details, see Makefile and .goreleaser.yml
var (
	name    = ""
	version = ""
	commit  = ""
	date    = ""
	url     = ""
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	io := &internal.IO{
		InReader:  os.Stdin,
		OutWriter: os.Stdout,
		ErrWriter: os.Stderr,
	}
	internal.AppName = name
	internal.AppVersion = fmt.Sprintf("%s version %s (%s:%s)\n%s\n", name, version, commit, date, url)
	return internal.NewApp(io).Run(ctx, os.Args[1:])
}

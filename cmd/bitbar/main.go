package main

import (
	"os"

	"github.com/muffix/mvg-info/internal/bitbar"
	"github.com/muffix/mvg-info/pkg/interruption"
)

const defaultMaxLineLength = 120

func main() {
	p := bitbar.Printer{MaxLineLength: defaultMaxLineLength}
	p.Interruptions(interruption.NewClient().Interruptions())
	p.Print(os.Stdout)
}

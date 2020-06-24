package bitbar

import (
	"fmt"
	"io"
	"strings"

	"github.com/muffix/mvg-info/pkg/interruption"
)

// Printer is a struct that can print interruptions in the BitBar style
type Printer struct {
	MaxLineLength int
	builder       strings.Builder
	count         int
	err           error
}

// Interruptions is used to set the interruptions that the printer is supposed to print or an error
func (p *Printer) Interruptions(interruptions []interruption.Interruption, err error) {
	if err != nil {
		p.err = err
		return
	}

	for _, i := range interruptions {
		if !i.Ticker {
			continue
		}

		p.addInterruption(i)
	}
}

func (p *Printer) addInterruption(i interruption.Interruption) {
	p.count++

	p.writeln("", "---")
	p.writeln("Updated", i.ModificationDate.String())
	if len(i.Lines) > 0 {
		p.writeln("Affected lines", i.Lines.String())
	}

	p.writeln("", i.Title)
	p.writeln("", i.Text)

	if i.Duration.Text != "" {
		p.writeln("Duration", i.Duration.Text)
	}
}

func (p *Printer) writeln(label, text string) {
	fullText := text
	if label != "" {
		fullText = fmt.Sprintf("%s: %s", label, text)
	}
	fullText = strings.Replace(fullText, "<br />", "\n", -1)

	p.builder.WriteString(fmt.Sprintf("%s\n", p.trimLineLength(fullText)))
}

// Print writes the previously added Interruptions out to a writer
func (p *Printer) Print(w io.Writer) {
	if p.err != nil {
		_, _ = fmt.Fprintf(w, "🚇⚠️\n%s", p.err)
		return
	}

	_, _ = fmt.Fprintf(w, "🚇%d️\n%s", p.count, p.builder.String())
}

func (p *Printer) trimLineLength(text string) string {
	if p.MaxLineLength == 0 {
		return text
	}

	var shortLinesOut []string
	linesIn := strings.Split(text, "\n")

	for _, lineIn := range linesIn {
		var shortLines []string
		words := strings.Fields(lineIn)

		var currentLine string
		for _, word := range words {
			if len(word)+len(currentLine) > p.MaxLineLength {
				shortLines = append(shortLines, strings.TrimSpace(currentLine))
				currentLine = ""
			}
			currentLine = fmt.Sprintf("%s %s", currentLine, word)
		}
		shortLines = append(shortLines, strings.TrimSpace(currentLine))
		shortLinesOut = append(shortLinesOut, strings.Join(shortLines, "\n"))
	}

	return strings.Join(shortLinesOut, "\n")
}

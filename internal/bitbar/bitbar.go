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

	p.writeSeparator()
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

func (p *Printer) writeSeparator() {
	p.builder.WriteString("---\n")
}

func (p *Printer) writeln(label, text string) {
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	fullText := text
	if label != "" {
		fullText = fmt.Sprintf("%s: %s", label, text)
	}
	fullText = strings.Replace(fullText, "<br />", "\n", -1)

	trimmedLines := strings.Split(p.trimLineLength(fullText), "\n")
	p.builder.WriteString(strings.Join(trimmedLines, " | trim=false\n"))
}

// Print writes the previously added Interruptions out to a writer
func (p *Printer) Print(w io.Writer) {
	if p.err != nil {
		_, _ = fmt.Fprintf(w, "🚇⚠️\n---\n%s", p.err)
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
		shortLinesOut = append(shortLinesOut, strings.Join(p.reshapeLineIntoMaxLength(lineIn), "\n"))
	}

	return strings.Join(shortLinesOut, "\n")
}

func (p *Printer) reshapeLineIntoMaxLength(line string) []string {
	var shortLines []string
	words := strings.Fields(line)

	var currentLine string
	for _, word := range words {
		if len(word)+len(currentLine) > p.MaxLineLength {
			shortLines = append(shortLines, strings.TrimSpace(currentLine))
			currentLine = ""
		}
		currentLine = fmt.Sprintf("%s %s", currentLine, word)
	}
	shortLines = append(shortLines, strings.TrimSpace(currentLine))
	return shortLines
}

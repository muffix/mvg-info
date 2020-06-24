package bitbar

import (
	"bytes"
	"testing"
	"time"

	"github.com/muffix/mvg-info/pkg/interruption"
)

const (
	longMockText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. <br /><br />Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
)

var (
	mockDate = interruption.Epoch(time.Date(2020, 01, 01, 10, 00, 00, 0, time.UTC))
)

type mockWriter struct {
	buf bytes.Buffer
}

func (w *mockWriter) Write(p []byte) (n int, err error) {
	return w.buf.Write(p)
}

func (w *mockWriter) String() string {
	return w.buf.String()
}

func TestPrinter_Print(t *testing.T) {
	testCase := struct {
		interruptions []interruption.Interruption
		want          string
	}{
		[]interruption.Interruption{
			{
				Ticker: true,
				Lines: []interruption.Line{
					{Line: "U1"},
					{Line: "42"},
					{Line: "X999"},
				},
				Title:            "Important message",
				Text:             "Simple text",
				ModificationDate: mockDate,
				Duration:         interruption.Duration{From: mockDate, Until: mockDate, Text: "Some time"},
			},
			{
				Ticker:           true,
				Title:            "Short message",
				Text:             "Simple text",
				ModificationDate: mockDate,
			},
			{
				Ticker:           false,
				Title:            "I shouldn't appear",
				Text:             "I shouldn't appear",
				ModificationDate: mockDate,
			},
			{
				Ticker: true,
				Lines: []interruption.Line{
					{Line: "U1"},
					{Line: "42"},
					{Line: "X999"},
				},
				Title:            longMockText,
				Text:             longMockText,
				ModificationDate: mockDate,
				Duration:         interruption.Duration{From: mockDate, Until: mockDate, Text: "Some time"},
			},
		},
		`🚇3️
--- | trim=false
Updated: Wed Jan 1 10:00:00 UTC | trim=false
Affected lines: U1, 42, X999 | trim=false
Important message | trim=false
Simple text | trim=false
Duration: Some time | trim=false
--- | trim=false
Updated: Wed Jan 1 10:00:00 UTC | trim=false
Short message | trim=false
Simple text | trim=false
--- | trim=false
Updated: Wed Jan 1 10:00:00 UTC | trim=false
Affected lines: U1, 42, X999 | trim=false
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna | trim=false
aliqua. | trim=false
 | trim=false
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute | trim=false
irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat | trim=false
cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. | trim=false
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna | trim=false
aliqua. | trim=false
 | trim=false
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute | trim=false
irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat | trim=false
cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. | trim=false
Duration: Some time | trim=false
`,
	}

	w := &mockWriter{}

	underTest := &Printer{MaxLineLength: 120}
	underTest.Interruptions(testCase.interruptions, nil)
	underTest.Print(w)

	got := w.String()

	if got != testCase.want {
		t.Fatalf("Wrong format printed. Got: \n%s\nWant: \n%s", got, testCase.want)
	}
}

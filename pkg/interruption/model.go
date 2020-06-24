package interruption

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// Epoch is a time.Time that can be unmarshalled from a timestamp in milliseconds as returned by the
// API
type Epoch time.Time

// Lines is an array of public transport lines
type Lines []Line

// Interruption is an interruption message as returned by the API
type Interruption struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Lines    Lines    `json:"lines"`
	Duration Duration `json:"duration"`
	Text     string   `json:"text"`
	Files    struct {
		FileDescription interface{} `json:"fileDescription"`
	} `json:"files"`
	Links struct {
		Link interface{} `json:"link"`
	} `json:"links"`
	EventTypes struct {
		EventType interface{} `json:"eventType"`
	} `json:"eventTypes"`
	ModificationDate Epoch `json:"modificationDate"`
	Ticker           bool  `json:"ticker"`
}

// Line is a public transport line
type Line struct {
	ID      int    `json:"id"`
	Product string `json:"product"`
	Line    string `json:"line"`
}

// Duration is the duration of an interruption
type Duration struct {
	From  Epoch  `json:"from"`
	Until Epoch  `json:"until"`
	Text  string `json:"text"`
}

// UnmarshalJSON converts a timestamp in milliseconds to a time.Time
func (e *Epoch) UnmarshalJSON(timestamp []byte) (err error) {
	timestampStr := string(timestamp)
	if timestampStr == "null" {
		return nil
	}

	epoch, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(e) = time.Unix(epoch/1000, 0)
	return
}

func (e *Epoch) String() string {
	return (*time.Time)(e).Format("Mon Jan 2 15:04:05 MST")
}

// UnmarshalJSON removes a layer of nesting that is returned by the API
func (l *Lines) UnmarshalJSON(lines []byte) (err error) {
	data := struct {
		Lines []Line `json:"line"`
	}{}
	err = json.Unmarshal(lines, &data)
	*l = data.Lines
	return
}

func (l *Lines) String() string {
	lineNames := make([]string, 0, len(*l))
	for _, line := range *l {
		lineNames = append(lineNames, line.Line)
	}
	return strings.Join(lineNames, ", ")
}

func (l *Line) String() string {
	return l.Line
}

// Response is the top-level response from the MVG interruptions API
type Response struct {
	Interruptions []Interruption `json:"interruption"`
	AffectedLines Lines          `json:"affectedLines"`
}

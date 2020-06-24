package interruption

import (
	"encoding/json"
	"net/http"
	"time"
)

const url = "https://www.mvg.de/.rest/betriebsaenderungen/api/interruptions"

// Client is a client for the MVG interruptions API
type Client interface {
	// Interruptions fetches and returns all interruptions available
	Interruptions() ([]Interruption, error)
}

type client http.Client

func (c *client) Interruptions() ([]Interruption, error) {
	response, err := c.fetchInterruptions()
	if err != nil {
		return nil, err
	}
	return response.Interruptions, nil

}

func (c *client) fetchInterruptions() (*Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	resp, err := (*http.Client)(c).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response Response
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// NewClient creates a new API client
func NewClient() Client {
	return &client{
		Timeout: time.Second * 10,
	}
}

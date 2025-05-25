package langsmith

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Platzhalter für den LangSmith-Client
// Hier werden später Funktionen zur Abfrage der Runs, Fehler und Kosten implementiert.

type Client struct {
	APIKey    string
	ProjectID string
	BaseURL   string // Standard: https://api.langchain.com
}

func NewClient(apiKey, projectID string) *Client {
	return &Client{APIKey: apiKey, ProjectID: projectID, BaseURL: "https://api.langchain.com"}
}

func (c *Client) GetRuns() (int, error) {
	url := fmt.Sprintf("%s/projects/%s/runs", c.BaseURL, c.ProjectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("LangSmith API Fehler: %s", resp.Status)
	}
	var data struct {
		Total int `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}
	return data.Total, nil
}

func (c *Client) GetFailedRuns() (int, error) {
	url := fmt.Sprintf("%s/projects/%s/runs?status=failed", c.BaseURL, c.ProjectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("LangSmith API Fehler: %s", resp.Status)
	}
	var data struct {
		Total int `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}
	return data.Total, nil
}

func (c *Client) GetTotalCosts() (float64, error) {
	url := fmt.Sprintf("%s/projects/%s/costs", c.BaseURL, c.ProjectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0.0, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0.0, fmt.Errorf("LangSmith API Fehler: %s", resp.Status)
	}
	var data struct {
		Total float64 `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0.0, err
	}
	return data.Total, nil
}

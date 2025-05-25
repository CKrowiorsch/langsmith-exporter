package langsmith

// Platzhalter für den LangSmith-Client
// Hier werden später Funktionen zur Abfrage der Runs, Fehler und Kosten implementiert.

type Client struct {
	APIKey string
	ProjectID string
}

func NewClient(apiKey, projectID string) *Client {
	return &Client{APIKey: apiKey, ProjectID: projectID}
}

func (c *Client) GetRuns() (int, error) {
	// TODO: Implementiere API-Call zu LangSmith, um die Anzahl der Runs zu erhalten
	return 0, nil
}

func (c *Client) GetFailedRuns() (int, error) {
	// TODO: Implementiere API-Call zu LangSmith, um die Anzahl der fehlgeschlagenen Runs zu erhalten
	return 0, nil
}

func (c *Client) GetTotalCosts() (float64, error) {
	// TODO: Implementiere API-Call zu LangSmith, um die Gesamtkosten zu erhalten
	return 0.0, nil
}

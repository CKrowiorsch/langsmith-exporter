package langsmith

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"io"
)

func TestGetRuns(t *testing.T) {
	fakeResponse := `{"total": 42}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, fakeResponse)
	}))
	defer ts.Close()

	client := &Client{APIKey: "test", ProjectID: "test", BaseURL: ts.URL}

	runs, err := client.GetRuns()
	if err != nil {
		t.Fatalf("Fehler beim Abrufen der Runs: %v", err)
	}
	if runs != 42 {
		t.Errorf("Erwartet: 42, erhalten: %d", runs)
	}
}

func TestGetFailedRuns(t *testing.T) {
	fakeResponse := `{"total": 5}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fakeResponse)
	}))
	defer ts.Close()

	client := &Client{APIKey: "test", ProjectID: "test", BaseURL: ts.URL}

	failed, err := client.GetFailedRuns()
	if err != nil {
		t.Fatalf("Fehler beim Abrufen der fehlgeschlagenen Runs: %v", err)
	}
	if failed != 5 {
		t.Errorf("Erwartet: 5, erhalten: %d", failed)
	}
}

func TestGetTotalCosts(t *testing.T) {
	fakeResponse := `{"total": 123.45}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fakeResponse)
	}))
	defer ts.Close()

	client := &Client{APIKey: "test", ProjectID: "test", BaseURL: ts.URL}

	costs, err := client.GetTotalCosts()
	if err != nil {
		t.Fatalf("Fehler beim Abrufen der Kosten: %v", err)
	}
	if costs != 123.45 {
		t.Errorf("Erwartet: 123.45, erhalten: %f", costs)
	}
}

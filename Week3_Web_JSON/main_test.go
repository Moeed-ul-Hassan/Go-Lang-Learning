package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	homeHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	body := w.Body.String()
	expected := "Week 3 – Hello World\n"
	if body != expected {
		t.Fatalf("unexpected body: %s", body)
	}
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello/John", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	expected := "Hello, John!"
	if w.Body.String() != expected {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestSearchHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/search?q=go", nil)
	w := httptest.NewRecorder()
	searchHandler(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp["query"] != "go" {
		t.Fatalf("expected query 'go', got %s", resp["query"])
	}
}

func TestEchoHandler(t *testing.T) {
	payload := []byte(`{"text":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	echoHandler(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var msg Message
	if err := json.NewDecoder(w.Body).Decode(&msg); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if msg.Text != "test" {
		t.Fatalf("expected text 'test', got %s", msg.Text)
	}
}

func TestSubmitHandler(t *testing.T) {
	data := "name=Alice&email=alice%40example.com"
	req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	submitHandler(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp["name"] != "Alice" || resp["email"] != "alice@example.com" {
		t.Fatalf("unexpected form data: %v", resp)
	}
}

func TestDataHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/data", nil)
	w := httptest.NewRecorder()
	dataHandler(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var courses []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	if err := json.NewDecoder(w.Body).Decode(&courses); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(courses) != 2 {
		t.Fatalf("expected 2 courses, got %d", len(courses))
	}
}

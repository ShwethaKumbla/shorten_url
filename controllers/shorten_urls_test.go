package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//testcase for Post json data successfully
func TestCreate(t *testing.T) {
	jsonInput := []byte(`{"long_url": "https://tour.golang.org/welcome/1"}`)

	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonInput))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	//create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestUpdate(t *testing.T) {
	jsonInput := []byte(`{"long_url": "https://tour.golang.org/welcome/1"}`)

	req, err := http.NewRequest("PUT", "/update", bytes.NewBuffer(jsonInput))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	//create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Update)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateWhichNotExist(t *testing.T) {
	jsonInput := []byte(`{"long_url": "https://abcdefg/welcome/hello"}`)

	req, err := http.NewRequest("PUT", "/update", bytes.NewBuffer(jsonInput))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	//create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Update)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestRetrieve(t *testing.T) {
	req, err := http.NewRequest("GET", "/list/", nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	//create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Retrieve)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateExisting(t *testing.T) {
	jsonInput := []byte(`{"long_url": "https://tour.golang.org/welcome/1"}`)

	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonInput))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	//create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestRedirect(t *testing.T) {

	req, err := http.NewRequest("GET", "/JkCAP", nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Redirect)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete?endpoint=https://tour.golang.org/welcome/1", nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Delete)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

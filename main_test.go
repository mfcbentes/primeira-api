package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetPeople(t *testing.T) {
	req, err := http.NewRequest("GET", "/people", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPeople)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler retornou código de status errado: obtido %v desejado %v", status, http.StatusOK)
	}
}

func TestGetPerson(t *testing.T) {
	req, err := http.NewRequest("GET", "/people/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/people/{id}", getPerson)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler retornou código de status errado: obtido %v desejado %v", status, http.StatusOK)
	}
}

func TestCreatePerson(t *testing.T) {
	person := &Person{ID: 3, Firstname: "Fernando", Lastname: "Bentes", Address: Address{City: "Santarem", State: "PA"}}
	jsonPerson, _ := json.Marshal(person)

	req, err := http.NewRequest("POST", "/people", bytes.NewBuffer(jsonPerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createPerson)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler retornou código de status errado: obtido %v desejado %v", status, http.StatusCreated)
	}
}

func TestDeletePerson(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/people/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/people/{id}", deletePerson)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler retornou código de status errado: obtido %v desejado %v", status, http.StatusOK)
	}
}

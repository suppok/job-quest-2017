package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("suppok", "1234", "todo_list_api")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/todos", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentTodo(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/todo/45", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Todo not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Todo not found'. Got '%s'", m["error"])
	}
}

func TestCreateTodo(t *testing.T) {
	clearTable()
	payload := []byte(`{"title":"test todo","description":"This is test description"}`)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["title"] != "test todo" {
		t.Errorf("Expected todo title to be 'test todo'. Got '%v'", m["title"])
	}
	if m["description"] != "This is test description" {
		t.Errorf("Expected todo description to be 'This is test description'. Got '%v'", m["description"])
	}
	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected todo ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetTodo(t *testing.T) {
	clearTable()
	addTodos(1)
	req, _ := http.NewRequest("GET", "/todo/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateTodo(t *testing.T) {
	clearTable()
	addTodos(1)
	req, _ := http.NewRequest("GET", "/todo/1", nil)
	response := executeRequest(req)
	var originalTodo map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalTodo)
	payload := []byte(`{"title":"test todo - updated title","description":"updated description"}`)
	req, _ = http.NewRequest("PUT", "/todo/1", bytes.NewBuffer(payload))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["id"] != originalTodo["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalTodo["id"], m["id"])
	}
	if m["title"] == originalTodo["title"] {
		t.Errorf("Expected the title to change from '%v' to '%v'. Got '%v'", originalTodo["title"], m["title"], m["title"])
	}
	if m["description"] == originalTodo["description"] {
		t.Errorf("Expected the description to change from '%v' to '%v'. Got '%v'", originalTodo["description"], m["description"], m["description"])
	}
}

func TestDeleteTodo(t *testing.T) {
	clearTable()
	addTodos(1)
	req, _ := http.NewRequest("GET", "/todo/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	req, _ = http.NewRequest("DELETE", "/todo/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	req, _ = http.NewRequest("GET", "/todo/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func addTodos(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO todos(title, description) VALUES('%s', '%s')", "Todo "+strconv.Itoa(i+1), "Description "+strconv.Itoa(i+1))
		a.DB.Exec(statement)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	a.DB.Exec("DELETE FROM todos")
	a.DB.Exec("ALTER TABLE todos AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS todos
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100)
)`

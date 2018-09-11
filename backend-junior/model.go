package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *todo) getTodo(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT title, description FROM todos WHERE id=%d", t.ID)
	return db.QueryRow(statement).Scan(&t.Title, &t.Description)
}

func (t *todo) updateTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (t *todo) deleteTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (t *todo) createTodo(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO todos(title, description) VALUES('%s', '%s')", t.Title, t.Description)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&t.ID)
	if err != nil {
		return err
	}
	return nil
}

func getTodos(db *sql.DB, start, count int) ([]todo, error) {
	statement := fmt.Sprintf("SELECT id, title, description FROM todos LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []todo{}
	for rows.Next() {
		var u todo
		if err := rows.Scan(&u.ID, &u.Title, &u.Description); err != nil {
			return nil, err
		}
		todos = append(todos, u)
	}
	return todos, nil
}

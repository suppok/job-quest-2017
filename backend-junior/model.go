package main

import (
	"database/sql"
	"errors"
)

type todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	description string `json:"description"`
}

func (t *todo) getTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (t *todo) updateTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (t *todo) deleteTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (t *todo) createTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getTodos(db *sql.DB, start, count int) ([]todo, error) {
	return nil, errors.New("Not implemented")
}

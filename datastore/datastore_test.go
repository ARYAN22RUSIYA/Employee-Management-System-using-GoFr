package leavedatastore

import (
	"database/sql"
	"errors"
	"testing"

	"gofr.dev/pkg/gofr"

	"gofr.dev/model"
)

// MockDB is a mock implementation of the database interface for testing.
type MockDB struct{}

func (db *MockDB) QueryRowContext(ctx *gofr.Context, query string, args ...interface{}) *sql.Row {

	return nil
}

func (db *MockDB) ExecContext(ctx *gofr.Context, query string, args ...interface{}) (sql.Result, error) {

	return nil, nil
}

func TestGetByID(t *testing.T) {
	// Create a new instance of the student data store with a mock database.
	store := &student{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a student ID for testing.
	id := "1"

	// Call the GetByID method.
	_, err := store.GetByID(ctx, id)

	// Check if the error is the expected "EntityNotFound" error.
	if !errors.Is(err, errors.New("EntityNotFound")) {
		t.Errorf("Expected EntityNotFound error, got %v", err)
	}
}

func TestCreate(t *testing.T) {
	// Create a new instance of the student data store with a mock database.
	store := &student{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a student record for testing.
	newStudent := &model.Student{ID: 123, Name: "Monalisa", Age: "22", Class: "12"}

	// Call the Create method.
	_, err := store.Create(ctx, newStudent)

	// Check if there is no error.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	// Create a new instance of the student data store with a mock database.
	store := &student{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a student record for testing.
	updatedStudent := &model.Student{ID: 1, Name: "Lisa", Age: "20", Class: "Biology"}

	// Call the Update method.
	_, err := store.Update(ctx, updatedLeave)

	// Check if there is no error.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDelete(t *testing.T) {
	// Create a new instance of the student data store with a mock database.
	store := &student{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a student ID for testing.
	studentID := 1

	// Call the Delete method.
	err := store.Delete(ctx, studentID)

	// Check if there is no error.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
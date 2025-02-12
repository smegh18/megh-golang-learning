package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGenerateInsertSQL(t *testing.T) {
	doc := map[string]interface{}{
		"id":    1,
		"name":  "John",
		"email": "john@example.com",
	}
	expectedSQL := "INSERT INTO users (id, name, email) VALUES ('1', 'John', 'john@example.com');"

	generatedSQL := generateInsertSQL("users", doc)
	assert.Equal(t, expectedSQL, generatedSQL)
}

func TestGenerateUpdateSQL(t *testing.T) {
	filter := map[string]interface{}{"id": 1}
	updates := map[string]interface{}{"name": "Jane", "email": "jane@example.com"}
	expectedSQL := "UPDATE users SET name='Jane', email='jane@example.com' WHERE id='1';"

	generatedSQL := generateUpdateSQL("users", filter, updates)
	assert.Equal(t, expectedSQL, generatedSQL)
}

func TestGenerateDeleteSQL(t *testing.T) {
	filter := map[string]interface{}{"id": 1}
	expectedSQL := "DELETE FROM users WHERE id='1';"

	generatedSQL := generateDeleteSQL("users", filter)
	assert.Equal(t, expectedSQL, generatedSQL)
}

func TestProcessOplogEntry_Insert(t *testing.T) {
	db, _ := gorm.Open(postgres.New(postgres.Config{DSN: "host=localhost user=postgres password=secret dbname=test port=5432 sslmode=disable"}), &gorm.Config{})
	op := &OplogProcessor{DB: db}
	doc := map[string]interface{}{"id": 1, "name": "Alice"}
	op.ProcessOplogEntry(OplogEntry{Operation: "i", Namespace: "test.users", Document: doc})
}

func TestProcessOplogEntry_Update(t *testing.T) {
	db, _ := gorm.Open(postgres.New(postgres.Config{DSN: "host=localhost user=postgres password=secret dbname=test port=5432 sslmode=disable"}), &gorm.Config{})
	op := &OplogProcessor{DB: db}
	filter := map[string]interface{}{"id": 1}
	updates := map[string]interface{}{"name": "Bob"}
	op.ProcessOplogEntry(OplogEntry{Operation: "u", Namespace: "test.users", UpdateFields: filter, Document: updates})
}

func TestProcessOplogEntry_Delete(t *testing.T) {
	db, _ := gorm.Open(postgres.New(postgres.Config{DSN: "host=localhost user=postgres password=secret dbname=test port=5432 sslmode=disable"}), &gorm.Config{})
	op := &OplogProcessor{DB: db}
	filter := map[string]interface{}{"id": 1}
	op.ProcessOplogEntry(OplogEntry{Operation: "d", Namespace: "test.users", UpdateFields: filter})
}

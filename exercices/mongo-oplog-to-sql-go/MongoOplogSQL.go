package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OplogEntry struct {
	Timestamp    int64  `bson:"ts"`
	Operation    string `bson:"op"`
	Namespace    string `bson:"ns"`
	Document     bson.M `bson:"o"`
	UpdateFields bson.M `bson:"o2"`
}

type OplogProcessor struct {
	DB            *gorm.DB
	LastProcessed int64
	Mutex         sync.Mutex
}

func NewOplogProcessor(dsn string) (*OplogProcessor, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &OplogProcessor{DB: db, LastProcessed: 0}, nil
}

func (op *OplogProcessor) ProcessOplogEntry(entry OplogEntry) {
	table := parseNamespace(entry.Namespace)

	switch entry.Operation {
	case "i":
		sql := generateInsertSQL(table, entry.Document)
		executeSQL(op.DB, sql)
	case "u":
		sql := generateUpdateSQL(table, entry.UpdateFields, entry.Document)
		executeSQL(op.DB, sql)
	case "d":
		sql := generateDeleteSQL(table, entry.UpdateFields)
		executeSQL(op.DB, sql)
	}

	op.Mutex.Lock()
	op.LastProcessed = entry.Timestamp
	op.Mutex.Unlock()
}

func parseNamespace(ns string) string {
	parts := strings.Split(ns, ".")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func generateInsertSQL(table string, doc bson.M) string {
	columns := []string{}
	values := []string{}

	for key, value := range doc {
		columns = append(columns, key)
		values = append(values, fmt.Sprintf("'%v'", value))
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, strings.Join(columns, ", "), strings.Join(values, ", "))
}

func generateUpdateSQL(table string, filter bson.M, updates bson.M) string {
	setClauses := []string{}
	whereClauses := []string{}

	for key, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s='%v'", key, value))
	}
	for key, value := range filter {
		whereClauses = append(whereClauses, fmt.Sprintf("%s='%v'", key, value))
	}

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s;", table, strings.Join(setClauses, ", "), strings.Join(whereClauses, " AND "))
}

func generateDeleteSQL(table string, filter bson.M) string {
	whereClauses := []string{}
	for key, value := range filter {
		whereClauses = append(whereClauses, fmt.Sprintf("%s='%v'", key, value))
	}

	return fmt.Sprintf("DELETE FROM %s WHERE %s;", table, strings.Join(whereClauses, " AND "))
}

func executeSQL(db *gorm.DB, sql string) {
	if err := db.Exec(sql).Error; err != nil {
		log.Printf("Error executing SQL: %s, %v", sql, err)
	}
}

func main() {
	dsn := "host=localhost user=postgres password=secret dbname=test port=5432 sslmode=disable"
	op, err := NewOplogProcessor(dsn)
	if err != nil {
		log.Fatal(err)
	}

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("local")
	collection := db.Collection("oplog.rs")

	cursor, err := collection.Find(nil, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(nil)
	for cursor.Next(nil) {
		var entry OplogEntry
		if err := cursor.Decode(&entry); err != nil {
			log.Println("Error decoding oplog entry:", err)
			continue
		}

		op.ProcessOplogEntry(entry)
	}
}

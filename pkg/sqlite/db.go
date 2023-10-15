package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

type Operator interface {
	DBFileExists(dbPath string) error
	CreateDB(filepath string) (*sql.DB, error)
}

type DB struct {
	ctx  context.Context
	Name string
}

func NewDBOperator(ctx context.Context, dbName string) Operator {
	return &DB{
		ctx:  ctx,
		Name: dbName,
	}
}

func (d *DB) DBFileExists(dbPath string) error {
	if dbPath == "" {
		return fmt.Errorf("db path is empty")
	}

	if err := utils.FileExistAndItIsAFile(dbPath); err != nil {
		return fmt.Errorf("db file does not exist: %w", err)
	}

	return nil
}

func (d *DB) CreateDB(filepath string) (*sql.DB, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		db, err := sql.Open("sqlite3", filepath)
		if err != nil {
			return nil, fmt.Errorf("failed to create the handle: %w", err)
		}
		// Check if connection with database is established
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("failed to keep connection alive: %w", err)
		}
		return db, nil
	}

	return nil, fmt.Errorf("database with the filepath %s already exists", filepath)
}

package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	// golangci-lint: ignore
	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"
	_ "github.com/mattn/go-sqlite3"
)

type SQLRunner interface {
	RunSQL(statement string) error
}

type DBClient struct {
	ctx    context.Context
	logger o11y.LoggerInterface
	c      *sql.DB
}

func newClient(dbFilePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	defer db.Close()

	return db, nil
}

func NewDBClient(ctx context.Context, l o11y.LoggerInterface, dbFilePath string,
	existingSQLClient *sql.DB) (SQLRunner, error) {
	if existingSQLClient != nil {
		return &DBClient{
			ctx:    ctx,
			logger: l,
			c:      existingSQLClient,
		}, nil
	}

	client, err := newClient(dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %w", err)
	}

	return &DBClient{
		ctx:    ctx,
		logger: l,
		c:      client,
	}, nil
}

func (d *DBClient) RunSQL(statement string) error {
	if statement == "" {
		return fmt.Errorf("statement is empty")
	}

	_, err := d.c.ExecContext(d.ctx, statement)
	if err != nil {
		return fmt.Errorf("failed to run sql statement: %w", err)
	}

	return nil
}

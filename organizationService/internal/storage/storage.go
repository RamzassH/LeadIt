package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var AllowedTables = map[string]struct{}{
	"employees":     {},
	"organizations": {},
	"roles":         {},
	"projects":      {},
}
var (
	ErrAlreadyExists = errors.New("entity already exists")
	ErrNotFound      = errors.New("entity not found")
	ErrInvalidTable  = errors.New("invalid table")
)

type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func GetById(ctx context.Context, db sqlx.ExtContext, table string, id int64, dest interface{}) error {
	if _, ok := AllowedTables[table]; !ok {
		return ErrInvalidTable
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table)
	return sqlx.GetContext(ctx, db, dest, query, id)
}

func Delete(ctx context.Context, db Executor, table string, id int64) (rowsAffected int64, err error) {

	if _, ok := AllowedTables[table]; !ok {
		return 0, ErrInvalidTable
	}

	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, table)

	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("delete failed: %w", err)
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("%w: table=%s id=%d", ErrNotFound, table, id)
	}

	return rowsAffected, nil
}

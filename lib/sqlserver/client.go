package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
)

// Client is a wrapper for sql.DB
// jetti:bean Client
type Client struct {
	baseClient *sql.DB
	serverHost string
	port       int
	user       string
	password   string
	database   string
}

func New(ctx context.Context, serverHost string, port int, user string, password string, database string) (*Client, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		serverHost, user, password, port, database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	context.AfterFunc(ctx, func() {
		db.Close()
	})
	return &Client{db, serverHost, port, user, password, database}, nil
}

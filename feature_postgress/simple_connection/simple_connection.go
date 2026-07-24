package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	val := os.Getenv("connection")
	return pgx.Connect(ctx, val)

}

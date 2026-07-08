package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(conn *pgx.Conn, ctx context.Context,
	task TaskModel,
) error {
	sqlQuery := `
INSERT INTO tasks (title,
 description , 
 completed ,
 created_at  )
VALUES($1 , $2 ,$3 , $4);

`
	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.Created_at, task.Completed_at)

	return err

}

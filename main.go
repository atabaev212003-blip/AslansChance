package main

import (
	"Study/feature_postgress/simple_connection"
	"Study/feature_postgress/simple_sql"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	//if err := simple_sql.InsertRow(conn, ctx, "Lizard", "погулять", false, time.Now()); err != nil {
	//	panic(err)
	//}

	//	if err := simple_sql.DeleteRow(ctx, conn); err != nil {
	//		panic(err)
	//	}
	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		if task.ID == 5 {
			task.Title = "Lalalala"
			task.Description = "Sing a song"
			task.Completed = true
			now := time.Now()
			task.Completed_at = &now
			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}

	simple_sql.GoodText(tasks)
	fmt.Println("Succeed!")
}

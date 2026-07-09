package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// 1 . we create rows.Next
// 1.2 we create variables to store the data information
// 2. we create rows.Scan
// 3. we close the rows

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
SELECT  id,title , description ,completed , created_at , completed_at 
FROM tasks 
ORDER BY id ASC

`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GoodText(task []TaskModel) {
	for i := 0; i < len(task); i++ {

		fmt.Println("Id :", task[i].ID)

		fmt.Println("Title :", task[i].Title)

		fmt.Println("Description :", task[i].Description)

		fmt.Println("Completed :", task[i].Completed)

		fmt.Println("Created at :", task[i].Created_at)

		fmt.Println("Completed at :", task[i].Completed_at)
		fmt.Println("-----------------------")
	}

}

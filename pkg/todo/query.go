package todo

import (
	"database/sql"
	"log"
)

func (c *core) GetTodoListFromDB() ([]TodoDB, error) {
	var todos []TodoDB

	query := `SELECT * FROM todos WHERE deleted_at IS NULL`

	err := c.db.Select(&todos, query)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[DB] Error query get todos:", err)
		return todos, err
	}

	return todos, nil
}

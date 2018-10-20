package todo

import (
	"database/sql"
	"log"
)

func (c *core) GetTodoListFromDB() ([]TodoDB, error) {
	var todos []TodoDB

	query := `SELECT 
		id,
		COALESCE(title, '') as title,
		COALESCE(completed, 1) as completed,
		COALESCE(created_at, now()) as created_at 
		FROM todos WHERE deleted_at IS NULL`

	err := c.db.Select(&todos, query)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[DB] Error query get todos:", err)
		return todos, err
	}

	return todos, nil
}

func (c *core) CreateTodoFromDB(todoData *TodoData) ([]TodoDB, error) {
	stmt, err := c.db.PrepareNamed(`INSERT INTO todos (title) VALUES (:title)`)
	if err != nil {
		log.Println("[DB] Error prepared name query create todo:", err)
		return nil, err
	}

	_, err = stmt.Exec(&todoData)
	if err != nil {
		log.Println("[DB] Error query insert todo:", err)
		return nil, err
	}

	return nil, nil
}

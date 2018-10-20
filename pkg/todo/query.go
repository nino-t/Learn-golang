package todo

import (
	"database/sql"
	"log"
)

func (c *core) GetTodoListFromDB() ([]TodoDB, error) {
	var todos []TodoDB

	query := `
		SELECT 
				id,
				COALESCE(title, '') as title,
				COALESCE(completed, 1) as completed,
				COALESCE(created_at, now()) as created_at 
			FROM
				todos
			WHERE
				deleted_at IS NULL`

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

func (c *core) GetTodoDetailFromDB(primaryId interface{}) ([]TodoDB, error) {
	var todos []TodoDB

	query :=
		`SELECT 
				id,
				COALESCE(title, '') as title,
				COALESCE(completed, 1) as completed,
				COALESCE(created_at, now()) as created_at 
			FROM 
				todos 
			WHERE
				deleted_at IS NULL AND
				id = ?`

	err := c.db.Select(&todos, query, primaryId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[DB] Error query get todo:", err)
		return todos, err
	}

	return todos, nil
}

func (c *core) UpdateTodoFromDB(todoData *TodoData) ([]TodoDB, error) {
	stmt, err := c.db.PrepareNamed(`
		UPDATE 
			todos
		SET
			title=:title
		WHERE
			id = :id
	`)
	if err != nil {
		log.Println("[DB] Error prepared name query update todo:", err)
		return nil, err
	}

	_, err = stmt.Exec(todoData)
	if err != nil {
		log.Println("[DB] Error query update todo:", err)
		return nil, err
	}

	return nil, nil
}

func (c *core) DeleteTodoFromDB(primaryId interface{}) ([]TodoDB, error) {
	var todos []TodoDB

	query :=
		`DELETE
			FROM
				todos 
			WHERE
				id = ?`

	err := c.db.Select(&todos, query, primaryId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[DB] Error query delete todo:", err)
		return todos, err
	}

	return todos, nil
}

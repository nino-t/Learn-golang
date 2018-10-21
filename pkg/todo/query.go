package todo

import (
	"database/sql"
	"log"
)

func (c *core) GetTodoListFromDB() ([]TodoModel, error) {
	var todos []TodoModel

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

func (c *core) CreateTodoFromDB(todoModel *TodoModel) ([]TodoModel, error) {
	stmt, err := c.db.PrepareNamed(`INSERT INTO todos (title) VALUES (:title)`)
	if err != nil {
		log.Println("[DB] Error prepared name query create todo:", err)
		return nil, err
	}

	_, err = stmt.Exec(&todoModel)
	if err != nil {
		log.Println("[DB] Error query insert todo:", err)
		return nil, err
	}

	return nil, nil
}

func (c *core) GetTodoDetailFromDB(todoModel *TodoModel) ([]TodoModel, error) {
	var todos []TodoModel

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

	err := c.db.Select(&todos, query, todoModel.ID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[DB] Error query get todo:", err)
		return todos, err
	}

	return todos, nil
}

func (c *core) UpdateTodoFromDB(todoModel *TodoModel) ([]TodoModel, error) {
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

	_, err = stmt.Exec(todoModel)
	if err != nil {
		log.Println("[DB] Error query update todo:", err)
		return nil, err
	}

	return nil, nil
}

func (c *core) DeleteTodoFromDB(todoModel *TodoModel) ([]TodoModel, error) {
	stmt, err := c.db.PrepareNamed(`
				DELETE
					FROM 
						todos
					WHERE
						id = :id
			`)
	if err != nil {
		log.Println("[DB] Error prepared name query update todo:", err)
		return nil, err
	}

	_, err = stmt.Exec(todoModel)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[DB] Error query delete todo:", err)
		return nil, err
	}

	return nil, nil
}

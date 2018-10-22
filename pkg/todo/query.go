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

func (c *core) CreateTodoFromDB(todoModel *TodoModel) (interface{}, error) {
	stmt, err := c.db.PrepareNamed(`INSERT INTO todos (title) VALUES (:title)`)
	if err != nil {
		log.Println("[DB] Error prepared name query create todo:", err)
		return nil, err
	}

	result, err := stmt.Exec(&todoModel)
	if err != nil {
		log.Println("[DB] Error query insert todo:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("[DB] Error query failed to get last insertID:", err)
		return nil, err
	}

	todo := TodoModel{}
	todo.ID = int(id)

	res, err := c.GetTodoDetailFromDB(&todo)
	return res, err
}

func (c *core) GetTodoDetailFromDB(todoModel *TodoModel) (interface{}, error) {
	var todo TodoModel

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

	row := c.db.QueryRow(query, todoModel.ID)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		log.Println("[DB] Error query failed to Fetch row:", err)
		return nil, err
	}

	return todo, nil
}

func (c *core) UpdateTodoFromDB(todoModel *TodoModel) (interface{}, error) {
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

	res, err := c.GetTodoDetailFromDB(todoModel)
	return res, err
}

func (c *core) DeleteTodoFromDB(todoModel *TodoModel) (interface{}, error) {
	stmt, err := c.db.PrepareNamed(`
				DELETE
					FROM 
						todos
					WHERE
						id = :id
			`)
	if err != nil {
		log.Println("[DB] Error prepared name query delete todo:", err)
		return nil, err
	}

	_, err = stmt.Exec(todoModel)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[DB] Error query delete todo:", err)
		return nil, err
	}

	return nil, err
}

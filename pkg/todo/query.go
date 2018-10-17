package todo

import "log"

func (c *core) GetTodoListFromDB(todo TodoDB) error {
	stmt, err := c.db.PrepareNamed(`SELECT * FROM todos WHERE deleted_at IS NULL`)
	if err != nil {
		log.Println("[DB] Error prepared name query get todo:", err)
		return err
	}

	_, err = stmt.Exec(todo)
	if err != nil {
		log.Println("[DB] Error query get todo:", err)
		return err
	}

	return nil
}

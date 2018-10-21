package todo

type (
	TodoModel struct {
		ID        int    `db:"id" json:"id"`
		Title     string `db:"title" json:"title"`
		Completed int8   `db:"completed" json:"completed"`
		CreatedAt string `db:"created_at" json:"created_at"`
	}
)

package todo

type (
	TodoAttributes struct {
		Title     string `json:"title" msg:"title"`
		Completed int8   `json:"completed" msg:"completed"`
	}

	TodoDB struct {
		ID        int    `db:"id"`
		Title     string `db:"title"`
		Completed int8   `db:"completed"`
	}
)

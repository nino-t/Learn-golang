package user

type (
	TodoModel struct {
		ID        int    `db:"id" json:"id"`
		Name      string `db:"name" json:"name"`
		Email     string `db:"email" json:"email"`
		Password  string `db:"password" json:"password"`
		CreatedAt string `db:"created_at" json:"created_at"`
	}
)

package database

type Clicker struct {
	UserID     string `db:"user_id"`
	ClickCount uint64 `db:"click_count"`
}

type User struct {
	ID    string `db:"user_id"`
	Name  string `db:"user_name"`
	Email string `db:"user_email"`
}

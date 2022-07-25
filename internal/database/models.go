package database

type Clicker struct {
	UserID     string `db:"user_id"`
	ClickCount uint64 `db:"click_count"`
}

type User struct {
	UserID uint `db:"user_id"`
}

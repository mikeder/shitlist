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

func (u *User) Authentications(us UserStore) ([]*Authentication, error) {
	return us.GetUserAuthentications(u.ID)
}

type UserAuthentications struct {
	User            User
	Authentications []Authentication
}

type Authentication struct {
	ID       string                 `db:"authentication_id"`
	UserID   string                 `db:"authentication_user_id"`
	Provider AuthenticationProvider `db:"authentication_provider"`
}

type AuthenticationProvider string

const (
	AuthenticationTypeGithub = AuthenticationProvider("github")
	AuthenticationTypeGoogle = AuthenticationProvider("google")
)

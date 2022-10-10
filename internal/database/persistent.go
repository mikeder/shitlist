package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// enforce PersistentDataStore inplements required interfaces
var _ ClickStore = (*PersistentDataStore)(nil)
var _ UserStore = (*PersistentDataStore)(nil)

type PersistentDataStore struct {
	db *sql.DB
}

type PersistentDataStoreConfig struct {
	Host       string `required:"true"`
	Port       int    `default:"5432"`
	User       string `required:"true"`
	Password   string `required:"true"`
	SchemaName string `required:"true" split_words:"true"`
}

func NewPersistentClickStore(cfg PersistentDataStoreConfig) (*PersistentDataStore, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.SchemaName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	// check db
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PersistentDataStore{
		db: db,
	}, nil
}

func (p *PersistentDataStore) AddClick(userID string) (*Clicker, error) {
	var c *Clicker = &Clicker{
		UserID: userID,
	}

	q := `
UPDATE clicks 
SET click_count=click_count+1 
WHERE user_id=$1 
RETURNING click_count`

	err := p.db.QueryRow(q, userID).Scan(&c.ClickCount)
	if err != nil {
		return nil, fmt.Errorf("add clicker %v: %w", userID, err)
	}

	return c, nil
}

func (p *PersistentDataStore) GetClickers() ([]*Clicker, error) {
	var cs []*Clicker = []*Clicker{}

	q := `
SELECT user_id, click_count 
FROM clicks`

	rows, err := p.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("get clickers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		c := &Clicker{}
		if err := rows.Scan(&c.UserID, &c.ClickCount); err != nil {
			return nil, fmt.Errorf("scan clicker: %w", err)
		}
		cs = append(cs, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate clicker rows: %w", err)
	}

	return cs, nil
}

func (p *PersistentDataStore) AddUser(name, email string) (*User, error) {
	var u *User = &User{
		Name:  name,
		Email: email,
	}

	q := `
INSERT INTO users(user_name, user_email) VALUES($1, $2) returning user_id`

	err := p.db.QueryRow(q, name, email).Scan(&u.ID)
	if err != nil {
		return nil, fmt.Errorf("add user %v - %v: %w", name, email, err)
	}
	return p.addClicker(u)
}

func (p *PersistentDataStore) addClicker(u *User) (*User, error) {
	q := `
INSERT INTO clicks(user_id, click_count) VALUES($1, $2)`

	_, err := p.db.Exec(q, u.ID, 0)
	if err != nil {
		return nil, fmt.Errorf("add clicker %v: %w", u.ID, err)
	}
	return u, nil
}

func (p *PersistentDataStore) GetUserByEmail(email string) (*User, error) {
	var u *User = &User{}

	q := `
SELECT user_id, user_email, user_name 
FROM users 
WHERE user_email=$1`

	err := p.db.QueryRow(q, email).Scan(&u.ID, &u.Email, &u.Name)
	if err != nil {
		return nil, fmt.Errorf("get user by email %v: %w", email, err)
	}
	return u, nil
}

func (p *PersistentDataStore) GetUserByName(name string) (*User, error) {
	var u *User = &User{}

	q := `
SELECT user_id, user_email, user_name 
FROM users 
WHERE user_name=$1`

	err := p.db.QueryRow(q, name).Scan(&u.ID, &u.Email, &u.Name)
	if err != nil {
		return nil, fmt.Errorf("get user by name %v: %w", name, err)
	}
	return u, nil
}

func (p *PersistentDataStore) GetUserAuthentications(userID string) (*UserAuthentications, error) {
	var ua *UserAuthentications = &UserAuthentications{}

	q := `
SELECT user_id, user_email, user_name 
FROM users 
INNER JOIN authentications 
ON user_id=fk_authentication_user
WHERE user_id=$1`

	rows, err := p.db.Query(q, userID)
	if err != nil {
		return ua, fmt.Errorf("get %v authentications: %w", userID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a Authentication
		rows.Scan(
			&ua.User.ID,
			&ua.User.Email,
			&ua.User.Name,
			&a.ID,
			&a.Provider,
		)
		ua.Authentications = append(ua.Authentications, a)
	}
	return ua, nil
}

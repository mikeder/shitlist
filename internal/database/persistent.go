package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

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

	q := `update clicks set click_count=click_count+1 where user_id=$1 returning click_count`
	err := p.db.QueryRow(q, userID).Scan(&c.ClickCount)
	if err != nil {
		log.Println("add click: " + err.Error())
		return nil, errors.New("unable to add click")
	}

	return c, nil
}

func (p *PersistentDataStore) GetClickers() ([]*Clicker, error) {
	var cs []*Clicker = []*Clicker{}

	q := `select user_id, click_count from clicks`
	rows, err := p.db.Query(q)
	if err != nil {
		log.Println("get clickers: " + err.Error())
		return nil, errors.New("failed to get clickers")
	}
	defer rows.Close()

	for rows.Next() {
		c := &Clicker{}
		if err := rows.Scan(&c.UserID, &c.ClickCount); err != nil {
			log.Println("scan clicker row: " + err.Error())
			return nil, errors.New("failed to scan clickers")
		}
		cs = append(cs, c)
	}
	if err := rows.Err(); err != nil {
		log.Println("clicker rows err: " + err.Error())
		return nil, errors.New("failed to iterate clicker rows")
	}

	return cs, nil
}

func (p *PersistentDataStore) AddUser(name, email string) (*User, error) {
	var u *User = &User{
		Name:  name,
		Email: email,
	}

	q := `insert into users(user_name, user_email) values($1, $2) returning user_id`
	err := p.db.QueryRow(q, name, email).Scan(&u.ID)
	if err != nil {
		log.Println("add user: " + err.Error())
		return nil, errors.New("unable to add user")
	}
	return p.addClicker(u)
}

func (p *PersistentDataStore) addClicker(u *User) (*User, error) {
	q := `insert into clicks(user_id, click_count) values($1, $2)`
	_, err := p.db.Exec(q, u.ID, 0)
	if err != nil {
		log.Println("add clicker: " + err.Error())
		return nil, errors.New("unable to add clicker")
	}
	return u, nil
}

func (p *PersistentDataStore) GetUserByEmail(email string) (*User, error) {
	var u *User = &User{}

	q := `select user_id, user_email, user_name from users where user_email=$1`
	err := p.db.QueryRow(q, email).Scan(&u.ID, &u.Email, &u.Name)
	if err != nil {
		log.Println("get user by email: " + err.Error())
		return nil, errors.New("unable to get user by email")
	}
	return u, nil
}

func (p *PersistentDataStore) GetUserByName(name string) (*User, error) {
	var u *User = &User{}

	q := `select user_id, user_email, user_name from users where user_name=$1`
	err := p.db.QueryRow(q, name).Scan(&u.ID, &u.Email, &u.Name)
	if err != nil {
		log.Println("get user by name: " + err.Error())
		return nil, errors.New("unable to get user by name")
	}
	return u, nil
}

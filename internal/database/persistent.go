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

func (p *PersistentDataStore) AddClick(userID string) (Clicker, error) {
	var c Clicker

	q := `update "clicks" set "clicks"="clicks"+1 where "user_id"=$1 returning *`
	err := p.db.QueryRow(q, userID).Scan(&c)
	if err != nil {
		log.Println("add click: " + err.Error())
		return c, errors.New("unable to add click")
	}

	return c, nil
}

func (p *PersistentDataStore) GetClickers() ([]Clicker, error) {
	return []Clicker{}, errors.New("not implemented")
}

func (p *PersistentDataStore) AddUser(name, email string) (User, error) {
	var u User
	q := `insert into "users"("user_name", "user_email") values($1, $2) returning user_id`
	err := p.db.QueryRow(q, name, email).Scan(&u)
	if err != nil {
		log.Println("add user: " + err.Error())
		return u, errors.New("unable to add user")
	}
	return u, nil
}

func (p *PersistentDataStore) GetUserByEmail(email string) (User, error) {
	return User{}, errors.New("not implemented")
}

func (p *PersistentDataStore) GetUserByName(name string) (User, error) {
	return User{}, errors.New("not implemented")
}

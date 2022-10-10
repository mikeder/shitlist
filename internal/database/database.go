package database

type ClickStore interface {
	AddClick(userID string) (*Clicker, error)
	GetClickers() ([]*Clicker, error)
}

type UserStore interface {
	AddUser(name, email string) (*User, error)
	GetUserByEmail(userEmail string) (*User, error)
	GetUserByName(userName string) (*User, error)
	GetUserAuthentications(userID string) (*UserAuthentications, error)
}

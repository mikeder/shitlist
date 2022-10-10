package database

type ClickStore interface {
	AddClick(userID string) (*Clicker, error)
	GetClickers() ([]*Clicker, error)
}

type UserStore interface {
	// Users
	AddUser(name, email string) (*User, error)
	GetUserByEmail(userEmail string) (*User, error)
	GetUserByName(userName string) (*User, error)

	// Authentications
	AddUserAuthentication(userID string, provider AuthenticationProvider) (*Authentication, error)
	GetUserAuthentications(userID string) ([]*Authentication, error)
}

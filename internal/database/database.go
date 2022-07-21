package database

type ClickStore interface {
	AddClick(userID string) (Clicker, error)
	GetClickers() ([]Clicker, error)
}

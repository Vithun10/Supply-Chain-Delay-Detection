package models

// User represents a registered user in the system.
type User struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`    // never expose in JSON
	Role         string `json:"role"` // Owner, Admin, Customer
}

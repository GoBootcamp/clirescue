// Copyright 2013

// Package pivotaluser represents a pivotal user
package pivotaluser

// User represents a Pivotal Labs user.
type User struct {
	Username string
	Password string
	APIToken string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Initials string `json:"initials"`
	Timezone struct {
		Kind      string `json:"kind"`
		Offset    string `json:"offset"`
		OlsonName string `json:"olson_name"`
	} `json:"time_zone"`
}

// New returns a *User.
func New() *User {
	return new(User)
}

// SetLogin sets the username and password.
func (u *User) SetLogin(name, pass string) {
	u.Username = name
	u.Password = pass
}

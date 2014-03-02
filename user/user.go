package user

func New() *User {
	return new(User)
}

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

func (u *User) Login(name, pass string) {
	u.Username = name
	u.Password = pass
}

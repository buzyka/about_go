package entity


type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

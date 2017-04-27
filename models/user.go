package models

import "time"

// User Struct
type User struct {
	ID        int
	Name      string
	Email     string
	Password  []byte
	CreatedAt time.Time
}

// Create User
func (user *User) Create() (err error) {
	statement := "insert into users (name, email, password, created_at) values ($1, $2, $3, $4) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID)
	return
}

// GetUser : Test
func GetUser(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from users where email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

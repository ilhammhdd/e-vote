package models

import (
	"github.com/ilhammhdd/e-vote/utils"
)

type User struct {
	Id        uint       `json:"id"`
	Email     NullString `json:"email"`
	Username  NullString `json:"username"`
	Token     NullString `json:"token"`
	CreatedAt NullTime   `json:"created_at"`
	UpdatedAt NullTime   `json:"updated_at"`
}

func (u User) PrimaryKey() map[string]uint {
	return map[string]uint{"id": u.Id}
}

func (u User) ForeignKey() map[string]uint {
	return nil
}

func (u User) Columns() map[string]interface{} {
	return map[string]interface{}{"email": u.Email, "username": u.Username, "token": u.Token, "created_at": u.CreatedAt, "updated_at": u.UpdatedAt}
}

func (u User) GetAll(stmnt string, params ...interface{}) *[]User {
	var users []User
	var user User
	if len(params) == 0 {
		rows, err := utils.DB.Query(stmnt)
		handleError(err)
		for rows.Next() {
			rows.Scan(&user.Id, &user.Email, &user.Username, &user.Token, &user.CreatedAt, &user.UpdatedAt)
			users = append(users, user)
		}
	} else {
		stmntOut, err := utils.DB.Prepare(stmnt)
		handleError(err)
		rows, err := stmntOut.Query(params...)
		handleError(err)
		for rows.Next() {
			rows.Scan(&user.Id, &user.Email, &user.Username, &user.Token, &user.CreatedAt, &user.UpdatedAt)
			users = append(users, user)
		}
	}
	return &users
}

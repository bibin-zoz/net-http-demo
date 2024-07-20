package models

type Admin struct {
	ID       int    `json:id`
	UserName string `json:username`
	Email    string `json:email`
}

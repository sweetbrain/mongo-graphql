// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Email string `json:"email"`
}

type Notification struct {
	ID    string `json:"id"`
	Seen  bool   `json:"seen"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

type UpdateNotification struct {
	ID     string `json:"id"`
	UserID string `json:"userID"`
	Seen   bool   `json:"seen"`
}

type UpdateUser struct {
	ID    string  `json:"id"`
	First *string `json:"first"`
	Last  *string `json:"last"`
	Email *string `json:"email"`
}

type User struct {
	ID            string          `json:"id" bson:"_id"`
	First         string          `json:"first"`
	Last          string          `json:"last"`
	Email         string          `json:"email"`
	Notifications []*Notification `json:"notifications"`
}

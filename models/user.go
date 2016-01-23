package models

type User struct {
	Id        int
	Name      string
	Email     string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

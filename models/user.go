package models

type UserCreate struct {
	Username string `json:"username" db:"username"`
	PWHash   string `json:"pw_hash" db:"pw_hash"`
}

type UserView struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	PWHash   string `json:"-" db:"-"`
}

package models

type User struct {
	ID        uint64 `json:"id" binding:"required" example:"1"`
	Username  string `json:"username" binding:"required" example:"my_username"`
	FirstName string `json:"first_name" binding:"required" example:"Ivan"`
	LastName  string `json:"last_name" binding:"required" example:"Ivanov"`
	BirthDate *Date  `json:"birthdate,omitempty" format:"date" example:"1990-01-01"`
	Sex       Sex    `json:"sex" binding:"required" example:"male"`
	Biography string `json:"biography,omitempty" binding:"omitempty" example:"London is the capital of Great Britain"`
	City      string `json:"city,omitempty" binding:"omitempty,min=1,max=50" example:"Saint-Petersburg"`
}

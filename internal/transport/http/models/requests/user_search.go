package requests

type UserSearchRequest struct {
	FirstName string `json:"first_name" form:"first_name" example:"van" description:"Часть имени"`
	LastName  string `json:"last_name" form:"last_name" example:"vano" description:"Часть фамилии"`
}

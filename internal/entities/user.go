package entities

import "time"

type User struct {
	ID        UserID
	Username  Username
	Password  Password
	FirstName UserFirstName
	LastName  UserLastName
	Birthdate *Birthdate
	Sex       Sex
	Biography Biography
	City      City
}

type UserID uint64
type Username string
type Password string
type UserFirstName string
type UserLastName string
type Birthdate time.Time
type Sex int
type Biography string
type City string

const (
	SexAny    Sex = 0
	SexMale   Sex = 1
	SexFemale Sex = 2
)

type UserFilter struct {
	FirstNameLike string
	LastNameLike  string
}

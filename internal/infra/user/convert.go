package user

import (
	"strings"
	"time"

	"highload-architect/internal/entities"
)

var fields = []string{
	"id",
	"username",
	"password",
	"first_name",
	"last_name",
	"sex",
	"birthdate",
	"biography",
	"city",
}

var fieldsStr = strings.Join(fields, ",")

type dbUser struct {
	ID        int64      `db:"id"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Sex       *int16     `db:"sex"`
	Birthdate *time.Time `db:"birthdate"`
	Biography *string    `db:"biography"`
	City      string     `db:"city"`
}

func convertManyToEntity(users []dbUser) []*entities.User {
	result := make([]*entities.User, len(users))
	for i, u := range users {
		result[i] = convertToEntity(u)
	}

	return result
}

func convertToEntity(u dbUser) *entities.User {
	sex := entities.SexAny
	if u.Sex != nil {
		sex = entities.Sex(*u.Sex)
	}

	biography := ""
	if u.Biography != nil {
		biography = *u.Biography
	}

	return &entities.User{
		ID:        entities.UserID(u.ID),
		Username:  entities.Username(u.Username),
		Password:  entities.Password(u.Password),
		FirstName: entities.UserFirstName(u.FirstName),
		LastName:  entities.UserLastName(u.LastName),
		Birthdate: (*entities.Birthdate)(u.Birthdate),
		Sex:       sex,
		Biography: entities.Biography(biography),
		City:      entities.City(u.City),
	}
}

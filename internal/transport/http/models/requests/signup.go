package requests

import (
	"regexp"

	"highload-architect/internal/entities"
	"highload-architect/internal/transport/http/models"
	pkgapperror "highload-architect/pkg/apperror"
)

var _usernameRe = regexp.MustCompile("^\\w+$")

type SignUpRequest struct {
	Username  string          `json:"username" binding:"required,min=1,max=50" minLength:"1" maxLength:"50" example:"myusername"`
	Password  models.Password `json:"password" binding:"required,min=1,max=50" minLength:"1" maxLength:"50" example:"123456"`
	FirstName string          `json:"first_name" binding:"required,min=1,max=50" minLength:"1" maxLength:"50" example:"Ivan"`
	LastName  string          `json:"last_name" binding:"required,min=1,max=50" minLength:"1" maxLength:"50" example:"Ivanov"`
	Birthdate *models.Date    `json:"birthdate,omitempty" format:"date" example:"2006-01-02"`
	Sex       *models.Sex     `json:"sex,omitempty" example:"male"`
	Biography string          `json:"biography,omitempty" example:"London is the capital of Great Britain"`
	City      string          `json:"city,omitempty" binding:"omitempty,min=1,max=50" example:"Saint-Petersburg"`
}

func (r SignUpRequest) Validate() error {
	if !_usernameRe.MatchString(r.Username) {
		return pkgapperror.ErrIncorrectParameter.WithDetails("username is invalid: alphanumeric and underscore characters only allowed")
	}

	if r.Sex != nil {
		if err := r.Sex.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r SignUpRequest) ConvertUsername() entities.Username {
	return entities.Username(r.Username)
}

func (r SignUpRequest) ConvertPassword() (entities.Password, error) {
	password, err := r.Password.Hash()
	if err != nil {
		return "", err
	}

	return entities.Password(password), nil
}

func (r SignUpRequest) ConvertFirstName() entities.UserFirstName {
	return entities.UserFirstName(r.FirstName)
}

func (r SignUpRequest) ConvertLastName() entities.UserLastName {
	return entities.UserLastName(r.LastName)
}

func (r SignUpRequest) ConvertBirthdate() *entities.Birthdate {
	if r.Birthdate == nil {
		return nil
	}

	res := r.Birthdate.ToEntity()
	return &res
}

func (r SignUpRequest) ConvertSex() entities.Sex {
	if r.Sex == nil {
		return entities.SexAny
	}

	return r.Sex.ToEntity()
}

func (r SignUpRequest) ConvertBiography() entities.Biography {
	return entities.Biography(r.Biography)
}

func (r SignUpRequest) ConvertCity() entities.City {
	return entities.City(r.City)
}

package transformers

import (
	"highload-architect/internal/entities"
	"highload-architect/internal/transport/http/models"
)

func TransformUsers(usrs []*entities.User) models.Users {
	response := make([]models.User, len(usrs))
	for i, u := range usrs {
		response[i] = TransformUser(u)
	}

	return response
}

func TransformUser(usr *entities.User) models.User {
	return models.User{
		ID:        uint64(usr.ID),
		Username:  string(usr.Username),
		FirstName: string(usr.FirstName),
		LastName:  string(usr.LastName),
		BirthDate: (*models.Date)(usr.Birthdate),
		Sex:       models.SexFromEntity(usr.Sex),
		Biography: string(usr.Biography),
		City:      string(usr.City),
	}
}

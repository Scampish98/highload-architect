package models

import (
	"highload-architect/internal/entities"
	pkgapperror "highload-architect/pkg/apperror"
)

type Sex string

const (
	SexAny    Sex = "any"
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

func (s Sex) Validate() error {
	switch s {
	case SexAny, SexMale, SexFemale:
		return nil
	default:
		return pkgapperror.ErrIncorrectParameter.WithDetails("sex is invalid")
	}
}

func (s Sex) ToEntity() entities.Sex {
	switch s {
	case SexAny:
		return entities.SexAny
	case SexMale:
		return entities.SexMale
	case SexFemale:
		return entities.SexFemale
	default:
		return entities.SexAny
	}
}

func SexFromEntity(entity entities.Sex) Sex {
	switch entity {
	case entities.SexMale:
		return SexMale
	case entities.SexFemale:
		return SexFemale
	default:
		return SexAny
	}
}

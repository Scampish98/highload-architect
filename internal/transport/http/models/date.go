package models

import (
	"encoding/json"
	"time"

	"highload-architect/internal/entities"
	pkgapperror "highload-architect/pkg/apperror"
)

type Date time.Time

var DateFormat = "2006-01-02"

func (dt *Date) UnmarshalJSON(bs []byte) error {
	var dateStr string
	err := json.Unmarshal(bs, &dateStr)
	if err != nil {
		return err
	}

	t, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return pkgapperror.ErrIncorrectParameter.WithDetails("date format is invalid")
	}

	*dt = Date(t)
	return nil
}

func (dt Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.String())
}

func (dt Date) String() string {
	return time.Time(dt).Format(DateFormat)
}

func (dt Date) ToEntity() entities.Birthdate {
	return entities.Birthdate(dt)
}

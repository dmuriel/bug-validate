package events

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"

	"bugvalidate/app/models"
)

func Validate1(tx *pop.Connection, event models.Event) *validate.Errors {
	return validate.Validate(
		&validators.FuncValidator{
			Field:   event.CompanyID.String(),
			Name:    "CompanyID",
			Message: "company with id '%s' not exists.",
			Fn: func() bool {
				cQ := tx.Q()
				cQ.Where("id = ?", event.CompanyID)
				cQ.Where("status = ?", "Active")

				exists, err := cQ.Exists(&models.Company{})
				if err != nil {
					fmt.Println("VALIDATE 1: ERROR 1 ----------------------->", err)
					return false
				}

				return exists
			},
		},
		&validators.FuncValidator{
			Name:    "Timestamp",
			Message: "invalid timestamp lapse.%s",
			Fn: func() bool {
				leQ := tx.Q()
				leQ.Where("type = ?", event.Type)

				var le models.Event
				err := leQ.Last(&le)
				if err != nil && !errors.Is(err, sql.ErrNoRows) {
					fmt.Println("VALIDATE 2: ERROR 2 ----------------------->", err)
					return false
				}

				if err != nil && errors.Is(err, sql.ErrNoRows) {
					return true
				}

				diff := event.Timestamp.Sub(le.Timestamp)
				return diff.Seconds() >= 5
			},
		},
	)
}

func Validate2(tx *pop.Connection, event models.Event) *validate.Errors {
	verrs := validate.Validate(
		&validators.FuncValidator{
			Field:   event.CompanyID.String(),
			Name:    "CompanyID",
			Message: "company with id '%s' not exists.",
			Fn: func() bool {
				cQ := tx.Q()
				cQ.Where("id = ?", event.CompanyID)
				cQ.Where("status = ?", "Active")

				exists, err := cQ.Exists(&models.Company{})
				if err != nil {
					return false
				}

				return exists
			},
		},
	)

	valFunc := &validators.FuncValidator{
		Name:    "Timestamp",
		Message: "invalid timestamp lapse.%s",
		Fn: func() bool {
			leQ := tx.Q()
			leQ.Where("type = ?", event.Type)

			var le models.Event
			err := leQ.Last(&le)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				fmt.Println("ERROR----------------------->", err)
				return false
			}

			if err != nil && errors.Is(err, sql.ErrNoRows) {
				return true
			}

			diff := event.Timestamp.Sub(le.Timestamp)
			return diff.Seconds() >= 5
		},
	}

	verrs.Append(validate.Validate(valFunc))

	return verrs
}

// ! This works to replicate conn busy error + index out of range error.
func Validate3(tx *pop.Connection, event models.Event) *validate.Errors {
	return validate.Validate(
		&validators.FuncValidator{
			Field:   event.CompanyID.String(),
			Name:    "CompanyID",
			Message: "company with id '%s' not exists.",
			Fn: func() bool {
				cQ := tx.Q()
				cQ.Where("id = ?", event.CompanyID)
				cQ.Where("status = ?", "Active")

				exists, err := cQ.Exists(&models.Company{})
				if err != nil {
					fmt.Println("VALIDATE 1: ERROR 1 ----------------------->", err)
					return false
				}

				return exists
			},
		},
		&validators.FuncValidator{
			Name:    "Timestamp",
			Message: "invalid timestamp lapse.%s",
			Fn: func() bool {
				leQ := tx.Q()
				leQ.Where("type = ?", event.Type)

				var le models.Event
				_, err := leQ.Exists(&le)
				if err != nil && !errors.Is(err, sql.ErrNoRows) {
					fmt.Println("VALIDATE 2: ERROR 2 ----------------------->", err)
					return false
				}

				if err != nil && errors.Is(err, sql.ErrNoRows) {
					return true
				}

				return true
			},
		},
	)
}

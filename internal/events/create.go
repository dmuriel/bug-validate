package events

import (
	"bugvalidate/app/models"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

func Create1(tx *pop.Connection, event models.Event) (*validate.Errors, error) {
	verrs := Validate1(tx, event)
	if verrs.HasAny() {
		return verrs, nil
	}

	if err := tx.Create(&event); err != nil {
		return nil, err
	}

	return nil, nil
}

func Create2(tx *pop.Connection, event models.Event) (*validate.Errors, error) {
	verrs := Validate2(tx, event)
	if verrs.HasAny() {
		return verrs, nil
	}

	if err := tx.Create(&event); err != nil {
		return nil, err
	}

	return nil, nil
}

func Update1(tx *pop.Connection, event models.Event) (*validate.Errors, error) {
	verrs := Validate3(tx, event)
	if verrs.HasAny() {
		return verrs, nil
	}

	if err := tx.Update(&event); err != nil {
		return nil, err
	}

	return nil, nil
}

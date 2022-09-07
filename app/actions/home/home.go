package home

import (
	"bugvalidate/app/models"
	"bugvalidate/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"

	"bugvalidate/internal/events"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func Index(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func Create1(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	event := models.Event{}

	if err := c.Bind(&event); err != nil {
		return err
	}

	verrs, err := events.Create1(tx, event)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Logger().Error(verrs)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("home/index.plush.html"))
	}

	return c.Redirect(http.StatusSeeOther, "homeIndexPath()")
}

func Update1(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	event := models.Event{}

	if err := c.Bind(&event); err != nil {
		return err
	}

	verrs, err := events.Update1(tx, event)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Logger().Error(verrs)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("home/index.plush.html"))
	}

	return c.Redirect(http.StatusSeeOther, "homeIndexPath()")
}

func Create2(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	event := models.Event{}

	if err := c.Bind(&event); err != nil {
		return err
	}

	verrs, err := events.Create2(tx, event)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Logger().Error(verrs)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("home/index.plush.html"))
	}

	return c.Redirect(http.StatusSeeOther, "homeIndexPath()")
}

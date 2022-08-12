package home_test

import (
	"bugvalidate/app"
	"bugvalidate/app/models"
	"net/http"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(app.New())}
	suite.Run(t, as)
}

func (as *ActionSuite) Test_Home_Create1() {
	company := models.Company{Name: "Test Company", Status: "Active"}
	as.NoError(as.DB.Create(&company))

	as.DBDelta(1, "events", func() {
		values := models.Event{CompanyID: company.ID}

		res := as.HTML("/create-1").Post(values)
		as.Equal(http.StatusSeeOther, res.Code)
	})
}

/*
func (as *ActionSuite) Test_Home_Create2() {
	company := models.Company{Name: "Test Company", Status: "Active"}
	as.NoError(as.DB.Create(&company))

	as.DBDelta(1, "events", func() {
		values := models.Event{CompanyID: company.ID}

		res := as.HTML("/create-2").Post(values)
		as.Equal(http.StatusSeeOther, res.Code)
	})
}
*/

package actions_test

import (
	"bugvalidate/app"
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

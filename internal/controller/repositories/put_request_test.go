package repositories

import (
	"testing"

	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/stretchr/testify/assert"
)

type PutRequestValidateTestCase struct {
	ID          int
	Name        string
	Link        string
	ExpectedErr error
}

var PutRequestValidateTestCases = map[string]PutRequestValidateTestCase{
	"should return nil error": {
		ID:          1,
		Name:        "Test Name",
		Link:        "https://test.com",
		ExpectedErr: nil,
	},
	"should return error on empty name": {
		ID:          1,
		Name:        "",
		Link:        "http://test.com",
		ExpectedErr: errors.NewParamErr("name is missing"),
	},
}

//TODO: implement test
func TestPutRequestValidate(t *testing.T) {
	for tName, tCase := range PutRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &putRequest{
				ID:   tCase.ID,
				Name: tCase.Name,
				Link: tCase.Link,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

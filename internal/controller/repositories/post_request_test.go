package repositories

import (
	"testing"

	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/stretchr/testify/assert"
)

type PostRequestValidateTestCase struct {
	Name        string
	Link        string
	ExpectedErr error
}

var PostRequestValidateTestCases = map[string]PostRequestValidateTestCase{
	"should return nil error": {
		Name:        "Test Name",
		Link:        "https://test.com",
		ExpectedErr: nil,
	},
	"should return error on empty name": {
		Name:        "",
		Link:        "http://test.com",
		ExpectedErr: errors.NewParamErr("name must be provided"),
	},
}

//TODO: implement test
func TestPostRequestValidate(t *testing.T) {
	for tName, tCase := range PostRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &postRequest{
				Name: tCase.Name,
				Link: tCase.Link,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

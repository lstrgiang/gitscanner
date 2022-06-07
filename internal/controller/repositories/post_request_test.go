package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PostRequestValidateTestCase struct {
	ExpectedErr error
}

var PostRequestValidateTestCases = map[string]PostRequestValidateTestCase{
	"should return nil error": {},
	"should return error on ": {},
}

//TODO: implement test
func TestPostRequestValidate(t *testing.T) {
	for tName, tCase := range PostRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &postRequest{}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

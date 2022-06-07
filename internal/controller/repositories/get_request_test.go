package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetRequestValidateTestCase struct {
	ExpectedErr error
}

var GetRequestValidateTestCases = map[string]GetRequestValidateTestCase{
	"should return nil error": {},
	"should return error on ": {},
}

//TODO: implement test
func TestGetRequestValidate(t *testing.T) {
	for tName, tCase := range GetRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &getRequest{}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

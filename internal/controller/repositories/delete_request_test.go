package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DeleteRequestValidateTestCase struct {
	ExpectedErr error
}

var DeleteRequestValidateTestCases = map[string]DeleteRequestValidateTestCase{
	"should return nil error": {},
	"should return error on ": {},
}

//TODO: implement test
func TestDeleteRequestValidate(t *testing.T) {
	for tName, tCase := range DeleteRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &deleteRequest{}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PutRequestValidateTestCase struct {
	ExpectedErr error
}

var PutRequestValidateTestCases = map[string]PutRequestValidateTestCase{
	"should return nil error": {},
	"should return error on ": {},
}

//TODO: implement test
func TestPutRequestValidate(t *testing.T) {
	for tName, tCase := range PutRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &putRequest{}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

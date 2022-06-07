package repositories

import (
	"math/rand"
	"testing"

	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/stretchr/testify/assert"
)

type DeleteRequestValidateTestCase struct {
	ID          int
	ExpectedErr error
}

var DeleteRequestValidateTestCases = map[string]DeleteRequestValidateTestCase{
	"should return nil error": {
		ID:          rand.Intn(100) + 1,
		ExpectedErr: nil,
	},
	"should return error on id less than or equal 0": {
		ID:          rand.Intn(100) - 100,
		ExpectedErr: errors.InvalidIdError,
	},
}

//TODO: implement test
func TestDeleteRequestValidate(t *testing.T) {
	for tName, tCase := range DeleteRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &deleteRequest{
				ID: tCase.ID,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

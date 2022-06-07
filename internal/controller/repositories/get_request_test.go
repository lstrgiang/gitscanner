package repositories

import (
	"testing"

	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/stretchr/testify/assert"
)

type GetRequestValidateTestCase struct {
	ID          int
	Page        int
	Limit       int
	ExpectedErr error
}

var GetRequestValidateTestCases = map[string]GetRequestValidateTestCase{
	"should return nil error on id provided": {
		ID:          1,
		Page:        0,
		Limit:       0,
		ExpectedErr: nil,
	},
	"should return nil error on page and limit provided": {
		ID:          0,
		Page:        1,
		Limit:       1,
		ExpectedErr: nil,
	},
	"should return error on empty params": {
		ID:          0,
		Page:        0,
		Limit:       0,
		ExpectedErr: errors.InvalidRequest,
	},
	"should return error on invalid id param": {
		ID:          -1,
		Page:        0,
		Limit:       0,
		ExpectedErr: errors.InvalidIdError,
	},
	"should return error on invalid page": {
		ID:          0,
		Page:        -1,
		Limit:       1,
		ExpectedErr: errors.NewParamErr("page must be non negative"),
	},
	"should return error on invalid limit": {
		ID:          0,
		Page:        1,
		Limit:       -1,
		ExpectedErr: errors.NewParamErr("limit must be non negative"),
	},
}

//TODO: implement test
func TestGetRequestValidate(t *testing.T) {
	for tName, tCase := range GetRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &getRequest{
				ID:    tCase.ID,
				Page:  tCase.Page,
				Limit: tCase.Limit,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}

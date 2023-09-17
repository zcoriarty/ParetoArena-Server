package query

import (
	"net/http"

	"github.com/zcoriarty/Backend/apperr"
	"github.com/zcoriarty/Backend/model"
)

// List prepares data for list queries
func List(u *model.AuthUser) (*model.ListQuery, error) {
	switch true {
	case int(u.Role) <= 2: // user is SuperAdmin or Admin
		return nil, nil
	default:
		return nil, apperr.New(http.StatusForbidden, "Forbidden")
	}
}

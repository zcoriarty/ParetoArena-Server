package magic

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/client"
	"github.com/magiclabs/magic-admin-go/token"
	"github.com/zcoriarty/Backend/apperr"
	"github.com/zcoriarty/Backend/config"
)

// NewMagic creates a new magic service implementation
func NewMagic(config *config.MagicConfig) *Magic {
	return &Magic{config}
}

// Magic provides a magic service implementation
type Magic struct {
	config *config.MagicConfig
}

// IsValidToken validates a token with magic link
func (m *Magic) IsValidToken(tkn string) (*token.Token, error) {
	authBearer := "Bearer"
	fmt.Printf("%s", authBearer)
	if tkn == "" {
		return nil, apperr.New(http.StatusUnauthorized, "Bearer token is required")
	}

	if !strings.HasPrefix(tkn, authBearer) {
		return nil, apperr.New(http.StatusUnauthorized, "Bearer token is required")
	}

	did := tkn[len(authBearer)+1:]
	if did == "" {
		return nil, apperr.New(http.StatusUnauthorized, "DID token is required")
	}

	tk, err := token.NewToken(did)
	if err != nil {

		return nil, apperr.New(http.StatusUnauthorized, "Malformed DID token error: "+err.Error())
	}

	if err := tk.Validate(); err != nil {
		return nil, apperr.New(http.StatusUnauthorized, "DID token failed validation: "+err.Error())
	}

	return tk, nil
}

// GetIssuer retrieves the issuer from token
// func (m *Magic) GetIssuer(c *gin.Context) error
func (m *Magic) GetIssuer(tk *token.Token) (*magic.UserInfo, error) {
	client := client.New(m.config.Secret, magic.NewDefaultClient())
	userInfo, err := client.User.GetMetadataByIssuer(tk.GetIssuer())

	if err != nil {
		return nil, apperr.New(http.StatusBadRequest, "Bad request")
	}

	return userInfo, nil
}

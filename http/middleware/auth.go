package middleware

import (
	"net/http"
	"strings"

	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/utils/jwt"
)

const (
	CONTEXT_TOKEN_KEY = "TOTOVAL_CONTEXT_TOKEN"
)

type TokenRevokeError struct{}

func (e TokenRevokeError) Error() string {
	return "token revoke failed"
}
/**
AuthRequired授权
*/
func AuthRequired() request.HandlerFunc {
	return func(c request.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.Request().Header.Get("Authorization")
			if s := strings.Split(token, " "); len(s) == 2 {
				token = s[1]
			}
		}

		// set token
		c.Set(CONTEXT_TOKEN_KEY, token)

		j := jwt.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				if token, _err := j.RefreshTokenUnverified(token); _err == nil {
					if claims, err := j.ParseToken(token); err == nil {
						c.SetAuthClaim(claims)
						c.Header("Authorization", "Bear "+token)
						return
					}
				}
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, toto.V{"error": err.Error()})
			return
		}
		c.SetAuthClaim(claims)
	}
}

/**
撤销 token
 */
func Revoke(c request.Context) error {
	j := jwt.NewJWT(signKey)
	if tokenString, exist := c.Get(CONTEXT_TOKEN_KEY); exist {
		if token, ok := tokenString.(string); ok {
			if err := j.RevokeToken(token); err == nil {
				c.Header("Authorization", "")
				return nil
			}
		}
	}
	return TokenRevokeError{}
}

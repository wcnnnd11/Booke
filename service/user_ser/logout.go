package user_ser

import (
	"GVB_server/service/redis_ser"
	"GVB_server/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {

	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)

}

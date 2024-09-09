package main

import (
	"GVB_server/core"
	"GVB_server/global"
	"GVB_server/utils/jwts"
	"fmt"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "lijiang",
		NickName: "xxx",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxpamlhbmciLCJuaWNrX25hbWUiOiJ4eHgiLCJyb2xlIjoxLCJ1c2VyX2lkIjoxLCJleHAiOjE3MjEyOTczNjAuMzA4NDQyOCwiaXNzIjoieHgifQ.Bn98XAyDl7iB-j5xXkJba48BDs5sStHfpqRe61y8wZw")
	fmt.Println(claims)
}

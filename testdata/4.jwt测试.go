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

	claims, err := jwts.ParseToken("\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsIm5pY2tfbmFtZSI6ImFkbWluIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiYXZhdGFyIjoiIiwiZXhwIjoxNzM2Nzc3NTA0LjExMzY4NywiaXNzIjoiMTIifQ.GAU8GRf_8jM06psPyWuwXu2KLSnHWjNxfDjeRYFw86E")
	fmt.Println(claims)
}

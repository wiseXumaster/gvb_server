package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: "hbbxht",
		NickName: "wise",
		Role:     1,
		UserID:   1,
	})
	fmt.Println(token, err)
	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImhiYnhodCIsIm5pY2tfbmFtZSI6Indpc2UiLCJyb2xlIjoxLCJ1c2VyX2lkIjoxLCJleHAiOjE3MzE1Nzk4MzcuMDMxNjIyLCJpc3MiOiIxMjM0In0._mhP9SYu5Odk9qp8-8cj9NbO0lsshqLeRGW2t0MbSwc")
	fmt.Println(claims, err)
}

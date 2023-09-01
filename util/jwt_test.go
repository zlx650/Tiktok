package util

import "testing"

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjQsIlVzZXJOYW1lIjoiMTIzNDU2Nzg5MDExIiwiUGFzc3dvcmQiOiIxMjM0NTY3IiwiaXNzIjoiR29Ib21lIiwiZXhwIjoxNjkzNTcwOTgxLCJuYmYiOjE2OTM1NjczODF9.QXMq7BoSSpSjtDW3YbBdFzOfvia_fR6frHCdn6razrw"
	parseToken, _ := ParseToken(token)
	println(parseToken.UserId)
	println(parseToken.UserName)
	println(parseToken.Password)
}

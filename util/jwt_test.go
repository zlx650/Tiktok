package util

import "testing"

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJVc2VyTmFtZSI6IjMyMTQ1Njg3MjEiLCJQYXNzd29yZCI6IjEyMzQ1NjciLCJpc3MiOiJHb0hvbWUiLCJleHAiOjE2OTM1NjM3MzQsIm5iZiI6MTY5MzU2MDEzNH0.zoGi5H3OxahK5iQgiYFjp-qZ0YDkRns_5YZRyC-jIVw"

	parseToken, _ := ParseToken(token)
	println(parseToken.UserId)
	println(parseToken.UserName)
	println(parseToken.Password)
}

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

//func TestCreateToken(t *testing.T) {
//
//}
//
//func TestParseToken(t *testing.T) {
//
//}

func TestJWTMiddleWareGet(t *testing.T) {
	token, _ := CreateToken(1)
	r := gin.Default()
	r.Use(JWTMiddleWare())
	r.GET("/test", func(context *gin.Context) {
		userID, exists := context.Get("user_id")
		assert.True(t, exists)
		assert.Equal(t, int64(1), userID)
		context.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	target := fmt.Sprintf("/test?token=%s", token)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"success"}`, w.Body.String())
}

func TestJWTMiddleWarePost(t *testing.T) {
	token, _ := CreateToken(18)
	r := gin.Default()
	r.Use(JWTMiddleWare())

	r.POST("/test", func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		assert.True(t, exists)
		assert.Equal(t, int64(18), userID)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	req.PostForm = url.Values{"token": {token}}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"success"}`, w.Body.String())
}

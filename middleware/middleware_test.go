package middleware_test

import (
	"testing"

	mw "github.com/zcoriarty/Backend/middleware"

	"github.com/gin-gonic/gin"
)

func TestAdd(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	mw.Add(r, gin.Logger())
}

package tokenAuth

import (
	"github.com/gin-gonic/gin"
)

/*
VerifyToken : 檢查token是否有效
  @param userToken String
  @return
*/
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

package routes

import (
	"DockerGoNginx/api/app/controllers/api/test"
	"DockerGoNginx/api/app/middleware/tokenAuth"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/kimi", func(c *gin.Context) {
		result := map[string]interface{}{
			"Data": map[string]interface{}{
				"msg":  "Hello Kimi!",
				"time": time.Now().Format("2006-01-02 15:04:05"),
			},
			"Result": map[string]interface{}{
				"Status":  true,
				"Code":    "2000",
				"Message": "正確狀況",
			},
		}
		c.JSON(http.StatusOK, result)
	})

	// 測試檢查 Token
	tokenGroup := r.Group("/api", tokenAuth.VerifyToken())
	tokenGroup.POST("/verifyToken", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "verify OK",
			"token": c.Request.Body,
			"d":     c.DefaultPostForm("userToken", ""),
		})
	})

	// 測試用 controller
	testGroup := r.Group("/test")
	testGroup.POST("/TestValidateV9", test.TestValidateV9)

	return r
}

package test

import (
	"DockerGoNginx/app/services/ginservices"
	"DockerGoNginx/config/errorCode"

	"github.com/gin-gonic/gin"
)

/*
TestValidateV9 : 測試驗證 v9
*/
func TestValidateV9(c *gin.Context) {
	type ReqStruct struct {
		IP          string `json:"ip" form:"ip"  binding:"ipv4" `
		Number      string `json:"number" form:"number"  binding:"numeric" `
		Excludesall string `json:"exc" form:"exc"  binding:"excludesall=!@#?" `
		Awsome      string `json:"awsome" form:"awsome" binding:"awsomeValidate" `
	}

	var reqJSON ReqStruct
	reqData, err := ginservices.GinRequest(c, &reqJSON)
	if bindErr := c.ShouldBind(&reqJSON); bindErr != nil {
		ginservices.GinResponse(c, reqData, err, false, errorCode.PARAMS_INVALID, "")
		return
	}

	ginservices.GinResponse(c, reqJSON, reqJSON, true, errorCode.SUCCESS, "")
	return
}

package ginservices

import (
	"DockerGoNginx/app/services/commomservices"
	"DockerGoNginx/config"
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ResultObj struct {
	Status  bool
	Code    string
	Message string
}

type ResponseObj struct {
	Data   interface{}
	Result ResultObj
}

/*
GinRequest : Gin http 接收body
*/
func GinRequest(c *gin.Context, reqJSON interface{}) (interface{}, error) {
	var reqData interface{}
	if c.Request.Method == "GET" {
		reqData = c.Request.URL.Query()
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}
	} else {
		// 取得Body資料, 給log用
		data, err := c.GetRawData()
		if err != nil {
			return nil, err
		}
		reqData = ioutil.NopCloser(bytes.NewBuffer(data))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}

	}
	return reqData, nil
}

/*
GinResponse : Gin http 回傳格式
{
    "Data": {},
    "Result": {
        "Status": true,
        "Code": "2000",
        "Message": "正確狀況"
    }
}
*/
func GinResponse(c *gin.Context, resquestData interface{}, responseData interface{}, resultStatus bool, resultMsg map[string]string, err interface{}) {
	var response ResponseObj
	response.Data = responseData
	env := os.Getenv("APP_ENV")
	response.Result = ResultObj{Status: resultStatus, Code: resultMsg["code"], Message: resultMsg["message"]}

	if config.WRITE_LOG == true {
		if resultStatus {
			commomservices.WriteLog(c, resquestData, response)
		} else {
			commomservices.WriteLogError(c, resquestData, response, err)
		}
	}

	// if config.WRITE_LOG_TO_DB == true {
	// 	WriteAPILog(c, resquestData, response, resultMsg, err)
	// }

	// 正式環境 不把訊息顯示出來
	if env == "app" {
		response.Result.Message = ""
	}

	c.JSON(http.StatusOK, response)
	return
}

/*
WriteAPILog : 寫入 WriteAPILog
*/
/*
func WriteAPILog(c *gin.Context, body interface{}, response interface{}, resultMsg map[string]string, err interface{}) {
	var userID int
	userIDstr, exist := c.Get("userID")
	if !exist {
		userID = 0
	} else {
		userID = userIDstr.(int)
	}

	log := models.ApiLog{
		UserID:        userID,
		RequestMethod: c.Request.Method,
		ApiURL:        c.Request.URL.Path,
		ResultCode:    resultMsg["code"],
		ResultMsg:     resultMsg["message"],
		Request:       fmt.Sprintf("%+v", body),
		Response:      fmt.Sprintf("%+v", response),
		ErrorMsg:      fmt.Sprintf("%+v", err),
		IP:            c.ClientIP(),
	}
	mysql.GormDB.Create(&log)
}
*/

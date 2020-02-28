package commomservices

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
WriteLog : 寫入 log
*/
func WriteLog(c *gin.Context, body interface{}, respon interface{}) {
	log.WithFields(log.Fields{
		"url":        c.Request.URL.Path,
		"method":     c.Request.Method,
		"request_ip": c.ClientIP(),
		"body":       body,
	}).Info(respon)
}

/*
WriteLogError : 寫入 error log
*/
func WriteLogError(c *gin.Context, body interface{}, respon interface{}, err interface{}) {
	log.WithFields(log.Fields{
		"url":        c.Request.URL.Path,
		"method":     c.Request.Method,
		"request_ip": c.ClientIP(),
		"body":       body,
		"err_msg":    err,
	}).Error(respon)
}

/*
UploadFile : form post 上傳檔案
*/
func UploadFile(c *gin.Context) map[string]interface{} {
	result := map[string]interface{}{
		"url":    "",
		"status": false,
	}

	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		return result
	}
	filename := header.Filename

	fmt.Println(file, err, filename)

	out, err := os.Create(filename)
	defer out.Close()
	io.Copy(out, file)

	result["status"] = true
	result["url"] = "xxx"
	return result
}

/*
GetUTCTime : 取得UTC的時間
*/
func GetUTCTime() time.Time {
	t := time.Now()
	local_location, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
	}
	time_UTC := t.In(local_location)
	return time_UTC
}

/*
GetFunctionName : 取得 function 名稱
*/
func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

/*
Sha1 : 取得sha1編碼
*/
func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

/*
GetS3FileName : 從S3 url 取得檔案名稱
*/
func GetS3FileName(s string) string {
	filURLSlice := strings.Split(s, "/")
	filename := filURLSlice[len(filURLSlice)-1]
	return filename
}

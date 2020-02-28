package main

import (
	"net"
	"time"

	"DockerGoNginx/app/validation/customValidate"
	"DockerGoNginx/app/validation/customValidateV9"
	"DockerGoNginx/config"
	"DockerGoNginx/routes"
	"net/http"
	"os"

	// db "DockerGoNginx/config/databases/mysql"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.AddHook(newRotateHook("logs", "stdout.log", 3*24*time.Hour, 24*time.Hour))
}

/*
newRotateHook : 寫log 的 hook 綁定時間
*/
func newRotateHook(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
	writer, err := rotatelogs.New(
		logPath+"/"+"%Y%m%d_"+logFileName,
		// logPath+"/"+"%Y%m%d_%H%M_"+logFileName,

		// WithLinkName 為最新的log建立連結
		rotatelogs.WithLinkName(logFileName),

		// WithRotationTime 設置log分割的時間
		rotatelogs.WithRotationTime(rotationTime),

		// WithMaxAge和WithRotationCount 只能設置一個
		// WithMaxAge 文件清理前最長保存時間
		// WithRotationCount 文件清理前最多保存個數
		rotatelogs.WithMaxAge(maxAge),
		// rotatelogs.WithRotationCount(maxAge),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}

func main() {
	// defer db.GormDB.Close()
	// defer db.SqlDB.Close()

	// GIN binding validation version
	switch config.VALIDATION_VERSION {
	case "v8":
		customValidate.Start()
	case "v9":
		customValidateV9.Start()
	default:
		customValidateV9.Start()
	}
	r := routes.SetupRouter()

	// GIN environment file
	var env = "dev"
	if len(os.Args) > 1 {
		env = os.Args[1]
		switch env {
		case "app":
			godotenv.Load(".env")
		case "dev":
			godotenv.Load("./.env.dev")
		case "qa":
			godotenv.Load("./.env.qa")
		default:
			godotenv.Load("./.env")
		}
	} else {
		godotenv.Load("./.env")
	}

	// Listen and Server
	port := os.Getenv("APP_URL")
	addrs, _ := net.InterfaceAddrs()

	log.WithFields(log.Fields{
		"ip":   addrs[0],
		"port": port,
		"env":  env,
	}).Info("===== Start Server ===== ")

	// r.Run(port)
	log.Fatal(http.ListenAndServe(port, r))
}

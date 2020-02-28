package customValidate

import (
	"DockerGoNginx/api/app/validation/validatefunction"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func Start() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("awsomeValidate", validatefunction.AwsomeValidate)
	}
}

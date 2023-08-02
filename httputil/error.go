package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"status ok"`
	Detail  string `json:"Detail" example:"id 30 is deleted"`
}

type HTTPFileSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Your file has been successfully uploaded."`
}

type HTTPFileError5 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Unable to save the file"`
}

type HTTPFileError4 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"No file is received"`
}

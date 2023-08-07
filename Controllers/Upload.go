package Controllers

import (
	"first-api/httputil"
	"net/http"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")

	file, err := c.FormFile("file")

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	extension := filepath.Ext(file.Filename)

	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "/uploaded/"+newFileName); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Your file has been successfully uploaded.",
	})
}

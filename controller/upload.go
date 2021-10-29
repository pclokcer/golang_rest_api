package controller

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/random"
)

type UploadController interface {
	Upload(c *gin.Context)
}

type uploadController struct {
}

func NewUploadController() UploadController {
	return &uploadController{}
}

func (upload uploadController) Upload(c *gin.Context) {

	// istekten gelen resim alınıyor
	file, header, err := c.Request.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"succes": false})
		return
	}

	// file adı uzantısına göre ayrılmak isteniyor
	split := strings.Split(header.Filename, ".")

	// file adı random olarak verildi ve dosya kaydedildi
	out, err := os.Create("./storage/tmp/" + random.String(25) + "." + split[len(split)-1])

	if err != nil {
		panic(err)
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"succes": true})
}

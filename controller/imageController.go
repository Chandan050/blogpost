package controller

import (
	"log"
	"math/rand"

	"github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randomletter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}

func Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatalln(err)
		}
		files := form.File["image"] // <file key>
		filename := ""
		for _, file := range files {
			filename = randomletter(5) + "-" + file.Filename
			if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
				c.AbortWithError(400, err)
			}
		}
		c.JSON(200, gin.H{"url": "http:/localhost:8080/app/uploads/" + filename, "file": filename})
		return
	}
}

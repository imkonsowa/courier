package main

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	parser := gin.Default()

	parser.POST("/upload", func(context *gin.Context) {
		file, _, err := context.Request.FormFile("file")

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to process your request",
			})
			return
		}

		if file == nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "File is missing!",
			})
			return
		}

		r := csv.NewReader(file)
		r.ReuseRecord = true

		all, err := csv.NewReader(file).Read()

		if len(all) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Empty files not allowed",
			})
			return
		}

		if len(all) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Empty files not allowed",
			})
			return
		}

		go func() {
			// process the lines
			fmt.Println("File lines", len(all))
		}()

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Well received, Processing!",
		})
		return
	})

	parser.Run(":1996")
}

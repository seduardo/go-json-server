package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	data = gin.H{
		"pessoas": []gin.H{
			gin.H{"id": 1, "name": "sergio"},
			gin.H{"id": 2, "name": "eduardo"},
			gin.H{"id": 3, "name": "diones"},
		},
		"animais": []gin.H{
			gin.H{"id": 1, "name": "zequinha"},
			gin.H{"id": 2, "name": "eva"},
		},
	}
)

func index(c *gin.Context) {
	c.String(200, "Index")
}

func findData(name string) []gin.H {
	t, ok := data[name]
	if ok {
		return t.([]gin.H)
	}
	return []gin.H{}
}

func listRows(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"data": findData(c.Param("table")),
	})
}

func findByID(table string, id int) gin.H {
	t, ok := data[table]
	if !ok {
		return gin.H{}
	}
	for _, r := range t.([]gin.H) {
		if r["id"].(int) == id {
			return r
		}
	}
	return gin.H{}
}

func atoi(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}
	return i
}

func listByID(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"data": findByID(c.Param("table"), atoi(c.Param("id"))),
	})
}

func main() {
	fmt.Println("Go-JSON-Server - Version 1.0")

	gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", index)
	r.GET("/:table", listRows)
	r.GET("/:table/:id", listByID)

	r.Run(":9090")
}

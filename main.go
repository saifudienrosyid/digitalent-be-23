package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"

)


func main(){
	router := gin.Default()
	//router.GET("/", getSomething)
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status", GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/:idAntrian/delete", DeleteAntrianHandler)
	router.Run(":8080")
}

// func getSomething(){
// 	c.JSON(http.statusOK, map[string]interface{}{
// 		"body1": "Something succes",
// 	})

// 	return
// }



package controllers

import (
	"github.com/gin-gonic/gin"
	"institute/config"
	// "fmt"
)

/*
 * Struct to handle response.
 *
 * Used by controllers to generate output.
 *
 */
type ResponseController struct {
	Code      int         `json:"code"`
	ApiStatus int         `json:"api_status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

/*
 * Function to generate the response in json.
 *
 * Params c 	   type  *gin.Context
 * & 	  response type ResponseController
 *
 * Generates json
 */
func GetResponse(c *gin.Context, response ResponseController) {
	defer c.Request.Body.Close()
	// fmt.Println(response)
	c.JSON(200, gin.H{
		config.Response: response,
	})
}

/*
 * Struct to handle response.
 *
 * Used by controllers to generate output.
 *
 */
type ResponseControllerList struct {
	Code        int         `json:"code"`
	ApiStatus   int         `json:"api_status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	TotalRecord interface{} `json:"total_record,omitempty"`
}

/*
 * Function to generate the response in json.
 *
 * Params c 	   type  *gin.Context
 * & 	  response type ResponseControllerList
 *
 * Generates json
 */
func GetResponseList(c *gin.Context, response ResponseControllerList) {
	defer c.Request.Body.Close()
	c.JSON(200, gin.H{
		config.Response: response,
	})
}

func UnauthorizedAccessResponse(c *gin.Context) {
	response := ResponseController{
		config.UnauthorizedCode,
		config.UnauthorizedStatus,
		config.UnauthorizedMsg,
		nil,
	}
	GetResponse(c, response)
}

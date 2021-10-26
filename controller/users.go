package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqBody struct {
	Test1 string `json:"test1"`
}

type Response struct {
	Body  interface{} `json: "body"`
	Param string      `json: "param"`
	Query string      `json: "query"`
}

func GetUsers(ctx *gin.Context) {
	body := ReqBody{}

	//Body Parameters
	err := ctx.BindJSON(&body)
	//Path Parameters
	param, _ := ctx.Params.Get("param")
	//Query Parameters
	query := ctx.Request.URL.Query().Get("query")

	if err != nil {
		panic(err)
	}

	var responseBody Response

	responseBody.Body = body
	responseBody.Param = param
	responseBody.Query = query

	fmt.Println(responseBody)

	ctx.JSON(http.StatusOK, responseBody)
}

package ret_code

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	ExtraMsg string `json:"extra_msg"`
}

func MakeSuc(c *gin.Context, data interface{}) {
	resp := Response{
		http.StatusOK,
		"ok",
		data,
	}
	c.JSON(http.StatusOK, resp)
}

func MakeRet(c *gin.Context, code ErrorResponse, extraMsg string) {
	if extraMsg != "" {
		code.ExtraMsg = extraMsg
	}
	c.JSON(http.StatusOK, code)
}

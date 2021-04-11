package controller

import (
	"api/model/ret_code"
	"github.com/gin-gonic/gin"
)

func Resp(c *gin.Context) {
	ret_code.MakeSuc(c, gin.H{"msg": "ok"})
}

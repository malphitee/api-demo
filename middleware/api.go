package middleware

import (
	"api/config"
	"api/model/ret_code"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SignatureParams struct {
	ModuleId  string `json:"module_id" binding:"required"`
	TimeStamp int64  `json:"timestamp" binding:"required"`
	Signature string `json:"signature" binding:"required"`
	Nonce     string `json:"nonce" binding:"required"`
	DataStr   string `json:"data_str" binding:"required"`
}

func ApiCheckSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signatureParams SignatureParams
		if err := c.ShouldBind(&signatureParams); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, ret_code.PARAMS_FORMAT_INVALID)
			log.Println(err.Error())
		}
		fmt.Println(signatureParams)
		var EnableCheckSignature bool
		if developConfig, err := config.GetDevelopConfig(); err != nil {
			EnableCheckSignature = true
		} else {
			EnableCheckSignature = developConfig.EnableCheckSignature
		}
		if EnableCheckSignature && !checkSignature(signatureParams) {
			c.AbortWithStatusJSON(http.StatusOK, ret_code.SIGNATURE_ERROR)
		}
	}
}

func checkSignature(params SignatureParams) bool {
	// 1. 判断参数中的时间戳是否过期
	timeStamp := time.Unix(params.TimeStamp, 0)
	if time.Since(timeStamp) > 3*time.Minute {
		return false
	}
	// @todo 2.从数据库中查找module_id，确认是否注册过，没注册直接返回错误
	moduleKey := "8ax431fg"
	// 3. 找到module信息的情况下，根据失效时间判断是否已失效，失效则删除模块信息
	// 4. 校验签名，根据module的key，用参数中提供的数据计算签名，与传来的签名对比，看是否相同，计算方法，sha1
	signStr := params.ModuleId + strconv.FormatInt(params.TimeStamp, 10) + params.Nonce + params.DataStr + moduleKey
	encryptedStr := sha1.New().Sum([]byte(signStr))
	return params.Signature == string(encryptedStr)
}

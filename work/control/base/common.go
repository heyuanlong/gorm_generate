package base

import (
	"encoding/json"
	"net/http"

	kinit "gorm_generate/initialize"
	kcode "gorm_generate/work/code"

	"github.com/gin-gonic/gin"
)

type DataIStruct struct {
	Status         int         `json:"code"`
	Info           string      `json:"info"`
	Data           interface{} `json:"data"`
	InnerException string      `json:"innerException"`
}

func init() {

}

func GetParam(c *gin.Context, key string) string {
	v := c.Query(key)
	if v == "" {
		v = c.PostForm(key)
	}
	if v == "" {
		v = c.Param(key)
	}
	if v == "" {
		v = c.GetHeader(key)
	}

	return v
}
func ReturnDataI(c *gin.Context, status int, v interface{}, callbackName string) []byte {
	object := DataIStruct{
		Status: status,
		Info:   kcode.GetCodeChnMsg(status),
		Data:   v,
	}
	return ReturnData(c, object, callbackName)
}
func ReturnData(c *gin.Context, v interface{}, callbackName string) []byte {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		kinit.LogError.Println(err)
	}
	if callbackName == "" {
		c.Data(http.StatusOK, "text/plain", jsonStr)
	} else {
		res := []byte(callbackName)
		res = append(res, []byte("(")...)
		res = append(res, jsonStr...)
		res = append(res, []byte(");")...)
		c.Data(http.StatusOK, "application/json; charset=utf-8", res)
	}
	return jsonStr
}
func ReturnDataStr(c *gin.Context, jsonStr []byte, callbackName string) []byte {
	if callbackName == "" {
		c.Data(http.StatusOK, "text/plain", jsonStr)
	} else {
		res := []byte(callbackName)
		res = append(res, []byte("(")...)
		res = append(res, jsonStr...)
		res = append(res, []byte(");")...)
		c.Data(http.StatusOK, "application/json; charset=utf-8", res)
	}
	return jsonStr
}

func SendErrorJsonStr(c *gin.Context, code int, innerException string, callbackName string) {
	jsonStr := GetErrorJsonStr(code, innerException)
	if callbackName == "" {
		c.Data(http.StatusOK, "text/plain", jsonStr)
	} else {
		res := []byte(callbackName)
		res = append(res, []byte("(")...)
		res = append(res, jsonStr...)
		res = append(res, []byte(");")...)
		c.Data(http.StatusOK, "application/json; charset=utf-8", res)
	}
}

func GetErrorJsonStr(code int, innerException string) []byte {

	str := ""
	if kcode.IS_TEST_SERVER == 1 {
		str = innerException
	}

	object := DataIStruct{
		Status:         code,
		Info:           kcode.GetCodeChnMsg(code),
		Data:           struct{}{},
		InnerException: str,
	}
	jsonStr, _ := json.Marshal(object)

	return jsonStr
}

func ReturnObjJson(c *gin.Context, v interface{}) []byte {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		kinit.LogError.Println(err)
	}

	c.Data(http.StatusOK, "text/plain", jsonStr)
	return jsonStr
}

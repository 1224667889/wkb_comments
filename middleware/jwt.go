package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Result struct {
	Status 	string 		`json:"status"`
	Msg	 	interface{} `json:"msg"`
}

//JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = 200
		tokenList := strings.Split(c.GetHeader("Authorization"), " ")
		if len(tokenList) < 2{
			c.JSON(code, gin.H{
				"status": code,
				"msg":  "无效token",
			})
			c.Abort()
			return
		}
		token := strings.Split(c.GetHeader("Authorization"), " ")[1]
		url := "**********************/api/id"
		r := Get(url, token)
		userId := 0
		if r.Status == "200" {
			userId = int(r.Msg.(float64))
		} else {
			c.JSON(code, gin.H{
				"status": code,
				"msg":  "身份验证失败",
			})
			c.Abort()
			return
		}
		fmt.Println(userId)
		c.Next()
	}
}

// Get 获取token信息
func Get(url string, token string) Result {
	resp, err := http.Get(url+"?token="+token)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}






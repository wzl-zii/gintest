package main

import "github.com/gin-gonic/gin"

type AAA struct {
	Title   string `json:"title"` //这种方法可以使显示出来的json小写
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "黄琴")
	})
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "hello world",
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "hello world!!!!",
		})
	})
	r.GET("/jsonp", func(c *gin.Context) { //jsonp，将callback值返回
		a := &AAA{
			Title:   "标题1--jsonp",
			Desc:    "描述一下1",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})
	r.Run()
}

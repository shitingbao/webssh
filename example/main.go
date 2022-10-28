package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shitingbao/webssh"
)

func main() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Println("Getwd err:", err)
		return
	}
	r := gin.Default()
	// 跨域设置
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		// AllowOrigins:     []string{"http://127.0.0.1", "http://localhost:8080"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// 将编译好的 dist 文件放在同级目录
	r.GET("/ws", ServeConn)
	r.Static("/static", path.Join(rootPath, "dist"))
	r.LoadHTMLFiles(path.Join(rootPath, "dist/index.html"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":8080")
}

func ServeConn(c *gin.Context) {
	opt := []webssh.Option{
		webssh.WithHostAddr("hostAddress"),
		webssh.WithUser("root"),
		webssh.WithKeyValue("yor PrivateKey"),
		webssh.WithTimeOut(time.Second)}
	webssh.SSHHandle(c.Writer, c.Request, opt...)
}

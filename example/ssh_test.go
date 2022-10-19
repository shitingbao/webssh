package example

import (
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shitingbao/webssh"
)

func TestSSh(t *testing.T) {
	r := gin.Default()
	//跨域设置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1", "http://localhost:8080"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/ws/:id", ServeConn)
	r.Static("/static", "/yourprojectpath/webssh/front/dist/")
	r.LoadHTMLFiles("/yourprojectpath/webssh/front/dist/index.html")
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

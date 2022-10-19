package example

import (
	"net/http"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shitingbao/webssh"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": err.Error()})
		return
	}
	defer wsConn.Close()
	webssh.SSHHandle(wsConn)
}

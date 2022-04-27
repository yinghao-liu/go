package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ws(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Printf("err is %v\n", err)
		return
	}
	defer ws.Close() //返回前关闭
	//for {
	//读取ws中的数据
	mt, message, err := ws.ReadMessage()
	if err != nil {
		//break
	}
	//写入ws数据
	err = ws.WriteMessage(mt, message)
	if err != nil {
		//break
	}
	//}
}
func main() {
	r := gin.Default()
	r.GET("/ws", ws)
	r.StaticFile("/", "index.html")
	r.Run(":8080")
}

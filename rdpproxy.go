package main

import (
	"encoding/json"
	"github.com/GoFeGroup/gordp/glog"
	"net/http"

	"github.com/GoFeGroup/gordp"
	"github.com/GoFeGroup/gordp/core"
	"github.com/GoFeGroup/gordp/proto/bitmap"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Processor struct {
	ws *websocket.Conn
}

func (p *Processor) ProcessBitmap(option *bitmap.Option, bitmap *bitmap.BitMap) {
	message := &Message{Bitmap: &Bitmap{
		X:    option.Left,
		Y:    option.Top,
		W:    option.Width,
		H:    option.Height,
		Data: bitmap.ToPng(),
	}}
	core.ThrowError(p.ws.WriteJSON(message))
}

func rdpProxy(ctx *gin.Context) {
	proto := ctx.Request.Header.Get("Sec-Websocket-Protocol")

	ws, err := upgrade.Upgrade(ctx.Writer, ctx.Request, http.Header{
		"Sec-Websocket-Protocol": {proto},
	})
	core.ThrowError(err)
	client := gordp.NewClient(&gordp.Option{
		Addr:     "192.168.1.129:3389",
		UserName: "anhk",
		Password: "[your-password-here]",
	})

	core.ThrowError(client.Connect())
	go func() {
		core.ThrowError(client.Run(&Processor{ws: ws}))
	}()
	for {
		_, msgBytes, err := ws.ReadMessage()
		core.ThrowError(err)

		msg := Message{}
		core.ThrowError(json.Unmarshal(msgBytes, &msg))
		glog.Debugf("msg: %+v", msg.Mouse)

		if msg.Keyboard != nil {

		}

		if msg.Mouse != nil {
			switch msg.Mouse.Type {
			case "mousemove":
				//client.SendMouseEvent()
			case "mousedown":
			case "mouseup":
			case "mouseright":
			}
		}

	}
}

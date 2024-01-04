package events

import (
	"io"
	"log"

	alf "github.com/PiterWeb/Alf-Router"
	"github.com/fasthttp/websocket"
	"github.com/pterm/pterm"
	"github.com/valyala/fasthttp"
)

type Event_plugin struct {
	ReadBufferSize  int
	WriteBufferSize int
	Port            string
}

func (plugin Event_plugin) Init_plugin(appConfig *alf.AppConfig) error {

	if plugin.Port == "" {
		plugin.Port = "8080"
	}

	if plugin.ReadBufferSize == 0 {
		plugin.ReadBufferSize = 1024
	}

	if plugin.WriteBufferSize == 0 {
		plugin.WriteBufferSize = 1024
	}

	upgrader := websocket.FastHTTPUpgrader{
		ReadBufferSize:  int(plugin.ReadBufferSize),
		WriteBufferSize: int(plugin.WriteBufferSize),
	}

	pterm.Info.Println("Event Server running on  port :" + plugin.Port)

	go func() {
		fasthttp.ListenAndServe(":"+plugin.Port, func(ctx *fasthttp.RequestCtx) {

			err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {

				for {
					_, message, err := conn.ReadMessage()
					if err != nil || err == io.EOF {
						log.Fatal("Error reading: ", err)
						break
					}

					log.Printf("recv: %s", message)
				}

			})

			if err != nil {
				return
			}

		})
	}()

	return nil

}

func readLoop(c *websocket.Conn) {

	for {

		var message interface{}

		c.ReadJSON(&message)

		pterm.Info.Print(message)

	}

}

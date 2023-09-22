package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

type chatroom struct {
	header
}

func Chatroom(c echo.Context) error {
	return c.Render(http.StatusOK, "chat.html", chatroom{
		header{
			Title: "Chatroom",
		},
	})
}

var groupCount int = -1

func ChatWS(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		//Write
		groupCount++
		err := ws.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(groupCount)))
		if err != nil {
			c.Logger().Error(err)
		}

		//Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}

		fmt.Printf("%s\n", msg)
	}
}

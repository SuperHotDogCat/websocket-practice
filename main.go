package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil) // 詳しくはわからないが, 既存のhttp通信をwebsocketにupgradeするもの
	if err != nil {
		return err
	}
	defer ws.Close()
	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)

		err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

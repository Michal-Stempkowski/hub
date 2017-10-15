package main

import (
	"hub/logic_designer"
	"hub/server"
)

func main() {
	urlHandlerMap := server.MergeHandlerMaps(
		logic_designer.GetHandlers())
	server.StartServer(urlHandlerMap)
}

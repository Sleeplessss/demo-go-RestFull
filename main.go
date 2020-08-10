package main

import ( 
	"restfull/handler"
	"restfull/configs"
)

func main() {

	configs.SetENV()
	router := handler.Router()

	router.Run(":3000") // Run server
}

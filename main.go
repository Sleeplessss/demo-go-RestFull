package main

import ( "restfull/handler" )

func main() {

	router := handler.Router()

	router.Run(":3000") // Run server
}

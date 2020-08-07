package main

import ( "restfull/routers/handlers" )

func main() {

	router := handler.Router()

	router.Run(":3000") // Run server
}

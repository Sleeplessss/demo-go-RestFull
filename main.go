package main

import ( "restfull/router" )

func main() {

	router := router.Router()

	router.Run(":3000") // Run server
}

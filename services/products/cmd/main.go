package main

import (
	"log"

	"git.mip-consult.de/sde/suzuki-framework/microservices/services/products"
)

const port = ":8888"

func main() {
	log.Printf("Listening on port %s\n", port)
	products.ListenGRPC()
}

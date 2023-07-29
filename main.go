package main

import (
	"fmt"
	//"log"

	//"os"
	//"yuvraj/controllers"
	//"yuvraj/config"
	"yuvraj/router"

	_ "github.com/lib/pq"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8000")
	fmt.Println("Connected")
}

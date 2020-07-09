package main

import (
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/models"

	//greeter "github.com/protoBuf"
	"github.com/routers"
)


func main() {
	r := routers.InitRouterForClient()
	r.Run("0.0.0.0:8081")
}

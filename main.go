package main

import (
	userservice "github.com/oneVegeDog/sourceNavigator/pkg/rpc/userService/userservice"
	"log"
)

func main() {
	svr := userservice.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

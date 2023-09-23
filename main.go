package main

import (
	sourcenavigatorservice "github.com/oneVegeDog/sourceNavigator/pkg/rpc/sourceNavigatorService/sourcenavigatorservice"
	"log"
)

func main() {
	svr := sourcenavigatorservice.NewServer(new(SourceNavigatorServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

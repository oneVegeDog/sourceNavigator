package main

import (
	source_rpc "github.com/oneVegeDog/sourceNavigator/pkg/source_rpc/sourcenavigatorservice"
	"log"
)

func main() {
	svr := source_rpc.NewServer(new(SourceNavigatorServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

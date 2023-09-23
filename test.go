package main

import "github.com/oneVegeGog/sourceNavigator/internal"

func main() {
	println(internal.GLOBAL_CONFIG.GetString("sourceNavigator.mysql.dns"))
}

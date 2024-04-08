package main

import (
	"fmt"
	"main/setting"
)

func main() {
	fmt.Println(setting.Config.PgPort)
}

package main

import (
	"fmt"

	"github.com/chaimakr/book_management_system/core/getter/config"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)
}

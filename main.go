package main

import (
	"fmt"
	"tools/config"
)

func main() {
	config.InitConfig()
	r := setupRouter()
	if err := r.Run(fmt.Sprintf(":%s", config.Settings.App.Port)); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

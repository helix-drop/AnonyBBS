package main

import (
	"fmt"
	"github.com/helix-drop/AnonyBBS/config"
	"github.com/helix-drop/AnonyBBS/routes"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	r := routes.NewRouter()

	addr := viper.GetString("gin.address")
	port := viper.GetString("gin.port")
	if err := r.Run(addr + ":" + port); err != nil {
		panic(fmt.Sprintf("gin启动失败:%s", err))
	}

}

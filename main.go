package main

import (
	"flag"
	"fmt"
	"github.com/corinnekrych/grapnel/config"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "Path to the config file to read")
	flag.Parse()
	fmt.Printf("Hello %s\n", configFilePath)
	configData, err := config.Read(configFilePath)
	if err != nil {
		fmt.Errorf("Can not read config file: %v error: %v", configFilePath, err)
		os.Exit(0)
	}
	fmt.Printf("ConfigData is: %#v", configData)
	app := gin.Default()
	v1 := app.Group("api/v1")
	{
		v1.POST("hooks", PostHook)
	}
	app.Run(configData.Address)
}

func PostHook(c *gin.Context) {

}
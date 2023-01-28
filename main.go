package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Datasource struct {
		Url string `yaml:"url"`
	} `yaml:"datasource"`
}

func readConfig() {
	config, err := os.Open("./config/config.yml")

	if err != nil {
		log.Println(err)
	}
	defer config.Close()
	var appConfig *Config
	decoder := yaml.NewDecoder(config)
	errors := decoder.Decode(&appConfig)

	if errors != nil {
		log.Println(errors)
	}
	fmt.Println("url: ", appConfig.Datasource.Url)
	fmt.Println("Hi")
}

func main() {
	readConfig()
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

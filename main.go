package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	// read config
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v\n", err)
	}

	// Get API token and refresh token
	tkn := viper.GetString("api_token")
	rfsh := viper.GetString("refresh_token")

	// start server
	fmt.Println("Hello")

}

func DoRequest(token string) error {

}

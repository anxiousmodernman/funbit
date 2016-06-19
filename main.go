package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
)

func printConfigFile() {
	fmt.Println(`
# Specify your settings from dev.fitbit.com here
server:
    client_id:
    secret:
    redirect_uri:

`)

}

func main() {

	if len(os.Args) > 1 {
		// gotta do this or we'll PANIC
		if os.Args[1] == "printConfig" {
			printConfigFile()
			os.Exit(0)
		}
	}

	// read config
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v\n", err)
	}

	var svr Server
	svr.ClientID = viper.GetString("server.client_id")
	svr.Secret = viper.GetString("server.secret")
	svr.RedirectURI = viper.GetString("server.redirect_uri")

	db, err := bolt.Open("funbit.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal("WTF", err)
	}
	svr.DB = db

	//go populateData(db)

	addr := "0.0.0.0:42069"
	log.Println("Starting server on", addr)
	log.Println("Server data:", svr)

	http.ListenAndServe(addr, &svr)
}

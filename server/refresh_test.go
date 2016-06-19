package server_test

import (
	"log"
	"testing"
	"time"

	"github.com/anxiousmodernman/funbit/server"
	"github.com/boltdb/bolt"
)

func TestGetRefreshToken(t *testing.T) {

	userID := "4NZ8B3"

	var svr server.Server
	db, err := bolt.Open("../funbit.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal("WTF", err)
	}
	svr.DB = db

	rt, err := svr.GetRefreshToken(userID)
	if err != nil {
		t.Errorf("Something bad happened: %v\n", err)
	}

	expected := "5801f6e29a30c15857d864185929d2668507a9092a02ea1323d95768c6a50c68"

	if rt != expected {
		t.Errorf("Got rt %s, but expected %s\n", rt)
	}
}

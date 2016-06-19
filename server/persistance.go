package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/anxiousmodernman/funbit/client"
	"github.com/boltdb/bolt"
)

// LoadDatabaseFixtures populates test databases.
func LoadDatabaseFixtures(fixturePath, databasePath string) error {

	var data []client.AuthResponse
	fxt := getFileContentsOrFatal(fixturePath)
	err := json.Unmarshal(fxt, &data)
	if err != nil {
		log.Fatal("Error loading database fixture:", err)
	}

	db, _ := bolt.Open(databasePath, 0600, nil)

	db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte("tokens"))
		if err != nil {
			return err
		}

		for _, item := range data {

			converted, err := json.Marshal(item)
			if err != nil {
				log.Println("Could not marshal AuthResponse for test fixture", err)
				return err
			}

			err = b.Put([]byte(item.UserID), converted)
			if err != nil {
				log.Println("Could not put AuthResponse for test fixture", err)
				return err
			}

		}

		return nil
	})

	return nil

}

// getFileContentsOrFatal will crash our program if file on disk can't be opened.
func getFileContentsOrFatal(path string) []byte {

	var contents []byte

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	return contents
}

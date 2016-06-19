package server

import (
	"encoding/json"

	"github.com/anxiousmodernman/funbit/client"
	"github.com/boltdb/bolt"
)

func (svr *Server) GetRefreshToken(userID string) (string, error) {

	var authData client.AuthResponse

	svr.DB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("tokens"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(k) == userID {
				err := json.Unmarshal(v, &authData)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	return authData.RefreshToken, nil
}

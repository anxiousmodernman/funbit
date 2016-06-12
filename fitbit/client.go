package fitbit

import "encoding/base64"

func FetchData(token string) error {
	// TODO use token to get that data, yo.
	// GET /1/user/[user-id]/body/[resource-path]/date/[base-date]/[end-date].json

	// return nil for error because everything's fine
	return nil
}

// NewAuthorizationHeader returns the base64 encoding of "clientID:secret".
func NewAuthorizationHeader(clientID, secret string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+secret))
}

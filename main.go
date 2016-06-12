package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"golang.org/x/net/context"

	"github.com/anxiousmodernman/funbit/fitbit"
	"github.com/spf13/viper"
)

func main() {

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

	addr := "0.0.0.0:42069"
	log.Println("Starting server on", addr)
	log.Println("Server data:", svr)

	http.ListenAndServe(addr, &svr)
}

// Server is our Handler
type Server struct {
	ClientID    string
	RedirectURI string
	Secret      string
}

// Use these constants for keys into the context.Context object.
const (
	AuthHdr = iota
)

// ServeHTTP lets Server satisfy the http.Handler interface.
func (svr *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// initialize a context
	ctx := context.Background()
	// Set the computed auth header on the context. Pass it to all handlers.
	ctx = context.WithValue(ctx, AuthHdr, fitbit.NewAuthorizationHeader(svr.ClientID, svr.Secret))

	// TODO use framework
	switch r.URL.Path {
	case "/auth":
		svr.Auth(ctx, w, r)
	default:
		Reply404(ctx, w, r)
	}

}

func AuthHdrFromContext(ctx context.Context) string {
	// we use the const key to nab that specific value, mothafucka!
	val := ctx.Value(AuthHdr)
	// val comes back as untyped interface{}, cast to string
	if s, ok := val.(string); !ok {
		log.Println("Warning: expected AuthHdr to be string")
		return ""
	} else {
		return s
	}
}

func (svr *Server) Auth(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	debugRequest(r)

	// get "code" off the url param
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Expected URL parameter \"code\"", 400)
		return
	}

	// Perform http request according to these docs: https://dev.fitbit.com/docs/oauth2/#access-token-request
	form := url.Values{}
	form.Add("clientId", svr.ClientID)
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", svr.RedirectURI)

	req, err := http.NewRequest("POST", "https://api.fitbit.com/oauth2/token", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal("Your server is whack, man.")
	}

	authHeader := AuthHdrFromContext(ctx)
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	debugRequest(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error requesting token: %v\n", err)
		http.Error(w, "Error requesting token", 500)
		return
	}

	debugResponse(resp)

	// Could panic if body is nil
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading JSON body: %v\n", err)
		// Error is a helper func that writes to the ResponseWriter
		http.Error(w, "Error reading JSON body", 500)
		// You MUST return early from HTTP handlers
		return
	}

	var auth fitbit.AuthResponse

	err = json.Unmarshal(contents, &auth)
	if err != nil {
		http.Error(w, "Error reading JSON body", 500)
		return
	}

	log.Println("Got this data")
	log.Println(auth)

}

func Reply404(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// helper funcs...
func debugResponse(resp *http.Response) {
	dump, _ := httputil.DumpResponse(resp, true)
	fmt.Println("DEBUG:")
	fmt.Printf("%q", dump)
}
func debugRequest(req *http.Request) {
	dump, _ := httputil.DumpRequest(req, true)
	fmt.Println("DEBUG:")
	fmt.Printf("%q", dump)
}

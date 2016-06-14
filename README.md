# funbit

Personal FitBit servers

## Goals

* Compute weight over time from Aria wifi scale
* ???


## Setup

Get the code and build it.

```
go get github.com/anxiousmodernman/funbit
cd $GOPATH/github.com/anxiousmodernman/funbit
go get ./...
go build
```

Run ngrok on port **42069**, and take note of your new url.

```
ngrok http 42069
```

funbit requires a conf.yml; Template it out with this built in command.

```
./funbit printConfig > conf.yml
```

Edit the values in conf.yml to match your application settings in dev.fitbit.com

```yaml
# conf.yml
server: 
    # From the "Manage My Apps" menu, find your app and fill in this info
    client_id: 123456
    secret: your-secret
    redirect_uri: https://5a5a9488.ngrok.io/auth  # must match exactly
```

Start the server

```
./funbit
```

Now you should be able to hit the **/oauth2/authorize** authorization URL 
constructed by the [interactive tutorial](https://dev.fitbit.com/apps/oauthinteractivetutorial),
and see debug output in the terminal.


## To Do

* Store Data in boltDB
* Poll/de-dupe data 
* Construct the authorization URL on some **/login** page
* Run on Digital Ocean



# go-auth
user authorization verification module

[![Coverage Status](https://coveralls.io/repos/github/deweppro/go-auth/badge.svg?branch=master)](https://coveralls.io/github/deweppro/go-auth?branch=master)
[![Release](https://img.shields.io/github/release/deweppro/go-auth.svg?style=flat-square)](https://github.com/deweppro/go-auth/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/go-auth)](https://goreportcard.com/report/github.com/deweppro/go-auth)
[![CI](https://github.com/deweppro/go-auth/actions/workflows/ci.yml/badge.svg)](https://github.com/deweppro/go-auth/actions/workflows/ci.yml)

## Install

```sh
go get -u github.com/deweppro/go-auth
```

## Examples

see more [here](internal/examples)


### Setup providers config

```go
import (
    "github.com/deweppro/go-auth/config"
    "github.com/deweppro/go-auth/providers"
)

var providerConfig = &config.Config{
		Provider: []config.ConfigItem{
			{
				Code:         "google",
				ClientID:     "****************.apps.googleusercontent.com",
				ClientSecret: "****************",
				RedirectURL:  "https://example.com/oauth/callback/google",
			},
		},
	}

providers := providers.New(providerConfig)
```

You can add our provider corresponding to the `providers.IProvider` interface

```go
providers.Add(provider1, provider2, ...)
```

### Initializing the authorization service

```go
import "github.com/deweppro/go-auth"

authServ := auth.New(providers)
```

### Request Handlers

the methods return an http handler `func(http.ResponseWriter, *http.Request)` when called 

#### Creating a redirect to the provider

```go
// <provider name> - provider name from the config
authServ.Request(<provider name>)
```

#### Processing the provider's response

```go
authServ.CallBack(<provider name>, <callback function>)
```
* `<provider name>` - provider name from config
* `<callback function>` - response results handler has signature `func(http.ResponseWriter, *http.Request, isp.IUser)`
  * `isp.IUser` - a new instance of the processed user model with ACL data filling


## License

BSD-3-Clause License. See the LICENSE file for details.
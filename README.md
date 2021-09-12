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
	"github.com/deweppro/go-auth/provider"
	"github.com/deweppro/go-auth/provider/isp"
)

var providerConfig = &provider.Config{
		Provider: []isp.Config{
			{
				Name:         "google",
				ClientID:     "****************.apps.googleusercontent.com",
				ClientSecret: "****************",
				RedirectURL:  "https://example.com/oauth/callback/google",
			},
		},
	}

providers := provider.New(providerConfig)
```

You can add our provider corresponding to the `provider.Provider` interface

```go
providers.Add(provider1, provider2, ...)
```

### You can use ACL storage based on the config

```go
import "github.com/deweppro/go-auth/storage"

var storageConfig = &storage.Config{
		ACL: map[string]string{
			"example.user@gmail.com": "01010101010",
		},
	}

store := storage.NewConfigStorage(storageConfig)
```

You can use your provider corresponding to the `storage.IStorage` interface

### Initializing the authorization service

```go
import "github.com/deweppro/go-auth"

authServ := auth.New(store, providers)
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
* `<callback function>` - response results handler has signature `func([]byte, http.ResponseWriter)`
  * `[]byte` - raw result of the response from the provider is transmitted in the form of a byte code
  * `http.ResponseWriter` - writer who accepts your processing of response results from provider

```go
authServ.CallBackWithACL(<provider name>, <user model>, <callback function>)
```
* `<provider name>` - provider name from config
* `<user model>` - user model corresponding to the interface `auth.IUser`
* `<callback function>` - response results handler has signature `func(auth.IUser, http.ResponseWriter)`
  * `auth.IUser` - a new instance of the processed user model with ACL data filling
  * `http.ResponseWriter` - writer who accepts your processing of response results from provider

## License

BSD-3-Clause License. See the LICENSE file for details.
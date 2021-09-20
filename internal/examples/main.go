package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deweppro/go-auth"
	"github.com/deweppro/go-auth/acl"
	"github.com/deweppro/go-auth/provider"
	"github.com/deweppro/go-auth/provider/isp"
	"github.com/deweppro/go-auth/storage"
	"github.com/deweppro/go-http/web/routes"
	"github.com/deweppro/go-http/web/server"
	"github.com/deweppro/go-logger"
)

var (
	provConf = &provider.Config{
		Provider: []isp.Config{
			{
				Name:         "google",
				ClientID:     "****************.apps.googleusercontent.com",
				ClientSecret: "****************",
				RedirectURL:  "https://example.com/oauth/callback/google",
			},
		},
	}

	storeConf = &storage.Config{
		ACL: map[string]string{
			"example.user@gmail.com": "01010101010",
		},
	}

	aclHub = acl.New(storage.NewConfigStorage(storeConf), 10)

	servConf = &server.Config{HTTP: server.ConfigItem{Addr: ":8080"}}
)

func main() {
	authServ := auth.New(provider.New(provConf))

	route := routes.NewRouter()
	route.Route("/oauth/request/google", routes.CtrlFunc(authServ.Request("google")), http.MethodGet)
	route.Route("/oauth/callback/google", routes.CtrlFunc(authServ.CallBackWithUser("google", &isp.UserGoogle{}, userHandler)), http.MethodGet)

	serv := server.New(servConf, route, logger.Default())
	serv.Up()

	<-time.After(60 * time.Minute)
}

const out = `
email: %s
name:  %s
ico:   %s
acl:   %+v
`

func userHandler(i isp.IUser, w http.ResponseWriter) {
	w.WriteHeader(200)
	levels := aclHub.GetAll(i.GetEmail())
	w.Write([]byte(fmt.Sprintf(out, i.GetEmail(), i.GetName(), i.GetIcon(), levels)))
}

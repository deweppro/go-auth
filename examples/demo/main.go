package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deweppro/go-auth"
	"github.com/deweppro/go-auth/acl"
	"github.com/deweppro/go-auth/config"
	"github.com/deweppro/go-auth/providers"
	"github.com/deweppro/go-auth/providers/isp"
	"github.com/deweppro/go-auth/providers/isp/google"
	"github.com/deweppro/go-auth/storage/inconfig"
	"github.com/deweppro/go-http/pkg/routes"
	"github.com/deweppro/go-http/servers"
	"github.com/deweppro/go-http/servers/web"
	"github.com/deweppro/go-logger"
)

var (
	provConf = &config.Config{
		Provider: []config.ConfigItem{
			{
				Code:         "google",
				ClientID:     "****************.apps.googleusercontent.com",
				ClientSecret: "****************",
				RedirectURL:  "https://example.com/oauth/callback/google",
			},
		},
	}

	storeConf = &inconfig.Config{
		ACL: map[string]string{
			"example.user@gmail.com": "01010101010",
		},
	}

	aclHub = acl.New(inconfig.New(storeConf), 10)

	servConf = servers.Config{Addr: ":8080"}
)

func main() {
	authServ := auth.New(providers.New(provConf))

	route := routes.NewRouter()
	route.Route("/oauth/request/google", routes.CtrlFunc(authServ.Request(google.CODE)), http.MethodGet)
	route.Route("/oauth/callback/google", routes.CtrlFunc(authServ.CallBack(google.CODE, userHandler)), http.MethodGet)

	serv := web.New(servConf, route, logger.Default())
	serv.Up()
	<-time.After(60 * time.Minute)
	serv.Down()
}

const out = `
email: %s
name:  %s
ico:   %s
acl:   %+v
`

func userHandler(w http.ResponseWriter, _ *http.Request, u isp.IUser) {
	w.WriteHeader(200)
	levels, _ := aclHub.GetAll(u.GetEmail())
	w.Write([]byte(fmt.Sprintf(out, u.GetEmail(), u.GetName(), u.GetIcon(), levels)))
}

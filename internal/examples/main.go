package main

import (
	"net/http"
	"time"

	"github.com/deweppro/go-auth"
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

	servConf = &server.Config{HTTP: server.ConfigItem{Addr: ":8080"}}
)

var _ auth.IUser = (*User)(nil)

type User struct {
	Email string `json:"email"`
	ACL   string
}

func (v *User) GetEmail() string  { return v.Email }
func (v *User) GetACL() string    { return v.ACL }
func (v *User) SetACL(acl string) { v.ACL = acl }

func main() {
	authServ := auth.New(storage.NewConfigStorage(storeConf), provider.New(provConf))

	route := routes.NewRouter()
	route.Route("/oauth/request/google", routes.CtrlFunc(authServ.Request("google")), http.MethodGet)
	route.Route("/oauth/callback/google", routes.CtrlFunc(authServ.CallBackWithACL("google", &User{}, userHandler)), http.MethodGet)

	serv := server.New(servConf, route, logger.Default())
	serv.Up()

	<-time.After(60 * time.Minute)
}

func userHandler(i auth.IUser, w http.ResponseWriter) {
	u := i.(*User)
	w.WriteHeader(200)
	w.Write([]byte(u.Email + " => " + u.ACL))
}

package main

import (
	"fmt"
	"io/ioutil"
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
	"gopkg.in/yaml.v2"
)

var (
	servConf = &server.Config{HTTP: server.ConfigItem{Addr: ":8080"}}
	aclHub   = acl.New(NewDebugStorage(), 10)
	raw      = false
)

func main() {
	b, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	provConf := &provider.Config{}
	err = yaml.Unmarshal(b, provConf)
	if err != nil {
		panic(err)
	}

	authServ := auth.New(provider.New(provConf))

	route := routes.NewRouter()
	route.Route("/oauth/r/google", routes.CtrlFunc(authServ.Request("google")), http.MethodGet)
	route.Route("/oauth/c/google", routes.CtrlFunc(switchHandler(raw, "google", &isp.UserGoogle{}, authServ)), http.MethodGet)

	route.Route("/oauth/r/yandex", routes.CtrlFunc(authServ.Request("yandex")), http.MethodGet)
	route.Route("/oauth/c/yandex", routes.CtrlFunc(switchHandler(raw, "yandex", &isp.UserYandex{}, authServ)), http.MethodGet)

	serv := server.New(servConf, route, logger.Default())
	serv.Up()

	<-time.After(60 * time.Minute)
}

func switchHandler(raw bool, name string, m isp.IUser, a *auth.Auth) auth.HttpHandler {
	if raw {
		return a.CallBack(name, userHandlerRaw)
	}
	return a.CallBackWithUser(name, m, userHandlerACL)
}

func userHandlerRaw(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

const out = `
email: %s
name:  %s
ico:   %s
acl:   %+v
`

func userHandlerACL(u isp.IUser, w http.ResponseWriter) {
	w.WriteHeader(200)
	levels := aclHub.GetAll(u.GetEmail())
	w.Write([]byte(fmt.Sprintf(out, u.GetEmail(), u.GetName(), u.GetIcon(), levels)))
}

type DebugStorage struct{}

func NewDebugStorage() storage.IStorage {
	return &DebugStorage{}
}

func (v *DebugStorage) FindACL(email string) (string, bool) {
	return "000", true
}

func (v *DebugStorage) ChangeACL(_, _ string) error {
	return nil
}

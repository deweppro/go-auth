package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/deweppro/go-auth"
	"github.com/deweppro/go-auth/acl"
	"github.com/deweppro/go-auth/config"
	"github.com/deweppro/go-auth/providers"
	"github.com/deweppro/go-auth/providers/isp"
	"github.com/deweppro/go-auth/providers/isp/google"
	"github.com/deweppro/go-auth/providers/isp/yandex"
	"github.com/deweppro/go-auth/storage"
	"github.com/deweppro/go-http/pkg/httputil/enc"
	"github.com/deweppro/go-http/pkg/routes"
	"github.com/deweppro/go-http/servers"
	"github.com/deweppro/go-http/servers/web"
	"github.com/deweppro/go-logger"
	"gopkg.in/yaml.v2"
)

var (
	servConf = servers.Config{Addr: ":8080"}
	aclHub   = acl.New(NewDebugStorage(), 10)
)

func main() {
	b, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	provConf := &config.Config{}
	err = yaml.Unmarshal(b, provConf)
	if err != nil {
		panic(err)
	}

	authServ := auth.New(providers.New(provConf))

	route := routes.NewRouter()
	route.Route("/oauth/r/google", routes.CtrlFunc(authServ.Request(google.CODE)), http.MethodGet)
	route.Route("/oauth/c/google", routes.CtrlFunc(switchHandler(google.CODE, authServ)), http.MethodGet)

	route.Route("/oauth/r/yandex", routes.CtrlFunc(authServ.Request(yandex.CODE)), http.MethodGet)
	route.Route("/oauth/c/yandex", routes.CtrlFunc(switchHandler(yandex.CODE, authServ)), http.MethodGet)

	serv := web.New(servConf, route, logger.Default())
	serv.Up()
	<-time.After(60 * time.Minute)
	serv.Down()
}

func switchHandler(name string, a *auth.Auth) auth.HttpHandler {
	return a.CallBack(name, userHandlerACL)
}

const out = `
email: %s
name:  %s
ico:   %s
acl:   %+v
`

func userHandlerACL(w http.ResponseWriter, _ *http.Request, u isp.IUser) {
	w.WriteHeader(200)
	levels, err := aclHub.GetAll(u.GetEmail())
	if err != nil {
		enc.Error(w, err)
		return
	}
	enc.Raw(w, []byte(fmt.Sprintf(out, u.GetEmail(), u.GetName(), u.GetIcon(), levels)))
}

type DebugStorage struct{}

func NewDebugStorage() storage.IStorage {
	return &DebugStorage{}
}

func (v *DebugStorage) FindACL(_ string) (string, error) {
	return "000", nil
}

func (v *DebugStorage) ChangeACL(_, _ string) error {
	return nil
}

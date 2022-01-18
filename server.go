package webhookForward

import (
	"net/http"

	"github.com/go-sdk/lib/app"
	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/log"
	"github.com/go-sdk/lib/srv"
)

var (
	handler = srv.New()

	addr = conf.Get("addr").StringD(":9988")
)

func init() {
	handler.Use(srv.Logger(), srv.Recovery())

	handler.GET("/_health", _health)
	handler.GET("/send", Send)
}

func _health(c *srv.Context) {
	c.JSON(http.StatusOK, app.VersionInfoMap())
}

func Start() error {
	go func() { srv.PrintRoutes(handler) }()

	log.Infof("listening and serving HTTP on %s", addr)
	return handler.Run(addr)
}

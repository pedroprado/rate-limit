package presentation

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ServerHttp interface {
	StartServer(ctx context.Context, port string)
	GetGinRouterGroup(relativePath string) *gin.RouterGroup
}

type serverHttpGin struct {
	ServerHttp
	ginEngine *gin.Engine
}

func (ref *serverHttpGin) StartServer(ctx context.Context, port string) {
	startServer(ctx, ref.ginEngine, port)
}

func startServer(ctx context.Context, handler http.Handler, port string) {

	address := fmt.Sprintf(":%s", port)

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	srvErrs := make(chan error, 1)
	go func() {
		srvErrs <- server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	shutdown := gracefulShutdown(server)

	select {
	case err := <-srvErrs:
		shutdown(err)
	case sig := <-quit:
		shutdown(sig)
	}

	logrus.Info(ctx, "Server shutdown successfully")
}

func NewServerHttpGin(pretty bool) ServerHttp {
	if !pretty {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	return &serverHttpGin{
		ginEngine: r,
	}
}

func (ref *serverHttpGin) GetGinRouterGroup(relativePath string) *gin.RouterGroup {
	ref.ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-LimitType"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	return ref.ginEngine.Group(relativePath)
}

func RegisterInfraApi(ginRouterGroup *gin.RouterGroup, diagnosticMode bool) {
	if diagnosticMode {
		pprof.RouteRegister(ginRouterGroup, "pprof")
	}
	ginRouterGroup.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
}

func gracefulShutdown(srv *http.Server) func(reason interface{}) {
	return func(reason interface{}) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		logrus.WithField("reason", reason).Info(ctx, "Server shutting down")

		if err := srv.Shutdown(ctx); err != nil {
			logrus.Error(ctx, "Error Gracefully Shutting Down API", err)
			return
		}
	}
}

package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "net/http/pprof"

	"github.com/SergeyDavidenko/subscription/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var (
	logGin = log.New()
	// Quit gc shutdown
	Quit = make(chan os.Signal, 1)
)

func apiRourter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.Use(authMiddleware())
		v1.POST("/subscriptions", createSubscriptions)
		v1.DELETE("/subscriptions", deleteSubscriptions)
		v1.PUT("/subscriptions", updateSubscriptions)
		v1.GET("/subscriptions", allSubscriptions)
		v1.GET("/info", info)
	}
}

func promRouter(r *gin.Engine) *ginprometheus.Prometheus {
	prom := ginprometheus.NewPrometheus("gin")
	prom.MetricsPath = config.Conf.API.MetricURI
	prom.SetMetricsPath(r)

	r.GET(config.Conf.API.HealthURI, health)
	return prom
}

// WEBServerRun run
func WEBServerRun() {
	// Set release mode
	gin.SetMode(gin.ReleaseMode)
	// Set log level
	logGin.SetLevel(config.LogLevel)
	router := gin.New()
	router2 := gin.New()

	router.Use(promRouter(router2).HandlerFunc())
	router.Use(ginlogrus.Logger(logGin), gin.Recovery())
	apiRourter(router)

	addressAPI := fmt.Sprintf(":%s", strconv.Itoa(config.Conf.API.Port))
	srv := &http.Server{
		Addr:              addressAPI,
		Handler:           router,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	log.Info("api port ", addressAPI)
	go func() {
		addressHealth := fmt.Sprintf(":%s", strconv.Itoa(config.Conf.API.HealthPort))
		log.Info("health port ", addressHealth)
		router2.Run(addressHealth)
	}()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ", err)
		}
	}()
	go func() {
		var pprofAddress string
		if config.Conf.API.PProfPort == 0 {
			pprofAddress = "localhost:6060"
		} else {
			pprofAddress = fmt.Sprintf(":%s", strconv.Itoa(config.Conf.API.PProfPort))
		}
		log.Info("pprof port ", pprofAddress)
		log.Fatal(http.ListenAndServe(pprofAddress, nil))
	}()
	<-Quit
	log.Info("Server shutdown ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	log.Info("Server exiting")
}

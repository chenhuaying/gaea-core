package ginserver

import (
	"context"
	"log"
	"net/http"

	//
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gaeanetwork/gaea-core/api"
	"github.com/gaeanetwork/gaea-core/smartcontract/fabric"
	"github.com/gaeanetwork/gaea-core/smartcontract/factory"
	"github.com/gaeanetwork/gaea-core/tee/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Start gin web http
func Start() {
	r := setupRouter()
	go factory.InitSmartContractService(&fabric.Chaincode{})
	go server.NewTeeServer(server.GRPCAddr).Start()
	gracefulStart(r)
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Serving static files
	r.Static("/site", "/dist")

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Allows all origins
	r.Use(cors.Default())

	v1 := r.Group("/api")
	{
		// Ping test - curl http://localhost:8080/api/ping
		v1.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

		// Transmission api register
		files := v1.Group("/files")
		api.RegisterAPI(files)
	}
	return r
}

func gracefulStart(router *gin.Engine) {
	srv := &http.Server{
		Addr:    server.ListenAddr,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Start profiling http endpoint if enabled
	if server.ProfileEnabled {
		go func() {
			log.Printf("Starting profiling server with listenAddress = %s\n", server.PProfAddr)
			if profileErr := http.ListenAndServe(server.PProfAddr, nil); profileErr != nil {
				log.Panicf("Error starting profiler: %v\n", profileErr)
			}
		}()
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

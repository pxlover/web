package initialize

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webProject/model"
	"webProject/router"
)

func Server() {
	NewServer(routers)
}

func routers(route *gin.Engine) {
	route.LoadHTMLGlob("/Users/noone/Documents/work/index.html")
	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	router.InitCommonRouter(route)
}

func NewServer(router func(*gin.Engine)) {
	var engine = gin.New()
	engine.NoMethod(model.HandleNotFound)
	engine.NoRoute(model.HandleNotFound)
	engine.Use(model.Recover, model.Cors)
	pprof.Register(engine)

	router(engine)

	address := fmt.Sprintf(":%d", 80)
	svc := initServer(address, engine)
	fmt.Printf("****欢迎使用****\n****运行地址:http://127.0.0.1%s\n", address)

	go func() {
		svc.ListenAndServe()
	}()

	shutDownWait(svc)
}

func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr: 			address,
		Handler: 		router,
		ReadTimeout: 	10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func shutDownWait(svc *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := svc.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
	fmt.Println("Server exiting ...")
}
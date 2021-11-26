package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"study_event_go/cmd/container"
	"study_event_go/conf"
	"study_event_go/router"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	banner = "\n" +
		"                                        ##\n" +
		"                                        ##\n" +
		"               #                        ##                                                                  #\n" +
		"              ##                        ##                            ##                                   ##\n" +
		"              ##                        ##                            ##                                   ##\n" +
		"   /###     ######## ##   ####      ### ##  ##   ####            /##   ##    ###      /##  ###  /###     ########        /###      /###\n" +
		"  / #### / ########   ##    ###  / ######### ##    ###  /       / ###   ##    ###    / ###  ###/ #### / ########        /  ###  / / ###  /\n" +
		" ##  ###/     ##      ##     ###/ ##   ####  ##     ###/       /   ###  ##     ###  /   ###  ##   ###/     ##          /    ###/ /   ###/\n" +
		"####          ##      ##      ##  ##    ##   ##      ##       ##    ### ##      ## ##    ### ##    ##      ##         ##     ## ##    ##\n" +
		"  ###         ##      ##      ##  ##    ##   ##      ##       ########  ##      ## ########  ##    ##      ##         ##     ## ##    ##\n" +
		"    ###       ##      ##      ##  ##    ##   ##      ##       #######   ##      ## #######   ##    ##      ##         ##     ## ##    ##\n" +
		"      ###     ##      ##      ##  ##    ##   ##      ##       ##        ##      ## ##        ##    ##      ##         ##     ## ##    ##\n" +
		" /###  ##     ##      ##      /#  ##    /#   ##      ##       ####    / ##      /  ####    / ##    ##      ##         ##     ## ##    ##\n" +
		"/ #### /      ##       ######/ ##  ####/      #########        ######/   ######/    ######/  ###   ###     ##          ########  ######\n" +
		"   ###/        ##       #####   ##  ###         #### ###        #####     #####      #####    ###   ###     ##           ### ###  ####\n" +
		"                                                      ###                                                                     ###\n" +
		"                                               #####   ###                                                              ####   ###\n" +
		"                                             /#######  /#                                                             /######  /#\n" +
		"                                            /      ###/                                                              /     ###/\n" +
		" => Starting listen %s\n"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	configs := conf.StudyEventGo
	e := echoInit(configs)
	signal := sigInit(e)
	rds := initRedis(configs, e)

	repoContainer := container.InitRepositoryContainer(rds)
	svcContainer := container.InitServiceContainer(repoContainer)
	ctrlContainer := container.InitControllerContainer(svcContainer, repoContainer)

	if err := initHandler(e, ctrlContainer, signal); err != nil {
		e.Logger.Error("InitHandler Error")
		os.Exit(1)
	}

	startServer(configs, e)
}

func echoInit(configs *conf.ViperConfig) (e *echo.Echo) {
	e = echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
	}))

	e.HideBanner = true

	return e
}

func sigInit(e *echo.Echo) chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	go func() {
		sig := <-quit
		e.Logger.Error("Got signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
		signal.Stop(quit)
		close(quit)
	}()
	return quit
}

func initRedis(configs *conf.ViperConfig, e *echo.Echo) redis.Cmdable {
	host := configs.GetString("redis_host")
	rds := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
	})
	if _, err := rds.Ping(context.Background()).Result(); err != nil {
		e.Logger.Error("initRedis NewClient", "host", host, "err", err) // todo: logger
		os.Exit(1)
	}

	return rds
}

func initHandler(e *echo.Echo, ctrlContainer *container.ControllerContainer, signal <-chan os.Signal) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	ver := api.Group("/v1")
	sys := ver.Group("/ludovico")

	router.Route(sys, ctrlContainer)

	return nil
}

func startServer(configs *conf.ViperConfig, e *echo.Echo) {
	// Start Server
	apiServer := fmt.Sprintf("0.0.0.0:%d", configs.GetInt("port"))
	e.Logger.Debugf("Starting server, Listen[%s]", apiServer)

	fmt.Printf(banner, apiServer)
	if err := e.Start(apiServer); err != nil {
		e.Logger.Fatal(err)
	}
}

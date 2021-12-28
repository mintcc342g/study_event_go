package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"study-event-go/cmd/container"
	"study-event-go/conf"
	"study-event-go/ent"
	"study-event-go/router"
	"study-event-go/types"

	"entgo.io/ent/dialect/sql"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
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
	rds := redisInit(configs, e)
	machineryServer := machineryInit(configs, e)
	db := dbInit(configs, e)

	repoContainer, err := container.InitRepositoryContainer(db, rds, machineryServer)
	if err != nil {
		e.Logger.Error("InitRepositoryContainer Error")
		os.Exit(1)
	}

	svcContainer := container.InitServiceContainer(repoContainer)
	ctrlContainer := container.InitControllerContainer(svcContainer, repoContainer)

	if err := handlerInit(e, ctrlContainer, signal); err != nil {
		e.Logger.Error("handlerInit Error")
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

func dbInit(configs *conf.ViperConfig, e *echo.Echo) *ent.Client {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC",
		configs.GetString("db_user"),
		configs.GetString("db_pass"),
		configs.GetString("db_host"),
		configs.GetInt("db_port"),
		configs.GetString("db_name"),
	)

	drv, err := sql.Open(configs.GetString("db"), dbURI)
	if err != nil {
		e.Logger.Error("dbInit Open ", " uri: ", dbURI, " err: ", err) // todo: logger
		os.Exit(1)
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv))
}

func redisInit(configs *conf.ViperConfig, e *echo.Echo) redis.Cmdable {
	host := configs.GetString("redis_host")
	rds := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
	})
	if _, err := rds.Ping(context.Background()).Result(); err != nil {
		e.Logger.Error("redisInit NewClient", "host", host, "err", err) // todo: logger
		os.Exit(1)
	}

	return rds
}

func machineryInit(configs *conf.ViperConfig, e *echo.Echo) *machinery.Server {
	broker := configs.GetString("broker_host")

	mConf := &config.Config{
		DefaultQueue:  types.DefaultMachineryQueue,
		Broker:        broker,
		ResultBackend: "eager",
	}

	server, err := machinery.NewServer(mConf)
	if err != nil {
		e.Logger.Error("machineryInit NewServer", "broker", broker, "err", err) // todo: logger
		os.Exit(1)
	}

	return server
}

func handlerInit(e *echo.Echo, ctrlContainer *container.ControllerContainer, signal <-chan os.Signal) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	ver := api.Group("/v1")
	sys := ver.Group("/assault-lily")

	router.AssaultLilyRoute(sys, ctrlContainer)

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

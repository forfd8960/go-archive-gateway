package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	rpcClient "github.com/forfd8960/go-archive-gateway/archive-client"
	"github.com/forfd8960/go-archive-gateway/conf"
	"github.com/forfd8960/go-archive-gateway/handler"
)

func main() {
	var configFile = flag.String("config", "conf/config.toml", "rpc server config")
	flag.Parse()

	if err := conf.LoadConfig(*configFile); err != nil {
		log.Fatalf("load config error: %v\n", err)
	}

	ctx := context.Background()
	rpcAddrs := conf.ArchiveConf.ArchiveRPC.Addrs
	rpcTimeout := conf.ArchiveConf.ArchiveRPC.Timeout.Duration
	if rpcTimeout <= 0 {
		rpcTimeout = 6 * time.Second
	}

	archiveHDL := handler.NewArchiveHandler(rpcClient.NewArchiveClient(ctx, rpcAddrs, rpcTimeout))

	router := gin.Default()
	router.LoadHTMLGlob("asset/*.html")
	router.GET("/", func(c *gin.Context) {
		archiveHDL.GetArchiveList(c)
	})

	router.GET("/search", func(c *gin.Context) {
		archiveHDL.SearchItem(c)
	})

	listenAddrs := os.Getenv("PORT")
	router.Run(listenAddrs)
}

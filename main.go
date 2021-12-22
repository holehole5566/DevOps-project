package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/holehole5566/goproject/model/mysql"
	"github.com/holehole5566/goproject/router"

)

func init() {
	mysql.Setup()
}

func main() {

	gin.SetMode("debug")
	routers := router.InitRouters()
	endPoint := fmt.Sprintf(":%d", 80)
	maxHeaderBytes := 1 << 20
	c := &tls.Config{MinVersion: tls.VersionTLS12}

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routers,
		MaxHeaderBytes: maxHeaderBytes,
		TLSConfig:      c,
	}

	server.ListenAndServe()

}


package main

import (
	"fmt"
	"gotweet/internal/config"
	"gotweet/internal/handler/user"
	userRepository "gotweet/internal/repository/user"
	userService "gotweet/internal/service/user"
	"gotweet/pkg/internalsql"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := internalsql.ConnectMySql(cfg)
	if err != nil {
		log.Fatal(err)
	}
	validate := validator.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.GET("/check-health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "It works",
		})
	})
	// r.SetTrustedProxies([]string{"192.168.2.2"})
	userRepo := userRepository.NewRepository(db)
	userService := userService.NewService(cfg, userRepo)
	userhandler := user.NewHandler(r, validate, userService)
	userhandler.RouteList(cfg.SecretJwt)
	r.GET("/", func(ctx *gin.Context) {
		fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
	})
	r.Run(server)
}

package main

import (
	"fmt"
	"gotweet/internal/config"
	"gotweet/internal/handler/post"
	"gotweet/internal/handler/user"
	postRepository "gotweet/internal/repository/post"
	userRepository "gotweet/internal/repository/user"
	postService "gotweet/internal/service/post"
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
	postRepo := postRepository.NewPostRepository(db)
	userService := userService.NewService(cfg, userRepo)
	postService := postService.NewPostService(cfg, postRepo)
	userhandler := user.NewHandler(r, validate, userService)
	postHandler := post.NewPostHandler(r, validate, &postService)
	userhandler.RouteList(cfg.SecretJwt)
	postHandler.RoutePostList(cfg.SecretJwt)
	r.GET("/", func(ctx *gin.Context) {
		fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
	})
	r.Run(server)
}

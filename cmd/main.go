package main

import (
	"fmt"
	"gotweet/internal/config"
	"gotweet/internal/handler/creator"
	"gotweet/internal/handler/locations"
	"gotweet/internal/handler/post"
	socketRoom "gotweet/internal/handler/room"
	"gotweet/internal/handler/user"
	"gotweet/internal/repository/address"
	creatorRepository "gotweet/internal/repository/creator"
	creatoraddress "gotweet/internal/repository/creator-address"
	creatorContact "gotweet/internal/repository/creator-contact"
	creatordata "gotweet/internal/repository/creator-data"
	creatoremail "gotweet/internal/repository/creator-email"
	"gotweet/internal/repository/email"
	locationsRepository "gotweet/internal/repository/locations"
	postRepository "gotweet/internal/repository/post"
	"gotweet/internal/repository/status"
	userRepository "gotweet/internal/repository/user"
	creatorService "gotweet/internal/service/creator"
	locationsService "gotweet/internal/service/locations"
	postService "gotweet/internal/service/post"
	userService "gotweet/internal/service/user"
	"gotweet/pkg/internalsql"
	socket "gotweet/websocket"
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
	locationRepo := locationsRepository.NewLocationsRepository(db)
	creatorRepo := creatorRepository.NewCreatorRepository(db)
	statusRepo := status.NewStatusRepository(db)
	addressRepo := address.NewAddressRepository(db)
	emailRepo := email.NewEmailRepository(db)
	creatorContactRepo := creatorContact.NewCreatorContactRepository(db)
	creatorAddressRepo := creatoraddress.NewCreatorAddressRepository(db)
	creatorEmailRepo := creatoremail.NewCreatorEmailRepository(db)
	creatorDataRepo := creatordata.NewCreatorDataRepository(db)
	userService := userService.NewService(cfg, userRepo)
	postService := postService.NewPostService(cfg, postRepo)
	locationsServces := locationsService.NewLocationService(r, locationRepo)
	creatorService := creatorService.NewCreatorService(statusRepo, addressRepo, emailRepo, creatorContactRepo, creatorRepo, creatorAddressRepo, creatorEmailRepo, creatorDataRepo, cfg)
	userhandler := user.NewHandler(r, validate, userService)
	postHandler := post.NewPostHandler(r, validate, postService)
	locationHandler := locations.NewLocationHandler(r, validate, locationsServces)
	creatorHandler := creator.NewCreatorHandler(r, validate, creatorService)
	userhandler.RouteList(cfg.SecretJwt)
	postHandler.RoutePostList(cfg.SecretJwt)
	locationHandler.RoutePostList(cfg.SecretJwt)
	creatorHandler.CreatorHandlerRouteList(cfg.SecretJwt)
	// Websocket

	hub := socket.NewHub()
	socketHandler := socketRoom.NewHandler(r, hub)
	socketHandler.RouteList(cfg.SecretJwt)
	go hub.Run()
	r.GET("/", func(ctx *gin.Context) {
		fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
	})

	r.Run(server)
}

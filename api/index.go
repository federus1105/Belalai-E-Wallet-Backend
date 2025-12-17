package handler

import (
	"fmt"
	"net/http"

	"github.com/Belalai-E-Wallet-Backend/vercel/config"
	"github.com/Belalai-E-Wallet-Backend/vercel/middleware"
	routers "github.com/Belalai-E-Wallet-Backend/vercel/router"
	"github.com/gin-gonic/gin"
)

var App *gin.Engine

func Handler(w http.ResponseWriter, r *http.Request) {
	if App == nil {
		App = setupApp()
	}
	App.ServeHTTP(w, r)
}

func setupApp() *gin.Engine {

	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.CORSMiddleware)

	// --- CONNECT DATABASE ---
	db, err := config.ConnectDB()
	if err != nil {
		panic("DB connection failed: " + err.Error())
	}

	// --- CONNECT REDIS ---
	rdb, err := config.NewRedis()
	if err != nil {
		panic("Redis connection failed: " + err.Error())
	}

	routers.InitRouter(db, rdb)
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Success": true,
			"Message": "Backend is running 🚀",
		})
	})

	fmt.Println("Router initialized successfully")
	return app
}

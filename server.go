package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// "github.com/gin-gonic/gin/binding"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/platformsh/template-golang/config/maria"
	"github.com/platformsh/template-golang/config/uploadcare"
	"github.com/platformsh/template-golang/handler/datahdl"
	"github.com/platformsh/template-golang/handler/mockhdl"
	"github.com/platformsh/template-golang/handler/userhdl"
	"github.com/platformsh/template-golang/repository/datarps"
	"github.com/platformsh/template-golang/repository/userrps"
	"github.com/platformsh/template-golang/service/datasvc"
	"github.com/platformsh/template-golang/service/usersvc"
	"github.com/uploadcare/uploadcare-go/upload"
)

type Test struct {
	Username string `json:"username"`
}

func main() {

	// The Config Reader library provides Platform.sh environment information mapped to Go structs.
	// config, err := psh.NewRuntimeConfig()
	// if err != nil {
	// 	panic("Not in a Platform.sh Environment.")
	// }

	// port := config.Port()
	port := "8080" // test

	maria, err := maria.Connect( /**config*/ )
	if err != nil {
		log.Panic("can't connect to maria")
	}

	router := gin.New()
	router.Use(gin.Logger(), cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Access-Control-Allow-Origin", "Content-Type", "User-Agent"},
	}))

	// initialize repository
	userRepo := userrps.NewUserRepository(maria)
	dataRepo := datarps.NewDataRepository(maria)

	// initialize service
	userService := usersvc.NewUserService(userRepo)
	dataService := datasvc.NewDataService(dataRepo)

	// initialize handler
	userHandler := userhdl.NewUserHandler(userService)
	dataHandler := datahdl.NewDataHandler(dataService, userService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "welcome",
		})
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "healthy",
		})
	})

	// mock api
	mockApi := router.Group("/mock")
	mockApi.POST("/login", mockhdl.Login)
	mockApi.POST("/logout", mockhdl.Logout)
	mockApi.POST("/datas/chart", mockhdl.ChartData)
	mockApi.POST("/datas/paginate", mockhdl.PaginateData)

	// user api
	userApi := router.Group("/users")
	userApi.POST("/", userHandler.Create)

	// auth api
	authApi := router.Group("/auth")
	authApi.POST("/login", userHandler.Login)
	authApi.POST("/logout", userHandler.Logout)

	// data api
	dataApi := router.Group("/datas")
	dataApi.POST("/", dataHandler.Create)
	dataApi.DELETE("/:uuid", dataHandler.Delete)
	dataApi.GET("/chart", dataHandler.ChartData)
	dataApi.GET("/paginate", dataHandler.PaginateData)

	client, err := uploadcare.Initialize()
	if err != nil {
		log.Fatal("failed to initialize uploadcare")
	}

	// image
	imageApi := router.Group("/images")
	imageApi.POST("/upload", func(c *gin.Context) {
		c.Request.ParseMultipartForm(10 * 1024 * 1024)

		if file, handler, err := c.Request.FormFile("myfile"); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			defer file.Close()

			fmt.Println("FILE INFO")
			fmt.Println("FILE NAME : ", handler.Filename)
			fmt.Println("FILE SIZE : ", handler.Size)
			fmt.Println("FILE TYPE : ", handler.Header.Get("Content-Type"))

			uploadSvc := upload.NewService(client)

			params := upload.FileParams{
				Data:        file,
				Name:        handler.Filename,
				ContentType: handler.Header.Get("Content-Type"),
			}
			fID, err := uploadSvc.File(context.Background(), params)
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "failed to upload file",
				})
			} else {
				c.JSON(http.StatusOK, map[string]interface{}{
					"url": fmt.Sprintf("https://ucarecdn.com/%s/", fID),
				})
			}
		}
	})

	// pusher trigger by raspy python
	

	router.Run(":" + port)
}

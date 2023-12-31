package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AkashJam/blog-server/server/pkg/article"
	"github.com/AkashJam/blog-server/server/pkg/http/rest"
	"github.com/gin-gonic/gin"
)

// gin.Accounts is a shortcut for map[string]string
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

var admins = gin.Accounts{
	"foo":    "bar",
	"austin": "1234",
	"lena":   "hello2",
}

func main() {
	log.Println("Starting server...")
	router := gin.Default()
	
	// Group using gin.BasicAuth() middleware
	// admin := router.Group("/admin", gin.BasicAuth(admins))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	// admin.GET("/secrets", func(c *gin.Context) {
	// 	// get user, it was set by the BasicAuth middleware
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	// 	if secret, ok := secrets[user]; ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	// 	}
	// })

	// rest.AccountRoutes(router.Group("/articles"))
	rest.NewHandler(&rest.Config{
		R: router,
		ArticleService: article.NewService(),
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
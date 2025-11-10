package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/golang-mini-project/url-shortener/config"
	"github.com/yafiakmal/golang-mini-project/url-shortener/internal/database"
	"github.com/yafiakmal/golang-mini-project/url-shortener/internal/handler"
)

type Visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu       sync.Mutex
	visitors = make(map[string]*Visitor)
)

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		log.Println("Create new limiter for: ", ip)
		limiter := rate.NewLimiter(1, 3) // 1 req/detik, burst 3
		visitors[ip] = &Visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

// hapus IP yang idle lebih dari 5 menit
func cleanupVisitors() {
	log.Println("cleanupVisitors started")
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 5*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func RateLimiterPerIP() gin.HandlerFunc {
	go cleanupVisitors()

	return func(c *gin.Context) {
		ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		limiter := getVisitor(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			return
		}
		c.Next()
	}
}

func main() {
	config.EnvLoad()

	// initialize postgres
	db, err := database.Connect(config.GetDBConfig())
	if err != nil {
		log.Panic(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"}, // atau ganti dengan domain frontend kamu
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	server.Use(RateLimiterPerIP())
	server.POST("/shortener", func(ctx *gin.Context) {
		handler.AddUrlHandler(ctx, db)
	})
	server.GET("/urls", func(ctx *gin.Context) {
		handler.GetAllUrlHandler(ctx, db)
	})
	server.GET("/:name", func(ctx *gin.Context) {
		handler.Redirect(ctx, db)
	})
	server.DELETE("/:name", func(ctx *gin.Context) {
		handler.DeleteUrl(ctx, db)
	})
	err = server.Run(":8080")
	if err != nil {
		log.Fatalln("failed to start server ", err)
	}
}

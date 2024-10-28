package bootstrap

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/volkankocaali/bi-taksi-case/config"
	"github.com/volkankocaali/bi-taksi-case/database"
	"github.com/volkankocaali/bi-taksi-case/pkg/circuitbreaker"
	"github.com/volkankocaali/bi-taksi-case/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"syscall"
	"time"
)

func StartServer(cfg config.Config, r *mux.Router) {
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: corsOptions(r),
	}

	// Graceful shutdown
	c := make(chan os.Signal, 1)

	go func() {
		logger.Printf("Starting server on port %s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server:", err)
			c <- syscall.SIGTERM
		}
	}()

	<-c
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Failed to shutdown server:", err)
	}
	logger.Info("Server shutdown successfully")
}

func LoggerInit(port string) {
	logger.SetLogger(logger.NewZapAdapter())
	logger.Info(fmt.Sprintf("Starting API on port %s", port))
}

func RouterInit() *mux.Router {
	r := mux.NewRouter()

	r.Use(contentTypeApplicationJsonMiddleware)
	return r
}

func MongoInit(cfg *config.Config) *mongo.Client {
	mongo, err := database.ConnectMongoDB(cfg.Database.URI)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB:", err)
	}

	return mongo
}

func CircuitBreakerInit() {
	circuitbreaker.InitCircuitBreaker()
}

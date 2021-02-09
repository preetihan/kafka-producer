package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run(":" + getEnv("PORT", Port))
}

func setupRouter() *gin.Engine {
	r := gin.New()

	environment := getEnv("ENV", Environment)
	ctxPath := getEnv("CTX_PATH", Ctxpath)
	bootstrapServers := getEnv("BOOTSTRAP_SERVERS", BootstrapServers)
	topics := getEnv("TOPICS", Topics)
	partition := getEnv("PARTITION", Partition)
	groupID := getEnv("GROUP_ID", GroupID)

	log.Println("ENV = " + environment)
	log.Println("CTX_PATH = " + ctxPath)
	log.Println("BOOTSTRAP_SERVERS = " + bootstrapServers)
	log.Println("TOPICS = " + topics)
	log.Println("PARTITION = " + partition)
	log.Println("GROUP_ID = " + groupID)

	if environment != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, ctxPath+"/api/health"),
		gin.Recovery(),
	)

	api := r.Group(ctxPath + "/api")
	{
		api.GET("/health", checkHealth)
		api.GET("/produce", produce)
	}

	return r
}

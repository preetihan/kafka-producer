package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func produce(c *gin.Context) {
	var m Message

	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("Got message: ", m.Message)

	bootstrapServers := getEnv("BOOTSTRAP_SERVERS", BootstrapServers)
	topics := getEnv("TOPICS", Topics)
	partition, _ := strconv.Atoi(getEnv("PARTITION", Partition))

	conn, err := kafka.DialLeader(context.Background(), "tcp", bootstrapServers, topics, partition)
	if err != nil {
		log.Println("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(m.Message)},
	)
	if err != nil {
		log.Println("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Println("failed to close writer:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"status": "Published",
	})
}

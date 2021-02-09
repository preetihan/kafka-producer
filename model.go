package main

type Message struct {
	Message string `uri:"message" binding:"required"`
}

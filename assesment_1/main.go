package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Message struct {
	UserID string
	Token  string
	Data   string
}

var (
	messages    = make(chan Message, 100)
	cache       = make([]Message, 0)
	cacheMutex  sync.Mutex
	validTokens = map[string]string{
		"user1": "token1",
		"user2": "token2",
		// Добавьте других пользователей и их токены
	}
)

func main() {
	go startUsers()
	go cacheMessages()
	go startWorker()

	// Обработка сигнала завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Graceful shutdown
	fmt.Println("Shutting down gracefully...")
	writeCacheToFile()
}

func startUsers() {
	users := []string{"user1", "user2"}
	tokens := []string{"token1", "token2"}
	for i := 0; i < 10; i++ {
		for j, user := range users {
			messages <- Message{UserID: user, Token: tokens[j], Data: fmt.Sprintf("Message %d from %s", i, user)}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func cacheMessages() {
	for msg := range messages {
		if isValidToken(msg.UserID, msg.Token) {
			cacheMutex.Lock()
			cache = append(cache, msg)
			cacheMutex.Unlock()
		} else {
			fmt.Printf("Invalid token from user: %s\n", msg.UserID)
		}
	}
}

func startWorker() {
	for {
		time.Sleep(1 * time.Second)
		writeCacheToFile()
	}
}

func writeCacheToFile() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if len(cache) == 0 {
		return
	}

	file, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	for _, msg := range cache {
		_, err := file.WriteString(fmt.Sprintf("UserID: %s, Data: %s\n", msg.UserID, msg.Data))
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}

	cache = make([]Message, 0)
}

func isValidToken(userID, token string) bool {
	validToken, exists := validTokens[userID]
	return exists && validToken == token
}

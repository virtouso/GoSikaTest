package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/virtouso/GoSikaTest/Domain"
	"github.com/virtouso/GoSikaTest/Dto"
	"github.com/virtouso/GoSikaTest/Infra"
	"net/http"
	"sync"
	"time"
)

func main() {
	Infra.Init()
	router := gin.Default()
	router.POST("/SendUsers", handleUserSubmit)

	router.Run(":8080")
}

func handleUserSubmit(c *gin.Context) {
	var input Dto.UserBulk

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received user input: %+v\n", input)

	// handle here
	insertUsersConcurrently(input.Users)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}

func insertUsersConcurrently(users []Dto.UserDto) {
	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a buffered channel to limit the number of concurrent inserts
	maxGoroutines := 10
	sem := make(chan struct{}, maxGoroutines)

	for _, user := range users {
		wg.Add(1)
		sem <- struct{}{} // Block when maxGoroutines are running

		// Start a new goroutine to insert the user
		go func(user Dto.UserDto) {
			defer wg.Done()
			defer func() { <-sem }() // Release the spot in the channel

			dbAddresses := []Domain.Address{}
			for _, item := range user.Addresses {
				dbAddresses = append(dbAddresses, Domain.Address{
					Address: item.Address,
				})
			}
			Infra.CrateUser(&Domain.User{

				Name:      user.Name,
				Email:     user.Email,
				Addresses: dbAddresses,
				CreatedAt: time.Now(),
			})

		}(user)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

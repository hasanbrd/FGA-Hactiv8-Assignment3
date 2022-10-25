package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go jsonUpdate()
	router := gin.Default()
	router.Run("localhost:8080")
}

func jsonUpdate() {
	for {
		var waterStatus, windStatus string
		waterValue, windValue := rand.Intn(100), rand.Intn(100)
		jsonString := fmt.Sprintf(`{"water": %d, "wind": %d }`, waterValue, windValue)
		jsonFile, err := os.Create("status.json")
		if err != nil {
			log.Println("Error in jsonUpdate:", err.Error())
			continue
		}
		_, err = jsonFile.Write([]byte(jsonString))
		if err != nil {
			log.Println("Error in jsonUpdate:", err.Error())
			continue
		}

		if waterValue < 5 {
			waterStatus = "Aman"
		} else if waterValue < 8 {
			waterStatus = "Siaga"
		} else {
			waterStatus = "Bahaya"
		}

		if windValue < 6 {
			windStatus = "Aman"
		} else if windValue < 15 {
			windStatus = "Siaga"
		} else {
			windStatus = "Bahaya"
		}

		log.Printf("Environmental Status: Water: %d meter (%s), Wind: %d meter/detik (%s) \n",
			waterValue, waterStatus, windValue, windStatus)
		time.Sleep(15 * time.Second)
	}
}

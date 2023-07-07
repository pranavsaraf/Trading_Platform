package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	accessToken = ""
	baseURL     = "https://api.dhan.co"
)

type Holdings struct {
}

func main() {
	router := gin.Default()

	router.GET("/holdings", Getholdings)
	_ = router.Run(":8080")
}

func Getholdings(c *gin.Context) {
	loadURL := "https://api.dhan.co" + "/holdings"
	payload := []map[string]interface{}{}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		c.JSON(0, gin.H{"error": err.Error()})
		return
	}
	body := bytes.NewBuffer(payloadJSON)
	req, err := http.NewRequest("GET", loadURL, body)
	if err != nil {
		c.JSON(0, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("access-token", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(0, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(0, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": string(response)})
}

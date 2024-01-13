package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"konsultanku-app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendRequest(c *gin.Context, url string, data map[string]interface{}) (userData map[string]interface{}, err error) {

	url += database.ApiKey

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err.Error())
		return userData, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err.Error())
		return userData, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err.Error())
		return userData, err
	}
	defer resp.Body.Close()

	// Handle the response here
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err.Error())
		return userData, err
	}

	userData = responseBody
	if resp.StatusCode == 200 {
		return userData, nil
	}
	return userData, err
}
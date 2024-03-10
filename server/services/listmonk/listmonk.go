package listmonk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"schej.it/server/logger"
)

// Adds the given user to the Listmonk contact list
// If subscriberId is not nil, then UPDATE the user instead of adding user
func AddUserToListmonk(email string, firstName string, lastName string, picture string, subscriberId *int) {
	url := os.Getenv("LISTMONK_URL")
	username := os.Getenv("LISTMONK_USERNAME")
	password := os.Getenv("LISTMONK_PASSWORD")
	listIdString := os.Getenv("LISTMONK_LIST_ID")

	listId, err := strconv.Atoi(listIdString)
	if err != nil {
		logger.StdErr.Println(err)
		return
	}

	// Create new subscriber
	body, _ := json.Marshal(bson.M{
		"email":  email,
		"name":   firstName + " " + lastName,
		"status": "enabled",
		"lists":  bson.A{listId},
		"attribs": bson.M{
			"firstName": firstName,
			"lastName":  lastName,
			"picture":   picture,
		},
	})
	bodyBuffer := bytes.NewBuffer(body)

	var req *http.Request
	if subscriberId != nil {
		// Existing subscriber
		req, _ = http.NewRequest("PUT", fmt.Sprintf("%s/api/subscribers/%d", url, *subscriberId), bodyBuffer)
	} else {
		// New subscriber
		req, _ = http.NewRequest("POST", fmt.Sprintf("%s/api/subscribers", url), bodyBuffer)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.StdErr.Println(err)
		return
	}
	defer resp.Body.Close()
}

// Check if the user is already in listmonk
// Returns a bool representing whether the subscriber exists and the id of the subscriber if it does exist
func DoesUserExist(email string) (bool, *int) {
	url := os.Getenv("LISTMONK_URL")
	username := os.Getenv("LISTMONK_USERNAME")
	password := os.Getenv("LISTMONK_PASSWORD")

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/subscribers?query=subscribers.email='%s'", url, email), nil)
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.StdErr.Println(err)
		return false, nil
	}
	defer resp.Body.Close()

	var response struct {
		Data struct {
			Results []struct {
				Id int `json:"id"`
			} `json:"results"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		logger.StdErr.Println(err)
		return false, nil
	}

	if len(response.Data.Results) > 0 {
		return true, &response.Data.Results[0].Id
	} else {
		return false, nil
	}
}

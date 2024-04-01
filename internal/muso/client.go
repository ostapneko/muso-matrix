package muso

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var musoURL string

func init() {
	value, ok := os.LookupEnv("MUSO_URL")
	if !ok {
		panic("MUSO_URL environment variable is not set")
	}
	musoURL = value
}

func GetTrackInfo() (*Track, error) {
	data, err := GetData("player:player/data")
	if err != nil {
		return nil, err
	}

	valueDoc := data["value"].(map[string]interface{})
	trackRolesDoc := valueDoc["trackRoles"].(map[string]interface{})
	mediaDataDoc := trackRolesDoc["mediaData"].(map[string]interface{})
	metaDataDoc := mediaDataDoc["metaData"].(map[string]interface{})

	track := &Track{
		Title:  trackRolesDoc["title"].(string),
		Artist: metaDataDoc["artist"].(string),
		Album:  metaDataDoc["album"].(string),
	}

	return track, nil
}

func GetData(path string) (map[string]interface{}, error) {
	rolesQS := strings.Join(Roles, ",")
	url := fmt.Sprintf("%s/api/getData?path=%s&roles=%s", musoURL, path, rolesQS)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var result []interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Map the result to a map[string]interface{} with values in the same order as Roles
	data := make(map[string]interface{})
	for i, role := range Roles {
		if i < len(result) {
			data[role] = result[i]
		}
	}

	return data, nil
}

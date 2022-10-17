package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	/*
	   "log"
	   "log/syslog"
	*/
	"net/http"
	"os"
	"os/exec"
)

func prefetchImages() error {

	cmd := exec.Command("glance-cache-prefetcher")
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("glance-cache-prefetcher failed to execute properly: %v", err)
	}

	return nil
}

func queueImages(hostname string, imageList []string) error {

	for _, image := range imageList {
		cmd := exec.Command("glance-cache-manage", "--host=", hostname, "queue-image", image)
		err := cmd.Run()

		if err != nil {
			return fmt.Errorf("glance-cache-manage failed to execute properly: %v", err)
		} else {
			fmt.Printf("Image %s queued", image)
		}
	}

	return nil
}

func getBody(method string, url string, headers map[string]string, body []byte) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	var bodyBytes []byte

	if res.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(res.Body)
	} else if err != nil {
		return nil, err
	} else {
		return nil, fmt.Errorf("The remote end did not return a HTTP 200 (OK) response.")
	}

	return bodyBytes, nil

}

func getImages(authToken string) ([]string, error) {

	type GlanceDetailResponse struct {
		Images []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
			ID     string `json:"id"`
		}
	}

	method := "GET"
	url := "http://192.168.1.2:9292/v1.1/images/detail"
	headers := map[string]string{"X-Auth-Token": authToken}

	bodyBytes, err := getBody(method, url, headers, nil)

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve the response body from the Glance API server: %v", err)
	}

	var glance GlanceDetailResponse
	err = json.Unmarshal(bodyBytes, &glance)

	if err != nil {
		return nil, fmt.Errorf("unable to parse the JSON response:", err)
	}

	imageList := make([]string, 10)

	for _, image := range glance.Images {
		if image.Status == "active" {
			imageList = append(imageList, image.ID)
		}
	}

	return imageList, nil

}

func getToken() (string, error) {

	type TokenResponse struct {
		Auth []struct {
			Token struct {
				Expires string `json:"expires"`
				ID      string `json:"id"`
			}
		}
	}

	method := "POST"
	url := "http://192.168.1.2:5000/v2.0/tokens"
	headers := map[string]string{"Content-type": "application/json"}
	creds := []byte(`{"auth":{"passwordCredentials":{"username": "glance", "password":"<password>"}, "tenantId":"<tenantkeygoeshere>"}}`)

	bodyBytes, err := getBody(method, url, headers, creds)

	if err != nil {
		return "", err
	}

	var keystone TokenResponse
	err = json.Unmarshal(bodyBytes, &keystone)

	if err != nil {
		return "", err
	}

	authToken := string((keystone.Auth[0].Token.ID))

	return authToken, nil
}

func main() {

	/*
	   slog, err := syslog.New(syslog.LOG_ERR, "[fido]")

	   if err != nil {
	       log.Fatalf("unable to connect to syslog: %v", err)
	       os.Exit(1)
	   } else {
	       defer slog.Close()
	   }
	*/

	hostname, err := os.Hostname()

	if err != nil {
		// slog.Err("Hostname not captured")
		os.Exit(1)
	}

	authToken, err := getToken()

	if err != nil {
		// slog.Err("The authentication token from the Glance API server was not retrieved")
		os.Exit(1)
	}

	imageList, err := getImages(authToken)

	err = queueImages(hostname, imageList)

	if err != nil {
		// slog.Err("Could not queue the images for pre-fetching")
		os.Exit(1)
	}

	err = prefetchImages()

	if err != nil {
		// slog.Err("Could not queue the images for pre-fetching")
		os.Exit(1)
	}

	return
}

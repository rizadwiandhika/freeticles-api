package cloudmersive

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
)

type cloudmersive struct {
	result response
}

type response struct {
	Successful            bool    `json:"Successful"`
	Score                 float64 `json:"Score"`
	ClassificationOutcome string  `json:"ClassificationOutcome"`
}

func NewImageAnalyzer() *cloudmersive {
	return &cloudmersive{}
}

func (c *cloudmersive) IsNSFW(imagePath string) (bool, error) {
	url := "https://api.cloudmersive.com/image/nsfw/classify"
	method := "POST"
	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(imagePath)
	if errFile1 != nil {
		return false, errFile1
	}
	defer file.Close()

	part1, errFile1 := writer.CreateFormFile("imageFile", filepath.Base(imagePath))
	if errFile1 != nil {
		return false, errFile1
	}
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return false, errFile1
	}
	err := writer.Close()
	if err != nil {
		return false, errFile1
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return false, errFile1
	}

	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Apikey", config.ENV.CLOUDMERSIVE_API_KEY)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return false, errFile1
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, errFile1
	}

	err = json.Unmarshal(body, &c.result)
	if err != nil {
		return false, err
	}

	// the less the score, the more chance image is safe
	const SAFE_THRESHOLD = 0.2

	// if the score is less that safe limit, then it's not NSFW
	if c.result.Successful && c.result.Score < SAFE_THRESHOLD {
		return false, nil
	}

	return true, nil
}

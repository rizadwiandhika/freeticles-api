package libre

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rizadwiandhika/miniproject-backend-alterra/third-parties/translate"
)

type libre struct {
}

type libreResponse struct {
	TranslatedText string `json:"translatedText"`
}

func NewTranslate() *libre {
	return &libre{}
}

func (l *libre) Translate(t translate.TranslateCore) (string, error) {
	var payloadEncoded []byte
	var err error
	var result libreResponse

	url := "https://libretranslate.de/translate"
	contentType := "application/json"
	payload := map[string]string{
		"q":      t.Text,
		"target": t.Target,
		"source": "id",
		"type":   "text",
	}
	if payloadEncoded, err = json.Marshal(payload); err != nil {
		return "", err
	}
	body := bytes.NewBuffer(payloadEncoded)

	responseBuffer, err := http.Post(url, contentType, body)
	if err != nil {
		return "", err
	}
	defer responseBuffer.Body.Close()

	responseBody, err := ioutil.ReadAll(responseBuffer.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return "", err
	}

	return result.TranslatedText, nil
}

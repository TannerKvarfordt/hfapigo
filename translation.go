package hfapigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	RecommendedRussianToEnglishModel = "Helsinki-NLP/opus-mt-ru-en"
)

// Request structure for the Translation endpoint
type TranslationRequest struct {
	// (Required) a string to be translated in the original languages
	Input string `json:"inputs,omitempty"`

	Options Options `json:"options,omitempty"`
}

// Response structure from the Translation endpoint
type TranslationResponse struct {
	// The translated Input string
	TranslationText string `json:"translation_text"`
}

func SendTranslationRequest(model string, request *TranslationRequest) (*TranslationResponse, error) {
	endpoint := APIBaseURL + model
	if request == nil {
		return nil, errors.New("nil TranslationRequest")
	}

	jsonBuf, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBuf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	SetAuthorizationHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Not sure why they return an array of responses since they
	// do not accept an array of inputs. But for now, just read
	// the array and return the single response.
	tresps := make([]*TranslationResponse, 1)
	err = json.Unmarshal(respBody, &tresps)
	if err != nil {
		return nil, errors.New(string(respBody))
	}

	if len(tresps) == 0 {
		return nil, errors.New("empty response list received")
	}

	if tresps[0] == nil {
		return nil, errors.New("nil response received")
	}

	return tresps[0], nil
}

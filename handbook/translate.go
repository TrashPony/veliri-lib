package handbook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func getYandexToken(oauthToken string) (string, error) {
	url := "https://iam.api.cloud.yandex.net/iam/v1/tokens"
	data := map[string]string{
		"yandexPassportOauthToken": oauthToken,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["iamToken"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse token from response")
	}

	return token, nil
}

type Translation struct {
	DetectedLanguageCode string `json:"detectedLanguageCode"`
	Text                 string `json:"text"`
}

type TranslationResponse struct {
	Translations []Translation `json:"translations"`
}

func TranslateText(oAuthToken, folderID, targetLanguage string, texts []string) (TranslationResponse, error) {
	token, err := getYandexToken(oAuthToken)
	if err != nil {
		return TranslationResponse{}, err
	}

	data, err := translateText(token, folderID, targetLanguage, texts)
	if err != nil {
		return TranslationResponse{}, err
	}

	return data, nil
}

func translateText(iamToken, folderID, targetLanguage string, texts []string) (TranslationResponse, error) {
	url := "https://translate.api.cloud.yandex.net/translate/v2/translate"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + iamToken,
	}

	postData := map[string]interface{}{
		"targetLanguageCode": targetLanguage,
		"texts":              texts,
		"folderId":           folderID,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return TranslationResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return TranslationResponse{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TranslationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return TranslationResponse{}, fmt.Errorf("failed to translate text: %s", resp.Status)
	}

	var result TranslationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return TranslationResponse{}, err
	}

	return result, nil
}

package dontpad

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL = "https://api.dontpad.com/"
)

type ClientAPI interface {
	UploadFile(content, key string) error
	DownloadFile(key string) (string, error)
}

type DontPad struct {
	client *http.Client
}

func NewDontPad(client *http.Client) *DontPad {
	return &DontPad{client: client}
}

func (d *DontPad) UploadFile(content, key string) error {
	_, err := d.DownloadFile(key)
	if err != nil {
		return err
	}
	form := url.Values{}
	form.Add("text", content)
	lastDateTime := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	form.Add("lastModified", lastDateTime)
	form.Add("force", "false")
	data := strings.NewReader(form.Encode())
	req, err := http.NewRequest(http.MethodPost, baseURL+key, data)
	if err != nil {
		return err
	}
	req.Header.Set("authority", "api.dontpad.com")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://dontpad.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://dontpad.com/")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err := d.client.Do(req)
	if err != nil {
		return err
	}
	log.Println("Upload file: ", resp.StatusCode)
	return nil
}

func (d *DontPad) DownloadFile(key string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s.body.json?lastModified=0", baseURL, key), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authority", "api.dontpad.com")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://dontpad.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://dontpad.com/")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err := d.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response response
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return "", err
	}

	log.Println("download file: ", resp.StatusCode)
	return response.Body, nil
}

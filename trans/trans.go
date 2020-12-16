package trans

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "https://trans.rayjar.com"
const googlePath = "/google/auto"
const bingPath = "/bing/dictf"

func fetch(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("网络错误")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("请求错误")
	}
	return resp.Body, nil
}

// FetchGoogle 获取翻译结果
func FetchGoogle(word string) (GoogleRes, error) {
	var res GoogleRes
	q := url.QueryEscape(word)
	body, err := fetch(baseURL + googlePath + "?q=" + q)
	if err != nil {
		return res, err
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

// FetchBing 获取词典结果
func FetchBing(word string) (DictRes, error) {
	var res DictRes
	q := url.QueryEscape(word)
	body, err := fetch(baseURL + bingPath + "?q=" + q)
	if err != nil {
		return res, err
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

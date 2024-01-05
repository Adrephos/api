package clients

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
)

var endpoint = "https://api.igdb.com/v4"

func headers(request *http.Request) {
	request.Header = map[string][]string{
		"Client-ID":     {os.Getenv("TWITCH_CLIENT_ID")},
		"Authorization": {"Bearer " + os.Getenv("TWITCH_TOKEN")},
	}
}

func SearchGame(query, fields string) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		endpoint+"/games",
		bytes.NewBuffer([]byte(
			`search "`+query+`"; fields `+fields+`;`,
		)),
	)
	if err != nil {
		return nil, err
	}
	headers(req)

	return http.DefaultClient.Do(req)
}

func GetCover(id int) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		endpoint+"/covers",
		bytes.NewBuffer([]byte(
			`where id = `+strconv.Itoa(id)+`; fields url;`,
		)),
	)
	if err != nil {
		return nil, err
	}
	headers(req)

	return http.DefaultClient.Do(req)
}

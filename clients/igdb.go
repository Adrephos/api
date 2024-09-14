package clients

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Adrephos/api/utils"
)

var endpoint = "https://api.igdb.com/v4"
var authURL = "https://id.twitch.tv/oauth2/token"

func headers(request *http.Request) {
  token, err := GetToken()
  if err != nil {
    return
  }
	request.Header = map[string][]string{
		"Client-ID":     {os.Getenv("TWITCH_CLIENT_ID")},
		"Authorization": {"Bearer " + token},
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

func GetToken() (string, error) {
  token, found := utils.GetCache().Get("igdb_token")
  if found {
    return token.(string), nil
  }
	req, err := http.NewRequest(
		http.MethodPost,
		authURL,
		bytes.NewBuffer([]byte(
			`client_id=`+os.Getenv("TWITCH_CLIENT_ID")+`&client_secret=`+os.Getenv("TWITCH_SECRET")+`&grant_type=client_credentials`,
		)),
	)
	if err != nil {
		return "", err
	}
	req.Header = map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	res, err := http.DefaultClient.Do(req)
  if err != nil {
    return "", err
  }
	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)
  expires := int(data["expires_in"].(float64))
  utils.GetCache().Set("igdb_token", data["access_token"].(string), time.Second*time.Duration(expires))
  return data["access_token"].(string), nil
}

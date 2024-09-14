package handlers_v1

import (
	"encoding/json"
	"errors"
	"github.com/Adrephos/api/clients"
	"github.com/Adrephos/api/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func getField(res *http.Response, field string, _ *gin.Context) (interface{}, error, int) {
	var arr []map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&arr)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	} else if len(arr) == 0 {
		return nil, errors.New("No cover found"), http.StatusNotFound
	}
	return arr[0][field], nil, http.StatusOK
}

func getURL(c *gin.Context) (string, error) {
	query := c.Param("query")
	gamesRes, _ := clients.SearchGame(strings.Replace(query, "_", " ", -1), "cover")
	coverInter, err, code := getField(gamesRes, "cover", c)
	if err != nil {
		handlers.Response(c, false, nil, err, code)
		return "", err
	}
	coverID, ok := coverInter.(float64)
	if !ok {
		handlers.Response(c, false, nil, errors.New("No cover found"), http.StatusNotFound)
		return "", errors.New("No cover found")
	}
	coverRes, _ := clients.GetCover(int(coverID))
	coverURLInter, err, code := getField(coverRes, "url", c)
	if err != nil {
		handlers.Response(c, false, nil, err, code)
		return "", err
	}
	coverURL, ok := coverURLInter.(string)
	if !ok {
		handlers.Response(c, false, nil, errors.New("No cover found"), http.StatusNotFound)
		return "", errors.New("No cover found")
	}

	return "https:" + coverURL, nil
}

func GetCover(c *gin.Context) {
	coverURL, err := getURL(c)
	if err != nil {
		return
	}

	res := map[string]string{"url": strings.Replace(coverURL, "t_thumb", "t_cover_big", -1)}
	handlers.Response(c, true, res, nil, http.StatusOK)
}

func GetThumbnail(c *gin.Context) {
	coverURL, err := getURL(c)
	if err != nil {
		return
	}

	res := map[string]string{"url": coverURL}
	handlers.Response(c, true, res, nil, http.StatusOK)
}

package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"jiyue.im/model"
	"jiyue.im/pkg/auth"
	"jiyue.im/pkg/errno"
	"jiyue.im/pkg/token"
)

type Info struct {
	Session_key string `json:"session_key"`
	Openid      string `json:"openid"`
}

func codeToToken(c *gin.Context, code string) {
	url := os.Getenv("LOGINURL")
	loginUrl := fmt.Sprintf(url, os.Getenv("APPID"), os.Getenv("APPSECRET"), code)
	fmt.Println(loginUrl)
	client := &http.Client{}

	reqest, err := http.NewRequest("GET", loginUrl, nil)
	if err != nil {
		SendResponse(c, errno.ErrWx, nil)
	}

	response, _ := client.Do(reqest)
	status := response.StatusCode
	if status != 200 {
		SendResponse(c, errno.ErrWx, nil)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var info Info
	err = json.Unmarshal([]byte(body), &info)
	if err != nil {
		SendResponse(c, errno.ErrWx, nil)
	}

	d, err := model.GetUserByOpenId(info.Openid)
	if err != nil {
		d.OpenId = info.Openid
		err = d.RegisterByOpenId()
		if err != nil {
			SendResponse(c, errno.ErrWxOpenId, nil)
		}
	}

	t, err := token.Sign(c, token.Context{ID: d.Id, Scope: auth.USER}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}

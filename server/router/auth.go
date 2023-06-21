package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (r *Router) login(c *gin.Context) {
	url := r.oAuth.OAuthConf.AuthCodeURL("random")
	fmt.Println(url)

	c.JSON(http.StatusOK, url)
}

func (r *Router) loginCallback(c *gin.Context) {
	request := c.Request

	data, err := r.getGoogleUserInfo(request.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusConflict, "Google Info Not Found : ")
		return
	}

	type User struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail string `json:"verifiedEmail"`
		Name          string `json:"name"`
	}

	var user User

	if err := json.Unmarshal(data, &user); err != nil {
		c.JSON(http.StatusConflict, "파싱 오류")
		return
	}

	if token, err := r.paseto.CreateToken(user.Name); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, "Token Create Failed")
		return
	} else {
		c.SetCookie("oauth", token, 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
	}

}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (r *Router) getGoogleUserInfo(code string) ([]byte, error) {
	token, err := r.oAuth.OAuthConf.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get UserInfo %s\n", err.Error())
	}

	return ioutil.ReadAll(resp.Body)
}

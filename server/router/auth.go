package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/04Akaps/Video_Chat_App/reposiroty"
	"github.com/04Akaps/Video_Chat_App/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type auth struct {
	router  Router
	paseto  *reposiroty.PasetoMaker
	service *service.Service
}

func newAuth(router Router, paseto *reposiroty.PasetoMaker, service *service.Service) *auth {
	a := &auth{
		router:  router,
		paseto:  paseto,
		service: service,
	}

	baseUri := "/auth"

	router.engine.GET(baseUri+"/login", a.login)
	router.engine.GET(baseUri+"/login/callback", a.loginCallback)
	router.engine.GET(baseUri+"/check-token", a.checkToken)

	return a
}

func (r *auth) checkToken(c *gin.Context) {
	if userName, err := r.router.extractToken(c); err != nil {
		c.JSON(http.StatusCreated, "failed")
	} else {
		c.JSON(http.StatusOK, userName)
	}
}

func (r *auth) login(c *gin.Context) {
	url := r.router.oAuth.OAuthConf.AuthCodeURL("random")
	c.JSON(http.StatusOK, url)
}

func (r *auth) loginCallback(c *gin.Context) {
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

	if _, err := r.service.Mysql.GetAuthByName(user.Name); err != nil {
		if err.Error() == "sql: no rows in result set" {
			if err = r.service.Mysql.InsertAuth(user.Name, user.Email, user.ID); err != nil {
				fmt.Println(err)
				c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
				return
			}
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
			return
		}
	}

	if token, err := r.router.paseto.CreateToken(user.Name); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, "Token Create Failed")
		return
	} else {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:  "oauth",
			Value: token,
			Path:  "/",
		})
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
	}
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (r *auth) getGoogleUserInfo(code string) ([]byte, error) {
	token, err := r.router.oAuth.OAuthConf.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get UserInfo %s\n", err.Error())
	}

	return ioutil.ReadAll(resp.Body)
}

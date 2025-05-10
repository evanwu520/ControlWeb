package controllers

import (
	"fmt"
	"net/http"

	"game.com/controlWeb/cache"
	"game.com/controlWeb/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserController() *userController {
	return &userController{}
}

// key:name value:password
var admins map[string]string

var menuList []Itme

func init() {
	// menu
	menuList = append(menuList, Itme{FuncName: "贈獎", FuncUrl: "reward"})
	menuList = append(menuList, Itme{FuncName: "贈獎記錄", FuncUrl: "rewardRecords"})
	menuList = append(menuList, Itme{FuncName: "輸贏報表", FuncUrl: "winlossRecords"})

	// admin
	admins = make(map[string]string)
	admins["admin1"] = "admin1XY!"
	admins["admin2"] = "admin2MN%"
	admins["admin3"] = "admin3PD$"
	admins["admin4"] = "admin4SS*"
}

type userController struct {
}

type LoginReq struct {
	Account  string `form:"account"`
	Password string `form:"password"`
}

type LoginRsp struct {
	Token string `json:"token"`
	Menus []Itme `json:"items"`
	Msg   string
}

type Itme struct {
	FuncName string `json:"funcName"`
	FuncUrl  string `json:"funcUrl"`
}

func (userController) Login(c *gin.Context) {
	var req LoginReq
	resp := LoginRsp{}

	err := c.ShouldBind(&req)

	if err != nil {
		fmt.Println(err)
		resp.Msg = "login fail!"
		c.HTML(http.StatusOK, "login.tmpl", resp)
		return
	}

	v, exist := admins[req.Account]

	if !exist || v != req.Password {
		resp.Msg = "login fail!"
		c.HTML(http.StatusOK, "login.tmpl", resp)
		return
	}

	resp.Token = uuid.New().String()
	resp.Menus = menuList
	cache.SetToken(resp.Token, req.Account)
	c.HTML(http.StatusOK, "main.tmpl", resp)
}

type OnlinePlayersResp struct {
	Players []*OnlinePlayerInfo
}

type OnlinePlayerInfo struct {
	GameID      int
	AccountName string
}

func (userController) OnlinePlayers(c *gin.Context) {

	resp := OnlinePlayersResp{}
	getAll := true
	account := c.Param("account")

	if account != "" {
		getAll = false
	}

	players := models.NewOnlinePlayersModel().OnlinePlayerList()

	for _, info := range players {
		data := &OnlinePlayerInfo{}
		data.AccountName = info.Player.AccountName
		data.GameID = info.GameID
		resp.Players = append(resp.Players, data)

		if !getAll && info.Player.AccountName == account {
			resp.Players = resp.Players[:0]
			resp.Players = append(resp.Players, data)
			break
		}
	}

	c.HTML(http.StatusOK, "main.tmpl", resp)
}

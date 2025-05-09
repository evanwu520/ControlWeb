package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserController() *userController {
	return &userController{}
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
}

type Itme struct {
	FuncName string `json:"funcName"`
	FuncUrl  string `json:"funcUrl"`
}

func (userController) Login(c *gin.Context) {
	var req LoginReq

	err := c.ShouldBind(&req)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req.Account, req.Password)

	// TODO read menu
	resp := LoginRsp{}

	var list []Itme
	list = append(list, Itme{FuncName: "贈獎", FuncUrl: "reward"})
	list = append(list, Itme{FuncName: "贈獎記錄", FuncUrl: "rewardRecords"})
	list = append(list, Itme{FuncName: "輸贏報表", FuncUrl: "winlossRecords"})

	resp.Token = uuid.New().String()
	resp.Menus = list

	c.HTML(http.StatusOK, "main.tmpl", resp)
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"game.com/controlWeb/models"
	"github.com/gin-gonic/gin"
)

func NewBonusController() *bonusController {
	return &bonusController{}
}

type bonusController struct {
}

type RewardReq struct {
	AccountName string `form:"accountName"`
	DeskNo      uint   `form:"deskNo"`
	BonusType   uint   `form:"bonusType"`
}

func (bonusController) Reward(c *gin.Context) {
	var req RewardReq

	err := c.ShouldBind(&req)

	if err != nil {
		fmt.Println(err)
	}

	bonusModel := models.NewBonusRankModel()

	record := &models.BonusRank{}
	record.BonusType = req.BonusType
	record.AccountName = req.AccountName
	record.Operator = "test"
	err = bonusModel.Insert(record)

	if err != nil {
		fmt.Println(err)
	}

	// c.HTML(http.StatusOK, "main.tmpl", nil)
}

type RewardRecordsReq struct {
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
	AccountName string `form:"accountName"`
}

type RewardRecordsRsp struct {
}

func (bonusController) RewardRecords(c *gin.Context) {
	var req RewardRecordsReq

	err := c.ShouldBind(&req)

	if err != nil {
		fmt.Println(err)
	}
	bonusModel := models.NewBonusRankModel()

	sTime := time.Unix(req.StartTime, 0).UTC()
	eTime := time.Unix(req.EndTime, 0).UTC()

	records, err := bonusModel.Recrods(req.AccountName, sTime, eTime)

	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := json.Marshal(records)
	fmt.Println(string(bytes))

	resp := &RewardRecordsRsp{}

	c.HTML(http.StatusOK, "main.tmpl", resp)
}

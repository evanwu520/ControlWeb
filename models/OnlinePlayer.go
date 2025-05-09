package models

import (
	"context"
	"encoding/json"
	"fmt"

	"game.com/controlWeb/redis"
)

type onlinePlyarModel struct {
}

func NewOnlinePlyarModel() *onlinePlyarModel {
	return &onlinePlyarModel{}
}

func (onlinePlyarModel) OnlinePlayerList() {

	redisClient := redis.NewRedis().GetClient()

	// Fetch raw JSON string
	val, err := redisClient.HGetAll(context.Background(), "GM:10020000:1").Result()
	if err != nil {
		fmt.Println("Redis GET error:", err)
		return
	}

	type Player struct {
		GameUserSetting       any            `json:"GameUserSetting"`
		KickOutCount          int            `json:"KickOutCount"`
		AuthKey               string         `json:"AuthKey"`
		AccessToken           string         `json:"AccessToken"`
		TypeID                int            `json:"TypeID"`
		TypeCode              string         `json:"TypeCode"`
		AgentID               int64          `json:"AgentID"`
		AgentCode             string         `json:"AgentCode"`
		MemberID              int64          `json:"MemberID"`
		AccountName           string         `json:"AccountName"`
		MemberTypeID          int            `json:"MemberTypeID"`
		MemberTypeCode        string         `json:"MemberTypeCode"`
		SubagentID            int64          `json:"SubagentID"`
		SubagentCode          string         `json:"SubagentCode"`
		SubagentTypeID        int            `json:"SubagentTypeID"`
		SubagentTypeCode      string         `json:"SubagentTypeCode"`
		NickName              string         `json:"NickName"`
		Avatar                any            `json:"Avatar"`
		IsBetPermissionDenied bool           `json:"IsBetPermissionDenied"`
		IsTest                bool           `json:"IsTest"`
		BillDateTimeZone      int            `json:"BillDateTimeZone"`
		Tags                  []any          `json:"Tags"`
		SignInIP              string         `json:"SignInIP"`
		AllowCreditTypes      []int          `json:"AllowCreditTypes"`
		MainCreditTypeID      int            `json:"MainCreditTypeID"`
		MemberCredits         map[string]any `json:"MemberCredits"`
	}

	var players []*Player

	fmt.Println(val)
	bytes, _ := json.Marshal(val)

	json.Unmarshal(bytes, &players)

	fmt.Println(players)

}

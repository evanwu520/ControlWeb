package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"game.com/controlWeb/redis"
)

type onlinePlayersModel struct {
}

func NewOnlinePlayersModel() *onlinePlayersModel {
	return &onlinePlayersModel{}
}

type Player struct {
	IsLoaded      bool         `json:"IsLoaded"`
	Player        PlayerDetail `json:"Player"`
	Status        Status       `json:"Status"`
	GameID        int          `json:"GameID"`
	TableID       int          `json:"TableID"`
	MachineID     int          `json:"MachineID"`
	MachineName   string       `json:"MachineName"`
	ConfigID      int          `json:"ConfigID"`
	Configs       []any        `json:"Configs"` // use specific type if known
	Setting       any          `json:"Setting"` // can change to specific struct if not null
	IsEnable      bool         `json:"IsEnable"`
	BuildDateTime time.Time    `json:"BuildDateTime"`
}

type PlayerDetail struct {
	GameUserSetting       any      `json:"GameUserSetting"` // can change to specific struct if not null
	KickOutCount          int      `json:"KickOutCount"`
	AuthKey               string   `json:"AuthKey"`
	AccessToken           string   `json:"AccessToken"`
	TypeID                int      `json:"TypeID"`
	TypeCode              string   `json:"TypeCode"`
	AgentID               int64    `json:"AgentID"`
	AgentCode             string   `json:"AgentCode"`
	MemberID              int64    `json:"MemberID"`
	AccountName           string   `json:"AccountName"`
	MemberTypeID          int      `json:"MemberTypeID"`
	MemberTypeCode        string   `json:"MemberTypeCode"`
	SubagentID            int64    `json:"SubagentID"`
	SubagentCode          string   `json:"SubagentCode"`
	SubagentTypeID        int      `json:"SubagentTypeID"`
	SubagentTypeCode      string   `json:"SubagentTypeCode"`
	NickName              string   `json:"NickName"`
	Avatar                any      `json:"Avatar"` // nullable
	IsBetPermissionDenied bool     `json:"IsBetPermissionDenied"`
	IsTest                bool     `json:"IsTest"`
	BillDateTimeZone      int      `json:"BillDateTimeZone"`
	Tags                  []any    `json:"Tags"`
	SignInIP              string   `json:"SignInIP"`
	AllowCreditTypes      []int    `json:"AllowCreditTypes"`
	MainCreditTypeID      int      `json:"MainCreditTypeID"`
	MemberCredits         struct{} `json:"MemberCredits"` // empty object
}

type Status struct {
	ThirtydayWinList        []int64   `json:"ThirtydayWinList"`
	SeatState               int       `json:"SeatState"`
	RoundID                 int64     `json:"RoundID"`
	SeatSetProbRTP          string    `json:"SeatSetProbRTP"`
	LastBetRecord           int64     `json:"LastBetRecord"`
	TodayBet                int64     `json:"TodayBet"`
	TodayBetDateTime        time.Time `json:"TodayBetDateTime"`
	YesterdayBet            int64     `json:"YesterdayBet"`
	ThirtydayBetList        []int64   `json:"ThirtydayBetList"`
	ThirtydayBet            int64     `json:"ThirtydayBet"`
	TodayWin                int64     `json:"TodayWin"`
	YesterdayWin            int64     `json:"YesterdayWin"`
	ThirtydayWin            int64     `json:"ThirtydayWin"`
	NoneFreeCount           int       `json:"NoneFreeCount"`
	FristForwardFree        int       `json:"FristForwardFree"`
	SecondForwardFree       int       `json:"SecondForwardFree"`
	ExpirationTime          time.Time `json:"ExpirationTime"`
	KeepAliveExpirationTime time.Time `json:"KeepAliveExpirationTime"`
	CtrlExpirationTime      time.Time `json:"CtrlExpirationTime"`
	GameCtrlOption          int       `json:"GameCtrlOption"`
	GameCtrlAccount         string    `json:"GameCtrlAccount"`
	MachineSpinkey          string    `json:"MachineSpinkey"`
	GameProbUpdate          int       `json:"GameProbUpdate"`
	MaintenanceDateTime     time.Time `json:"MaintenanceDateTime"`
	MaintenanceState        int       `json:"MaintenanceState"`
	CurrencyRate            int       `json:"CurrencyRate"`
}

func (onlinePlayersModel) OnlinePlayerList() []*Player {

	redisClient := redis.NewRedis().GetClient()

	// Fetch raw JSON string
	val, err := redisClient.HGetAll(context.Background(), "GM:10020000:1").Result()
	if err != nil {
		fmt.Println("Redis GET error:", err)
		return nil
	}

	var players []*Player

	bytes, _ := json.Marshal(val)
	err = json.Unmarshal(bytes, &players)

	if err != nil {
		fmt.Println("Redis data Unmarshal:", err)
		return nil
	}

	return players

}

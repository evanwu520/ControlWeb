package models

import (
	"time"

	"game.com/controlWeb/db"
)

func NewBonusRankModel() *bonusRankModel {
	return &bonusRankModel{TableName: "BonusRank"}
}

type bonusRankModel struct {
	TableName string
}

type BonusRank struct {
	ID            uint      `gorm:"primaryKey;column:ID;->"` // DB column: "id"
	AccountName   string    `gorm:"column:AccountName"`
	DeskNo        uint      `gorm:"column:DeskNo"`
	BonusType     uint      `gorm:"column:BonusType"`
	Operator      string    `gorm:"column:Operator"`
	BuildDateTime time.Time `gorm:"column:BuildDateTime;->"` // DB column: "build_date_time"
}

func (model *bonusRankModel) Insert(data *BonusRank) error {

	conn := db.NewSql().GetConn()

	err := conn.Table(model.TableName).Create(data).Error

	if err != nil {
		return err
	}
	return nil
}

func (model *bonusRankModel) Recrods(accountName string, startTime, endTime time.Time) (records []*BonusRank, err error) {

	conn := db.NewSql().GetConn()
	result := conn.Table(model.TableName).
		Where("AccountName = ? And BuildDateTime Between ? and ? ",
			accountName, startTime, endTime).
		Find(&records).Order("ID Desc")

	err = result.Error

	return
}

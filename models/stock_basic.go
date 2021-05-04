package models

import (
	"time"
)

type StockBasic struct {
	Id         int       `json:"Id"`
	TsCode     string    `json:"TsCode"`
	Code       string    `json:"Code"`
	Name       string    `json:"Name"`
	Area       string    `json:"Area"`       //所在地域
	Industry   string    `json:"Industry"`   // 所属行业
	FullName   string    `json:"FullName"`   //股票全称
	EnName     string    `json:"EnName"`     //英文名称
	Market     string    `json:"Market"`     //市场类型 （主板/中小板/创业板/科创板/CDR）
	ExChange   string    `json:"ExChange"`   //交易所代码
	CurrType   string    `json:"CurrType"`   //交易货币
	ListStatus string    `json:"ListStatus"` //上市状态： L上市 D退市 P暂停上市
	ListDate   time.Time `json:"ListDate"`   //上市日期
	DeListDate time.Time `json:"DeListDate"` //退市日期
	IsHs       string    `json:"IsHs"`       //是否沪深港通标的，N否 H沪股通 S深股通
}

func (StockBasic) TableName() string {
	return "stock_basic"
}

func (s *StockBasic) Insert() error {
	db := dbConn.Create(&s)
	return db.Error
}

func (s *StockBasic) GetStockBasicById() bool {
	res := dbConn.Where("id = ?", s.Id).First(&s).RecordNotFound()
	return res
}

func (s *StockBasic) Update() error {
	return dbConn.Model(&s).Where("Id = ?", s.Id).Updates(map[string]interface{}{
		"ListDate":   s.ListDate,
		"DeListDate": s.DeListDate,
		"ListStatus": s.ListStatus,
	}).Error
}

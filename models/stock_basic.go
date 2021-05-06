package models

import (
	"time"
)

type StockBasic struct {
	Id         int        `gorm:"column:Id"`
	TsCode     string     `gorm:"column:TsCode"`
	Code       string     `gorm:"column:Code"`
	Name       string     `gorm:"column:Name"`
	Area       string     `gorm:"column:Area"`       //所在地域
	Industry   string     `gorm:"column:Industry"`   // 所属行业
	FullName   string     `gorm:"column:FullName"`   //股票全称
	CnSpell    string     `gorm:"column:CnSpell"`    //拼音缩写
	EnName     string     `gorm:"column:EnName"`     //英文名称
	Market     string     `gorm:"column:Market"`     //市场类型 （主板/中小板/创业板/科创板/CDR）
	ExChange   string     `gorm:"column:ExChange"`   //交易所代码
	CurrType   string     `gorm:"column:CurrType"`   //交易货币
	ListStatus string     `gorm:"column:ListStatus"` //上市状态： L上市 D退市 P暂停上市
	ListDate   *time.Time `gorm:"column:ListDate"`   //上市日期
	DeListDate *time.Time `gorm:"column:DeListDate"` //退市日期
	IsHs       string     `gorm:"column:IsHs"`       //是否沪深港通标的，N否 H沪股通 S深股通
}

func (StockBasic) TableName() string {
	return "stock_basic"
}

func (s *StockBasic) Insert() error {
	db := dbConn.Create(&s)
	return db.Error
}

func (s *StockBasic) GetStockBasicById() *StockBasic {
	dbConn.Where("id = ?", s.Id).First(&s)
	return s
}

func (s *StockBasic) GetStockBasicByCode() int {
	dbConn.Raw("SELECT Id FROM stock_basic WHERE code =?;", s.Code).Scan(&s)
	return s.Id
}

func (s *StockBasic) Update() error {
	return dbConn.Model(&s).Where("Id = ?", s.Id).Updates(map[string]interface{}{
		"ListDate":   s.ListDate,
		"DeListDate": s.DeListDate,
		"ListStatus": s.ListStatus,
	}).Error
}

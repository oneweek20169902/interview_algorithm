package models

type StockDaily struct {
	Id             int     `json:"Id"`             //
	Code           string  `json:"Code"`           //
	LatestPrice    float64 `json:"LatestPrice"`    //最新价
	UpOrDown       int     `json:"UpOrDown"`       //涨跌幅
	UpOrDownAmount float64 `json:"UpOrDownAmount"` //涨跌额
	Volume         float64 `json:"Volume"`         //成交量
	Amplitude      float64 `json:"Amplitude"`      //振幅
	Turnover       float64 `json:"Turnover"`       //成交额
	YesterdayPrice float64 `json:"YesterdayPrice"` //昨日收盘价
	TodayOpenPrice float64 `json:"TodayOpenPrice"` //今日开盘价
	HighestPrice   float64 `json:"HighestPrice"`   //
	LowestPrice    float64 `json:"LowestPrice"`    //
	Exchange       float64 `json:"Exchange"`       //换手率
	VolumeRatio    float64 `json:"VolumeRatio"`    //量比
	PE             float64 `json:"PE"`             //市盈率
	MRQ            float64 `json:"MRQ"`            //每股净资产
	EPS            float64 `json:"EPS"`            //每股收益
	ROE            float64 `json:"ROE"`            //净资产收益率
	PB             float64 `json:"PB"`             //市净率
	GMV            float64 `json:"GMV"`            //总市值
	CMV            float64 `json:"CMV"`            //流通市值
	Industry       string  `json:"Industry"`       //所属行业
	CreateDate     float64 `json:"CreateDate"`
	ModifyDate     float64 `json:"ModifyDate"`
}

func (StockDaily) TableName() string {
	return "stock_daily"
}

func (s *StockDaily) Insert() error {
	db := dbConn.Create(&s)
	return db.Error
}

func (s *StockDaily) GetStockDailyById() bool {
	return dbConn.Where(" Id=?", s.Id).First(&s).RecordNotFound()
}

func (s *StockDaily) GetStockDailyByCode() bool {
	return dbConn.Where(" Code=?", s.Code).First(&s).RecordNotFound()
}

func (s *StockDaily) Update() error {
	return dbConn.Where(" Id = ?", s.Id).Updates(map[string]interface{}{
		"PE":             s.PE,
		"MRQ":            s.MRQ,
		"EPS":            s.EPS,
		"ROE":            s.ROE,
		"PB":             s.PB,
		"GMV":            s.GMV,
		"TodayOpenPrice": s.TodayOpenPrice,
		"YesterdayPrice": s.YesterdayPrice,
		"Turnover":       s.Turnover,
		"Volume":         s.Volume,
		"UpOrDownAmount": s.UpOrDownAmount,
		"UpOrDown":       s.UpOrDown,
		"LatestPrice":    s.LatestPrice,
		"HighestPrice":   s.HighestPrice,
		"LowestPrice":    s.LowestPrice,
		"CreateDate":     s.CreateDate,
		"ModifyDate":     s.ModifyDate,
	}).Error

}

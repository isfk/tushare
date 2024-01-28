// Code generated by protoc-gen-go-tushare. DO NOT EDIT.
// source: api/v1/us.proto

package v1

import (
	json "encoding/json"
)

var _ = json.Decoder{}

// api name

// 美股列表 api name
const ApiUsBasic string = "us_basic"

// 美股行情 api name
const ApiUsDaily string = "us_daily"

// 美股交易日历 api name
const ApiUsTradecal string = "us_tradecal"

// fields

// 美股列表 fields
var FieldsUsBasic = []string{"ts_code", "name", "enname", "classify", "list_date", "delist_date"}

// 美股行情 fields
var FieldsUsDaily = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_change", "vol", "amount", "vwap", "turnover_ratio", "total_mv", "pe", "pb"}

// 美股交易日历 fields
var FieldsUsTradecal = []string{"cal_date", "is_open", "pretrade_date"}

// struct

// 美股列表|us_basic
type UsBasic struct {
	TsCode     string `json:"ts_code"`     // 美股代码
	Name       string `json:"name"`        // 中文名称
	Enname     string `json:"enname"`      // 英文名称
	Classify   string `json:"classify"`    // 分类ADR/GDR/EQ
	ListDate   string `json:"list_date"`   // 上市日期
	DelistDate string `json:"delist_date"` // 退市日期
}

type UsBasicRequest struct {
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
	TsCode   string `json:"ts_code"`  // 股票代码, eg.AAPL（苹果）
	Classify string `json:"classify"` // 股票分类, ADR/GDR/EQ
}

type UsBasicResponse struct {
	List []*UsBasic `json:"list"`
}

func (x *UsBasicResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 美股行情|us_daily
type UsDaily struct {
	TsCode        string  `json:"ts_code"`        // 股票代码
	TradeDate     string  `json:"trade_date"`     // 交易日期
	Close         float64 `json:"close"`          // 收盘价
	Open          float64 `json:"open"`           // 开盘价
	High          float64 `json:"high"`           // 最高价
	Low           float64 `json:"low"`            // 最低价
	PreClose      float64 `json:"pre_close"`      // 昨收价
	Change        float64 `json:"change"`         // 涨跌额
	PctChange     float64 `json:"pct_change"`     // 涨跌幅
	Vol           float64 `json:"vol"`            // 成交量
	Amount        float64 `json:"amount"`         // 成交额
	Vwap          float64 `json:"vwap"`           // 平均价
	TurnoverRatio float64 `json:"turnover_ratio"` // 换手率
	TotalMv       float64 `json:"total_mv"`       // 总市值
	Pe            float64 `json:"pe"`             // PE
	Pb            float64 `json:"pb"`             // PB
}

type UsDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码（e.g. AAPL）
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD）
	StartDate string `json:"start_date"` // 开始日期（YYYYMMDD）
	EndDate   string `json:"end_date"`   // 结束日期（YYYYMMDD）
}

type UsDailyResponse struct {
	List []*UsDaily `json:"list"`
}

func (x *UsDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 美股交易日历|us_tradecal
type UsTradecal struct {
	CalDate      string `json:"cal_date"`      // 日历日期
	IsOpen       int64  `json:"is_open"`       // 是否交易 '0'休市 '1'交易
	PretradeDate string `json:"pretrade_date"` // 上一个交易日
}

type UsTradecalRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
	IsOpen    string `json:"is_open"`    // 是否交易
}

type UsTradecalResponse struct {
	List []*UsTradecal `json:"list"`
}

func (x *UsTradecalResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

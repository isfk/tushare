// Code generated by protoc-gen-go-tushare. DO NOT EDIT.
// source: api/v1/cb.proto

package v1

import (
	json "encoding/json"
)

var _ = json.Decoder{}

// api name

// 债券大宗交易 api name
const ApiBondBlk string = "bond_blk"

// 大宗交易明细 api name
const ApiBondBlkDetail string = "bond_blk_detail"

// 可转债基本信息 api name
const ApiCbBasic string = "cb_basic"

// 可转债赎回信息 api name
const ApiCbCall string = "cb_call"

// 可转债行情 api name
const ApiCbDaily string = "cb_daily"

// 可转债发行 api name
const ApiCbIssue string = "cb_issue"

// 可转债转股价变动 api name
const ApiCbPriceChg string = "cb_price_chg"

// 可转债票面利率 api name
const ApiCbRate string = "cb_rate"

// 可转债转股结果 api name
const ApiCbShare string = "cb_share"

// 财经日历 api name
const ApiEcoCal string = "eco_cal"

// 债券回购日行情 api name
const ApiRepoDaily string = "repo_daily"

// 国债收益率曲线 api name
const ApiYcCb string = "yc_cb"

// fields

// 债券大宗交易 fields
var FieldsBondBlk = []string{"trade_date", "ts_code", "name", "price", "vol", "amount"}

// 大宗交易明细 fields
var FieldsBondBlkDetail = []string{"trade_date", "ts_code", "name", "price", "vol", "amount", "buy_dp", "sell_dp"}

// 可转债基本信息 fields
var FieldsCbBasic = []string{"ts_code", "bond_full_name", "bond_short_name", "cb_code", "stk_code", "stk_short_name", "maturity", "par", "issue_price", "issue_size", "remain_size", "value_date", "maturity_date", "rate_type", "coupon_rate", "add_rate", "pay_per_year", "list_date", "delist_date", "exchange", "conv_start_date", "conv_end_date", "conv_stop_date", "first_conv_price", "conv_price", "rate_clause", "put_clause", "maturity_put_price", "call_clause", "reset_clause", "conv_clause", "guarantor", "guarantee_type", "issue_rating", "newest_rating", "rating_comp"}

// 可转债赎回信息 fields
var FieldsCbCall = []string{"ts_code", "call_type", "is_call", "ann_date", "call_date", "call_price", "call_price_tax", "call_vol", "call_amount", "payment_date", "call_reg_date"}

// 可转债行情 fields
var FieldsCbDaily = []string{"ts_code", "trade_date", "pre_close", "open", "high", "low", "close", "change", "pct_chg", "vol", "amount", "bond_value", "bond_over_rate", "cb_value", "cb_over_rate"}

// 可转债发行 fields
var FieldsCbIssue = []string{"ts_code", "ann_date", "res_ann_date", "plan_issue_size", "issue_size", "issue_price", "issue_type", "issue_cost", "onl_code", "onl_name", "onl_date", "onl_size", "onl_pch_vol", "onl_pch_num", "onl_pch_excess", "onl_winning_rate", "shd_ration_code", "shd_ration_name", "shd_ration_date", "shd_ration_record_date", "shd_ration_pay_date", "shd_ration_price", "shd_ration_ratio", "shd_ration_size", "shd_ration_vol", "shd_ration_num", "shd_ration_excess", "offl_size", "offl_deposit", "offl_pch_vol", "offl_pch_num", "offl_pch_excess", "offl_winning_rate", "lead_underwriter", "lead_underwriter_vol"}

// 可转债转股价变动 fields
var FieldsCbPriceChg = []string{"ts_code", "bond_short_name", "publish_date", "change_date", "convert_price_initial", "convertprice_bef", "convertprice_aft"}

// 可转债票面利率 fields
var FieldsCbRate = []string{"ts_code", "rate_freq", "rate_start_date", "rate_end_date", "coupon_rate"}

// 可转债转股结果 fields
var FieldsCbShare = []string{"ts_code", "bond_short_name", "publish_date", "end_date", "issue_size", "convert_price_initial", "convert_price", "convert_val", "convert_vol", "convert_ratio", "acc_convert_val", "acc_convert_vol", "acc_convert_ratio", "remain_size", "total_shares"}

// 财经日历 fields
var FieldsEcoCal = []string{"date", "time", "currency", "country", "event", "value", "pre_value", "fore_value"}

// 债券回购日行情 fields
var FieldsRepoDaily = []string{"ts_code", "trade_date", "repo_maturity", "pre_close", "open", "high", "low", "close", "weight", "weight_r", "amount", "num"}

// 国债收益率曲线 fields
var FieldsYcCb = []string{"trade_date", "ts_code", "curve_name", "curve_type", "curve_term", "yield"}

// struct

// 债券大宗交易|bond_blk
type BondBlk struct {
	TradeDate string  `json:"trade_date"` // 交易日期
	TsCode    string  `json:"ts_code"`    // 债券代码
	Name      string  `json:"name"`       // 债券名称
	Price     float64 `json:"price"`      // 成交价（元）
	Vol       float64 `json:"vol"`        // 累计成交数量（万股/万份/万张/万手）
	Amount    float64 `json:"amount"`     // 累计成交金额（万元）
}

type BondBlkRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 债券代码
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD格式，下同）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type BondBlkResponse struct {
	List []*BondBlk `json:"list"`
}

func (x *BondBlkResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 大宗交易明细|bond_blk_detail
type BondBlkDetail struct {
	TradeDate string  `json:"trade_date"` // 交易日期
	TsCode    string  `json:"ts_code"`    // 债券代码
	Name      string  `json:"name"`       // 债券名称
	Price     float64 `json:"price"`      // 成交价（元）
	Vol       float64 `json:"vol"`        // 成交数量（万股/万份/万张/万手）
	Amount    float64 `json:"amount"`     // 成交金额（万元）
	BuyDp     string  `json:"buy_dp"`     // 买方营业部
	SellDp    string  `json:"sell_dp"`    // 卖方营业部
}

type BondBlkDetailRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 债券代码
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD格式，下同）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type BondBlkDetailResponse struct {
	List []*BondBlkDetail `json:"list"`
}

func (x *BondBlkDetailResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 可转债基本信息|cb_basic
type CbBasic struct {
	TsCode           string  `json:"ts_code"`            // 转债代码
	BondFullName     string  `json:"bond_full_name"`     // 转债名称
	BondShortName    string  `json:"bond_short_name"`    // 转债简称
	CbCode           string  `json:"cb_code"`            // 转股申报代码
	StkCode          string  `json:"stk_code"`           // 正股代码
	StkShortName     string  `json:"stk_short_name"`     // 正股简称
	Maturity         float64 `json:"maturity"`           // 发行期限（年）
	Par              float64 `json:"par"`                // 面值
	IssuePrice       float64 `json:"issue_price"`        // 发行价格
	IssueSize        float64 `json:"issue_size"`         // 发行总额（元）
	RemainSize       float64 `json:"remain_size"`        // 债券余额（元）
	ValueDate        string  `json:"value_date"`         // 起息日期
	MaturityDate     string  `json:"maturity_date"`      // 到期日期
	RateType         string  `json:"rate_type"`          // 利率类型
	CouponRate       float64 `json:"coupon_rate"`        // 票面利率（%）
	AddRate          float64 `json:"add_rate"`           // 补偿利率（%）
	PayPerYear       int64   `json:"pay_per_year"`       // 年付息次数
	ListDate         string  `json:"list_date"`          // 上市日期
	DelistDate       string  `json:"delist_date"`        // 摘牌日
	Exchange         string  `json:"exchange"`           // 上市地点
	ConvStartDate    string  `json:"conv_start_date"`    // 转股起始日
	ConvEndDate      string  `json:"conv_end_date"`      // 转股截止日
	ConvStopDate     string  `json:"conv_stop_date"`     // 停止转股日(提前到期)
	FirstConvPrice   float64 `json:"first_conv_price"`   // 初始转股价
	ConvPrice        float64 `json:"conv_price"`         // 最新转股价
	RateClause       string  `json:"rate_clause"`        // 利率说明
	PutClause        string  `json:"put_clause"`         // 赎回条款
	MaturityPutPrice float64 `json:"maturity_put_price"` // 到期赎回价格(含税)
	CallClause       string  `json:"call_clause"`        // 回售条款
	ResetClause      string  `json:"reset_clause"`       // 特别向下修正条款
	ConvClause       string  `json:"conv_clause"`        // 转股条款
	Guarantor        string  `json:"guarantor"`          // 担保人
	GuaranteeType    string  `json:"guarantee_type"`     // 担保方式
	IssueRating      string  `json:"issue_rating"`       // 发行信用等级
	NewestRating     string  `json:"newest_rating"`      // 最新信用等级
	RatingComp       string  `json:"rating_comp"`        // 最新评级机构
}

type CbBasicRequest struct {
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
	TsCode   string `json:"ts_code"`   // 转债代码
	ListDate string `json:"list_date"` // 上市日期
	Exchange string `json:"exchange"`  // 上市地点
}

type CbBasicResponse struct {
	List []*CbBasic `json:"list"`
}

func (x *CbBasicResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 可转债赎回信息|cb_call
type CbCall struct {
	TsCode       string  `json:"ts_code"`        // 转债代码
	CallType     string  `json:"call_type"`      // 赎回类型：到赎、强赎
	IsCall       string  `json:"is_call"`        // 是否赎回：已满足强赎条件、公告提示强赎、公告实施强赎、公告到期赎回、公告不强赎
	AnnDate      string  `json:"ann_date"`       // 公告/提示日期
	CallDate     string  `json:"call_date"`      // 赎回日期
	CallPrice    float64 `json:"call_price"`     // 赎回价格(含税，元/张)
	CallPriceTax float64 `json:"call_price_tax"` // 赎回价格(扣税，元/张)
	CallVol      float64 `json:"call_vol"`       // 赎回债券数量(张)
	CallAmount   float64 `json:"call_amount"`    // 赎回金额(万元)
	PaymentDate  string  `json:"payment_date"`   // 行权后款项到账日
	CallRegDate  string  `json:"call_reg_date"`  // 赎回登记日
}

type CbCallRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	TsCode string `json:"ts_code"` // 转债代码，支持多值输入
}

type CbCallResponse struct {
	List []*CbCall `json:"list"`
}

func (x *CbCallResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 可转债行情|cb_daily
type CbDaily struct {
	TsCode       string  `json:"ts_code"`        // 转债代码
	TradeDate    string  `json:"trade_date"`     // 交易日期
	PreClose     float64 `json:"pre_close"`      // 昨收盘价(元)
	Open         float64 `json:"open"`           // 开盘价(元)
	High         float64 `json:"high"`           // 最高价(元)
	Low          float64 `json:"low"`            // 最低价(元)
	Close        float64 `json:"close"`          // 收盘价(元)
	Change       float64 `json:"change"`         // 涨跌(元)
	PctChg       float64 `json:"pct_chg"`        // 涨跌幅(%)
	Vol          float64 `json:"vol"`            // 成交量(手)
	Amount       float64 `json:"amount"`         // 成交金额(万元)
	BondValue    float64 `json:"bond_value"`     // 纯债价值
	BondOverRate float64 `json:"bond_over_rate"` // 纯债溢价率(%)
	CbValue      float64 `json:"cb_value"`       // 转股价值
	CbOverRate   float64 `json:"cb_over_rate"`   // 转股溢价率(%)
}

type CbDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	TradeDate string `json:"trade_date"` // 交易日期(YYYYMMDD格式，下同)
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type CbDailyResponse struct {
	List []*CbDaily `json:"list"`
}

func (x *CbDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 可转债发行|cb_issue
type CbIssue struct {
	TsCode              string  `json:"ts_code"`                // 转债代码
	AnnDate             string  `json:"ann_date"`               // 发行公告日
	ResAnnDate          string  `json:"res_ann_date"`           // 发行结果公告日
	PlanIssueSize       float64 `json:"plan_issue_size"`        // 计划发行总额（元）
	IssueSize           float64 `json:"issue_size"`             // 发行总额（元）
	IssuePrice          float64 `json:"issue_price"`            // 发行价格
	IssueType           string  `json:"issue_type"`             // 发行方式
	IssueCost           float64 `json:"issue_cost"`             // 发行费用（元）
	OnlCode             string  `json:"onl_code"`               // 网上申购代码
	OnlName             string  `json:"onl_name"`               // 网上申购简称
	OnlDate             string  `json:"onl_date"`               // 网上发行日期
	OnlSize             float64 `json:"onl_size"`               // 网上发行总额（张）
	OnlPchVol           float64 `json:"onl_pch_vol"`            // 网上发行有效申购数量（张）
	OnlPchNum           int64   `json:"onl_pch_num"`            // 网上发行有效申购户数
	OnlPchExcess        float64 `json:"onl_pch_excess"`         // 网上发行超额认购倍数
	OnlWinningRate      float64 `json:"onl_winning_rate"`       // 网上发行中签率（%）
	ShdRationCode       string  `json:"shd_ration_code"`        // 老股东配售代码
	ShdRationName       string  `json:"shd_ration_name"`        // 老股东配售简称
	ShdRationDate       string  `json:"shd_ration_date"`        // 老股东配售日
	ShdRationRecordDate string  `json:"shd_ration_record_date"` // 老股东配售股权登记日
	ShdRationPayDate    string  `json:"shd_ration_pay_date"`    // 老股东配售缴款日
	ShdRationPrice      float64 `json:"shd_ration_price"`       // 老股东配售价格
	ShdRationRatio      float64 `json:"shd_ration_ratio"`       // 老股东配售比例
	ShdRationSize       float64 `json:"shd_ration_size"`        // 老股东配售数量（张）
	ShdRationVol        float64 `json:"shd_ration_vol"`         // 老股东配售有效申购数量（张）
	ShdRationNum        int64   `json:"shd_ration_num"`         // 老股东配售有效申购户数
	ShdRationExcess     float64 `json:"shd_ration_excess"`      // 老股东配售超额认购倍数
	OfflSize            float64 `json:"offl_size"`              // 网下发行总额（张）
	OfflDeposit         float64 `json:"offl_deposit"`           // 网下发行定金比例（%）
	OfflPchVol          float64 `json:"offl_pch_vol"`           // 网下发行有效申购数量（张）
	OfflPchNum          int64   `json:"offl_pch_num"`           // 网下发行有效申购户数
	OfflPchExcess       float64 `json:"offl_pch_excess"`        // 网下发行超额认购倍数
	OfflWinningRate     float64 `json:"offl_winning_rate"`      // 网下发行中签率
	LeadUnderwriter     string  `json:"lead_underwriter"`       // 主承销商
	LeadUnderwriterVol  float64 `json:"lead_underwriter_vol"`   // 主承销商包销数量（张）
}

type CbIssueRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	AnnDate   string `json:"ann_date"`   // 发行公告日
	StartDate string `json:"start_date"` // 公告开始日期
	EndDate   string `json:"end_date"`   // 公告结束日期
}

type CbIssueResponse struct {
	List []*CbIssue `json:"list"`
}

func (x *CbIssueResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 可转债转股价变动|cb_price_chg
type CbPriceChg struct {
	TsCode              string  `json:"ts_code"`               // 转债代码
	BondShortName       string  `json:"bond_short_name"`       // 转债简称
	PublishDate         string  `json:"publish_date"`          // 公告日期
	ChangeDate          string  `json:"change_date"`           // 变动日期
	ConvertPriceInitial float64 `json:"convert_price_initial"` // 初始转股价格
	ConvertpriceBef     float64 `json:"convertprice_bef"`      // 修正前转股价格
	ConvertpriceAft     float64 `json:"convertprice_aft"`      // 修正后转股价格
}

type CbPriceChgRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	TsCode string `json:"ts_code"` // 转债代码，支持多值输入
}

type CbPriceChgResponse struct {
	List []*CbPriceChg `json:"list"`
}

func (x *CbPriceChgResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 可转债票面利率|cb_rate
type CbRate struct {
	TsCode        string  `json:"ts_code"`         // 转债代码
	RateFreq      int64   `json:"rate_freq"`       // 付息频率(次/年)
	RateStartDate string  `json:"rate_start_date"` // 付息开始日期
	RateEndDate   string  `json:"rate_end_date"`   // 付息结束日期
	CouponRate    float64 `json:"coupon_rate"`     // 票面利率(%)
}

type CbRateRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	TsCode string `json:"ts_code"` // 转债代码，支持多值输入
}

type CbRateResponse struct {
	List []*CbRate `json:"list"`
}

func (x *CbRateResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 可转债转股结果|cb_share
type CbShare struct {
	TsCode              string  `json:"ts_code"`               // 债券代码
	BondShortName       string  `json:"bond_short_name"`       // 债券简称
	PublishDate         string  `json:"publish_date"`          // 公告日期
	EndDate             string  `json:"end_date"`              // 统计截止日期
	IssueSize           float64 `json:"issue_size"`            // 可转债发行总额
	ConvertPriceInitial float64 `json:"convert_price_initial"` // 初始转换价格
	ConvertPrice        float64 `json:"convert_price"`         // 本次转换价格
	ConvertVal          float64 `json:"convert_val"`           // 本次转股金额
	ConvertVol          float64 `json:"convert_vol"`           // 本次转股数量
	ConvertRatio        float64 `json:"convert_ratio"`         // 本次转股比例
	AccConvertVal       float64 `json:"acc_convert_val"`       // 累计转股金额
	AccConvertVol       float64 `json:"acc_convert_vol"`       // 累计转股数量
	AccConvertRatio     float64 `json:"acc_convert_ratio"`     // 累计转股比例
	RemainSize          float64 `json:"remain_size"`           // 可转债剩余金额
	TotalShares         float64 `json:"total_shares"`          // 转股后总股本
}

type CbShareRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	TsCode string `json:"ts_code"` // 转债代码，支持多值输入
}

type CbShareResponse struct {
	List []*CbShare `json:"list"`
}

func (x *CbShareResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 财经日历|eco_cal
type EcoCal struct {
	Date      string `json:"date"`       // 日期
	Time      string `json:"time"`       // 时间
	Currency  string `json:"currency"`   // 货币代码
	Country   string `json:"country"`    // 国家
	Event     string `json:"event"`      // 经济事件
	Value     string `json:"value"`      // 今值
	PreValue  string `json:"pre_value"`  // 前值
	ForeValue string `json:"fore_value"` // 预测值
}

type EcoCalRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	Date      string `json:"date"`       // 日期（YYYYMMDD格式）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
	Currency  string `json:"currency"`   // 货币代码
	Country   string `json:"country"`    // 国家（比如：中国、美国）
	Event     string `json:"event"`      // 事件 （支持模糊匹配： *非农*）
}

type EcoCalResponse struct {
	List []*EcoCal `json:"list"`
}

func (x *EcoCalResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 债券回购日行情|repo_daily
type RepoDaily struct {
	TsCode       string  `json:"ts_code"`       // TS代码
	TradeDate    string  `json:"trade_date"`    // 交易日期
	RepoMaturity string  `json:"repo_maturity"` // 期限品种
	PreClose     float64 `json:"pre_close"`     // 前收盘(%)
	Open         float64 `json:"open"`          // 开盘价(%)
	High         float64 `json:"high"`          // 最高价(%)
	Low          float64 `json:"low"`           // 最低价(%)
	Close        float64 `json:"close"`         // 收盘价(%)
	Weight       float64 `json:"weight"`        // 加权价(%)
	WeightR      float64 `json:"weight_r"`      // 加权价(利率债)(%)
	Amount       float64 `json:"amount"`        // 成交金额(万元)
	Num          int64   `json:"num"`           // 成交笔数(笔)
}

type RepoDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	TradeDate string `json:"trade_date"` // 交易日期(YYYYMMDD格式，下同)
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type RepoDailyResponse struct {
	List []*RepoDaily `json:"list"`
}

func (x *RepoDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 国债收益率曲线|yc_cb
type YcCb struct {
	TradeDate string  `json:"trade_date"` // 交易日期
	TsCode    string  `json:"ts_code"`    // 曲线编码
	CurveName string  `json:"curve_name"` // 曲线名称
	CurveType string  `json:"curve_type"` // 曲线类型：0-到期，1-即期
	CurveTerm float64 `json:"curve_term"` // 期限(年)
	Yield     float64 `json:"yield"`      // 收益率(%)
}

type YcCbRequest struct {
	Limit     string  `json:"limit"`
	Offset    string  `json:"offset"`
	TsCode    string  `json:"ts_code"`    // 收益率曲线编码：1001.CB-国债收益率曲线
	CurveType string  `json:"curve_type"` // 曲线类型：0-到期，1-即期
	TradeDate string  `json:"trade_date"` // 交易日期
	StartDate string  `json:"start_date"` // 查询起始日期
	EndDate   string  `json:"end_date"`   // 查询结束日期
	CurveTerm float64 `json:"curve_term"` // 期限
}

type YcCbResponse struct {
	List []*YcCb `json:"list"`
}

func (x *YcCbResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

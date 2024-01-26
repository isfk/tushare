package tushare

// 这里定义的是接口的 `api_name`

const (

	/* 沪深股票 基础数据 */

	ApiStockBasic   = "stock_basic"
	ApiTradeCal     = "trade_cal"
	ApiNameChange   = "namechange"
	ApiHsConst      = "hs_const"
	ApiStockCompany = "stock_company"
	ApiStkManagers  = "stk_managers"
	ApiStkRewards   = "stk_rewards"
	ApiNewShare     = "new_share"
	ApiBakBasic     = "bak_basic"

	/* 沪深股票 行情数据 */

	ApiDaily         = "daily"
	ApiWeekly        = "weekly"
	ApiMonthly       = "monthly"
	ApiAdjFactor     = "adj_factor"
	ApiDailyBasic    = "daily_basic"
	ApiMoneyFlow     = "moneyflow"
	ApiStkLimit      = "stk_limit"
	ApiSuspendd      = "suspend_d"
	ApiMoneyFlowHsgt = "moneyflow_hsgt"
	ApiHsgtTop10     = "hsgt_top10"
	ApiGgtTop10      = "ggt_top10"
	ApiGgtDaily      = "ggt_daily"
	ApiGgtMonthly    = "ggt_monthly"
	ApiBakDaily      = "bak_daily"
)

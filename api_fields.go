package tushare

// 这里定义了每个接口的全部字段, 如果需要, 可以使用

var (
	/* 沪深股票 基础数据 */

	FieldsStockBasic   = []string{"ts_code", "symbol", "name", "area", "industry", "fullname", "enname", "cnspell", "market", "exchange", "curr_type", "list_status", "list_date", "delist_date", "is_hs", "act_name", "act_ent_type"}
	FieldsTradeCal     = []string{"exchange", "cal_date", "is_open", "pretrade_date"}
	FieldsNameChange   = []string{"ts_code", "name", "start_date", "end_date", "ann_date", "change_reason"}
	FieldsHsConst      = []string{"hs_type", "is_new"}
	FieldsStockCompany = []string{"ts_code", "exchange", "chairman", "manager", "secretary", "reg_capital", "setup_date", "province", "city", "introduction", "website", "email", "office", "ann_date", "business_scope", "employees", "main_business"}
	FieldsStkManagers  = []string{"ts_code", "ann_date", "name", "gender", "lev", "title", "edu", "national", "birthday", "begin_date", "end_date", "resume"}
	FieldsStkRewards   = []string{"ts_code", "ann_date", "end_date", "name", "title", "reward", "hold_vol"}
	FieldsNewShare     = []string{"ts_code", "sub_code", "name", "ipo_date", "issue_date", "amount", "market_amount", "price", "pe", "limit_amount", "funds", "ballot"}
	FieldsBakBasic     = []string{"trade_date", "ts_code", "name", "industry", "area", "pe", "float_share", "total_share", "total_assets", "liquid_assets", "fixed_assets", "reserved", "reserved_pershare", "eps", "bvps", "pb", "list_date", "undp", "per_undp", "rev_yoy", "profit_yoy", "gpr", "npr", "holder_num"}

	/* 沪深股票 行情数据 */

	FieldsDaily         = []string{"ts_code", "trade_date", "open", "high", "low", "close", "pre_close", "change", "pct_chg", "vol", "amount"}
	FieldsWeekly        = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"}
	FieldsMonthly       = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"}
	FieldsAdjFactor     = []string{"ts_code", "trade_date", "adj_factor"}
	FieldsDailyBasic    = []string{"ts_code", "trade_date", "close", "turnover_rate", "turnover_rate_f", "volume_ratio", "pe", "pe_ttm", "pb", "ps", "ps_ttm", "dv_ratio", "dv_ttm", "total_share", "float_share", "free_share", "total_mv", "circ_mv", "limit_status"}
	FieldsMoneyFlow     = []string{"ts_code", "trade_date", "buy_sm_vol", "buy_sm_amount", "sell_sm_vol", "sell_sm_amount", "buy_md_vol", "buy_md_amount", "sell_md_vol", "sell_md_amount", "buy_lg_vol", "buy_lg_amount", "sell_lg_vol", "sell_lg_amount", "buy_elg_vol", "buy_elg_amount", "sell_elg_vol", "sell_elg_amount", "net_mf_vol", "net_mf_amount", "trade_count"}
	FieldsStkLimit      = []string{"trade_date", "ts_code", "up_limit", "down_limit", "pre_close"}
	FieldsSuspendd      = []string{"ts_code", "trade_date", "suspend_timing", "suspend_type"}
	FieldsMoneyFlowHsgt = []string{"trade_date", "ggt_ss", "ggt_sz", "hgt", "sgt", "north_money", "south_money"}
	FieldsHsgtTop10     = []string{"trade_date", "ts_code", "name", "close", "change", "rank", "market_type", "amount", "net_amount", "buy", "sell"}
	FieldsGgtTop10      = []string{"trade_date", "ts_code", "name", "close", "p_change", "rank", "market_type", "amount", "net_amount", "sh_amount", "sh_net_amount", "sh_buy", "sh_sell", "sz_amount", "sz_net_amount", "sz_buy", "sz_sell"}
	FieldsGgtDaily      = []string{"trade_date", "buy_amount", "buy_volume", "sell_amount", "sell_volume"}
	FieldsGgtMonthly    = []string{"month", "day_buy_amt", "day_buy_vol", "day_sell_amt", "day_sell_vol", "total_buy_amt", "total_buy_vol", "total_sell_amt", "total_sell_vol"}
	FieldsBakDaily      = []string{"ts_code", "trade_date", "name", "pct_change", "close", "change", "open", "high", "low", "pre_close", "vol_ratio", "turn_over", "swing", "vol", "amount", "selling", "buying", "total_share", "float_share", "pe", "industry", "area", "float_mv", "total_mv", "avg_price", "strength", "activity", "avg_turnover", "attack", "interval_3", "interval_6"}
)

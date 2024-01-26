package tushare

import (
	"os"
	"strings"
	"testing"
	"time"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestStockBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.StockBasicRequest, v1.StockBasicResponse](ts, &Request[v1.StockBasicRequest]{
		Params:  &v1.StockBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiStockBasic,
		Fields:  v1.FieldsStockBasic,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestTradeCal(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.TradeCalRequest, v1.TradeCalResponse](ts, &Request[v1.TradeCalRequest]{
		Params:  &v1.TradeCalRequest{CalDate: TushareDate(time.Now().AddDate(0, 0, -1)), Limit: "2", Offset: "0"},
		ApiName: v1.ApiTradeCal,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestNameChange(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.NameChangeRequest, v1.NameChangeResponse](ts, &Request[v1.NameChangeRequest]{
		Params:  &v1.NameChangeRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiNameChange,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestHsConst(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.HsConstRequest, v1.HsConstResponse](ts, &Request[v1.HsConstRequest]{
		Params:  &v1.HsConstRequest{Limit: "2", Offset: "0", HsType: "SH"},
		ApiName: v1.ApiHsConst,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestStockCompany(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.StockCompanyRequest, v1.StockCompanyResponse](ts, &Request[v1.StockCompanyRequest]{
		Params:  &v1.StockCompanyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiStockCompany,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestStkManagers(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.StkManagersRequest, v1.StkManagersResponse](ts, &Request[v1.StkManagersRequest]{
		Params:  &v1.StkManagersRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiStkManagers,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestStkRewards(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.StkRewardsRequest, v1.StkRewardsResponse](ts, &Request[v1.StkRewardsRequest]{
		Params:  &v1.StkRewardsRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiStkRewards,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestNewShare(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.NewShareRequest, v1.NewShareResponse](ts, &Request[v1.NewShareRequest]{
		Params:  &v1.NewShareRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiNewShare,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestBakBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.BakBasicRequest, v1.BakBasicResponse](ts, &Request[v1.BakBasicRequest]{
		Params:  &v1.BakBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiBakBasic,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.DailyRequest, v1.DailyResponse](ts, &Request[v1.DailyRequest]{
		Params:  &v1.DailyRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiDaily,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestWeekly(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.WeeklyRequest, v1.WeeklyResponse](ts, &Request[v1.WeeklyRequest]{
		Params:  &v1.WeeklyRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiWeekly,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestMonthly(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.MonthlyRequest, v1.MonthlyResponse](ts, &Request[v1.MonthlyRequest]{
		Params:  &v1.MonthlyRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiMonthly,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestAdjFactor(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.AdjFactorRequest, v1.AdjFactorResponse](ts, &Request[v1.AdjFactorRequest]{
		Params:  &v1.AdjFactorRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiAdjFactor,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestDailyBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.DailyBasicRequest, v1.DailyBasicResponse](ts, &Request[v1.DailyBasicRequest]{
		Params:  &v1.DailyBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiDailyBasic,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestMoneyFlow(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.MoneyFlowRequest, v1.MoneyFlowResponse](ts, &Request[v1.MoneyFlowRequest]{
		Params:  &v1.MoneyFlowRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiMoneyFlow,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestStkLimit(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.StkLimitRequest, v1.StkLimitResponse](ts, &Request[v1.StkLimitRequest]{
		Params:  &v1.StkLimitRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiStkLimit,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestSuspendd(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.SuspenddRequest, v1.SuspenddResponse](ts, &Request[v1.SuspenddRequest]{
		Params:  &v1.SuspenddRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiSuspendd,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestMoneyFlowHsgt(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.MoneyFlowHsgtRequest, v1.MoneyFlowHsgtResponse](ts, &Request[v1.MoneyFlowHsgtRequest]{
		Params:  &v1.MoneyFlowHsgtRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiMoneyFlowHsgt,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestHsgtTop10(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.HsgtTop10Request, v1.HsgtTop10Response](ts, &Request[v1.HsgtTop10Request]{
		Params:  &v1.HsgtTop10Request{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiHsgtTop10,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestGgtTop10(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.GgtTop10Request, v1.GgtTop10Response](ts, &Request[v1.GgtTop10Request]{
		Params:  &v1.GgtTop10Request{Limit: "2", Offset: "0", TsCode: "00268.HK"},
		ApiName: v1.ApiGgtTop10,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestGgtDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.GgtDailyRequest, v1.GgtDailyResponse](ts, &Request[v1.GgtDailyRequest]{
		Params:  &v1.GgtDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiGgtDaily,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestGgtMonthly(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.GgtMonthlyRequest, v1.GgtMonthlyResponse](ts, &Request[v1.GgtMonthlyRequest]{
		Params:  &v1.GgtMonthlyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiGgtMonthly,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Log("该接口没有权限, 跳过测试")
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestBakDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.BakDailyRequest, v1.BakDailyResponse](ts, &Request[v1.BakDailyRequest]{
		Params:  &v1.BakDailyRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiBakDaily,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestIncome(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IncomeRequest, v1.IncomeResponse](ts, &Request[v1.IncomeRequest]{
		Params:  &v1.IncomeRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiIncome,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestBalanceSheet(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.BalanceSheetRequest, v1.BalanceSheetResponse](ts, &Request[v1.BalanceSheetRequest]{
		Params:  &v1.BalanceSheetRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiBalanceSheet,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestCashFlow(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CashFlowRequest, v1.CashFlowResponse](ts, &Request[v1.CashFlowRequest]{
		Params:  &v1.CashFlowRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiCashFlow,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestForeCast(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ForeCastRequest, v1.ForeCastResponse](ts, &Request[v1.ForeCastRequest]{
		Params:  &v1.ForeCastRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiForeCast,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestExpress(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ExpressRequest, v1.ExpressResponse](ts, &Request[v1.ExpressRequest]{
		Params:  &v1.ExpressRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiExpress,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestDividend(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.DividendRequest, v1.DividendResponse](ts, &Request[v1.DividendRequest]{
		Params:  &v1.DividendRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiDividend,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFinaIndicator(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FinaIndicatorRequest, v1.FinaIndicatorResponse](ts, &Request[v1.FinaIndicatorRequest]{
		Params:  &v1.FinaIndicatorRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiFinaIndicator,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFinaAudit(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FinaAuditRequest, v1.FinaAuditResponse](ts, &Request[v1.FinaAuditRequest]{
		Params:  &v1.FinaAuditRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiFinaAudit,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFinaMainbz(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FinaMainbzRequest, v1.FinaMainbzResponse](ts, &Request[v1.FinaMainbzRequest]{
		Params:  &v1.FinaMainbzRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiFinaMainbz,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestDisclosureDate(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.DisclosureDateRequest, v1.DisclosureDateResponse](ts, &Request[v1.DisclosureDateRequest]{
		Params:  &v1.DisclosureDateRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiDisclosureDate,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestMargin(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.MarginRequest, v1.MarginResponse](ts, &Request[v1.MarginRequest]{
		Params:  &v1.MarginRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiMargin,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestMarginDetail(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.MarginDetailRequest, v1.MarginDetailResponse](ts, &Request[v1.MarginDetailRequest]{
		Params:  &v1.MarginDetailRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiMarginDetail,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestMarginTarget(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.MarginTargetRequest, v1.MarginTargetResponse](ts, &Request[v1.MarginTargetRequest]{
		Params:  &v1.MarginTargetRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiMarginTarget,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Log("该接口没有权限, 跳过测试")
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestTop10Holders(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.Top10HoldersRequest, v1.Top10HoldersResponse](ts, &Request[v1.Top10HoldersRequest]{
		Params:  &v1.Top10HoldersRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiTop10Holders,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestTop10Floatholders(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.Top10FloatholdersRequest, v1.Top10FloatholdersResponse](ts, &Request[v1.Top10FloatholdersRequest]{
		Params:  &v1.Top10FloatholdersRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiTop10Floatholders,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestTopList(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.TopListRequest, v1.TopListResponse](ts, &Request[v1.TopListRequest]{
		Params:  &v1.TopListRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiTopList,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestTopInst(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.TopInstRequest, v1.TopInstResponse](ts, &Request[v1.TopInstRequest]{
		Params:  &v1.TopInstRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiTopInst,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestPledgeStat(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.PledgeStatRequest, v1.PledgeStatResponse](ts, &Request[v1.PledgeStatRequest]{
		Params:  &v1.PledgeStatRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiPledgeStat,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestPledgeDetail(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.PledgeDetailRequest, v1.PledgeDetailResponse](ts, &Request[v1.PledgeDetailRequest]{
		Params:  &v1.PledgeDetailRequest{Limit: "2", Offset: "0", TsCode: "000014.SZ"},
		ApiName: v1.ApiPledgeDetail,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestRepurchase(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.RepurchaseRequest, v1.RepurchaseResponse](ts, &Request[v1.RepurchaseRequest]{
		Params:  &v1.RepurchaseRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiRepurchase,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestConcept(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ConceptRequest, v1.ConceptResponse](ts, &Request[v1.ConceptRequest]{
		Params:  &v1.ConceptRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiConcept,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestConceptDetail(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ConceptDetailRequest, v1.ConceptDetailResponse](ts, &Request[v1.ConceptDetailRequest]{
		Params:  &v1.ConceptDetailRequest{Limit: "2", Offset: "0", Id: "TS2"},
		ApiName: v1.ApiConceptDetail,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestShareFloat(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ShareFloatRequest, v1.ShareFloatResponse](ts, &Request[v1.ShareFloatRequest]{
		Params:  &v1.ShareFloatRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiShareFloat,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestBlockTrade(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.BlockTradeRequest, v1.BlockTradeResponse](ts, &Request[v1.BlockTradeRequest]{
		Params:  &v1.BlockTradeRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiBlockTrade,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestStkHolderNumber(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.StkHolderNumberRequest, v1.StkHolderNumberResponse](ts, &Request[v1.StkHolderNumberRequest]{
		Params:  &v1.StkHolderNumberRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiStkHolderNumber,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestStkHolderTrade(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.StkHolderTradeRequest, v1.StkHolderTradeResponse](ts, &Request[v1.StkHolderTradeRequest]{
		Params:  &v1.StkHolderTradeRequest{Limit: "2", Offset: "0", TsCode: "000001.SZ"},
		ApiName: v1.ApiStkHolderTrade,
		Fields:  []string{},
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

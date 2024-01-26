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
		ApiName: ApiStockBasic,
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

func TestTradeCal(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.TradeCalRequest, v1.TradeCalResponse](ts, &Request[v1.TradeCalRequest]{
		Params:  &v1.TradeCalRequest{CalDate: time.Now().Format("20060102"), Limit: "2", Offset: "0"},
		ApiName: ApiTradeCal,
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
		ApiName: ApiNameChange,
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
		ApiName: ApiHsConst,
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
		ApiName: ApiStockCompany,
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
		ApiName: ApiStkManagers,
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
		ApiName: ApiStkRewards,
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
		ApiName: ApiNewShare,
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
		ApiName: ApiBakBasic,
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
		Params:  &v1.DailyRequest{Limit: "2", Offset: "0", TradeDate: time.Now().AddDate(-50, 0, -1).Format("20060102")},
		ApiName: ApiDaily,
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
		ApiName: ApiWeekly,
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
		ApiName: ApiMonthly,
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
		ApiName: ApiAdjFactor,
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
		ApiName: ApiDailyBasic,
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
		ApiName: ApiMoneyFlow,
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
		ApiName: ApiStkLimit,
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
		ApiName: ApiSuspendd,
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
		Params:  &v1.MoneyFlowHsgtRequest{Limit: "2", Offset: "0", TradeDate: time.Now().AddDate(0, 0, -1).Format("20060102")},
		ApiName: ApiMoneyFlowHsgt,
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
		ApiName: ApiHsgtTop10,
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
		ApiName: ApiGgtTop10,
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
		ApiName: ApiGgtDaily,
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
		ApiName: ApiGgtMonthly,
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
		ApiName: ApiBakDaily,
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

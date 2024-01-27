package tushare

import (
	"os"
	"strings"
	"testing"
	"time"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestFundBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundBasicRequest, v1.FundBasicResponse](ts, &Request[v1.FundBasicRequest]{
		Params:  &v1.FundBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFundBasic,
		Fields:  v1.FieldsFundBasic,
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundCompany(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundCompanyRequest, v1.FundCompanyResponse](ts, &Request[v1.FundCompanyRequest]{
		Params:  &v1.FundCompanyRequest{},
		ApiName: v1.ApiFundCompany,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundManager(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundManagerRequest, v1.FundManagerResponse](ts, &Request[v1.FundManagerRequest]{
		Params:  &v1.FundManagerRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFundManager,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundShare(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundShareRequest, v1.FundShareResponse](ts, &Request[v1.FundShareRequest]{
		Params:  &v1.FundShareRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFundShare,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundNav(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundNavRequest, v1.FundNavResponse](ts, &Request[v1.FundNavRequest]{
		Params:  &v1.FundNavRequest{Limit: "2", Offset: "0", NavDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiFundNav,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundDiv(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundDivRequest, v1.FundDivResponse](ts, &Request[v1.FundDivRequest]{
		Params:  &v1.FundDivRequest{Limit: "2", Offset: "0", AnnDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiFundDiv,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundPortfolio(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundPortfolioRequest, v1.FundPortfolioResponse](ts, &Request[v1.FundPortfolioRequest]{
		Params:  &v1.FundPortfolioRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFundPortfolio,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundDailyRequest, v1.FundDailyResponse](ts, &Request[v1.FundDailyRequest]{
		Params:  &v1.FundDailyRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiFundDaily,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestFundAdj(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FundAdjRequest, v1.FundAdjResponse](ts, &Request[v1.FundAdjRequest]{
		Params:  &v1.FundAdjRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiFundAdj,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

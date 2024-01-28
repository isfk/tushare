package tushare

import (
	"os"
	"strings"
	"testing"
	"time"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestCbBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CbBasicRequest, v1.CbBasicResponse](ts, &Request[v1.CbBasicRequest]{
		Params:  &v1.CbBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiCbBasic,
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

func TestCbIssue(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CbIssueRequest, v1.CbIssueResponse](ts, &Request[v1.CbIssueRequest]{
		Params:  &v1.CbIssueRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiCbIssue,
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

func TestCbCall(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CbCallRequest, v1.CbCallResponse](ts, &Request[v1.CbCallRequest]{
		Params:  &v1.CbCallRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiCbCall,
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

func TestCbRate(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CbRateRequest, v1.CbRateResponse](ts, &Request[v1.CbRateRequest]{
		Params:  &v1.CbRateRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiCbRate,
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

func TestCbDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CbDailyRequest, v1.CbDailyResponse](ts, &Request[v1.CbDailyRequest]{
		Params:  &v1.CbDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiCbDaily,
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

func TestCbPriceChg(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CbPriceChgRequest, v1.CbPriceChgResponse](ts, &Request[v1.CbPriceChgRequest]{
		Params:  &v1.CbPriceChgRequest{Limit: "2", Offset: "0", TsCode: "113556.SH"},
		ApiName: v1.ApiCbPriceChg,
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

func TestCbShare(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CbShareRequest, v1.CbShareResponse](ts, &Request[v1.CbShareRequest]{
		Params:  &v1.CbShareRequest{Limit: "2", Offset: "0", TsCode: "113001.SH"},
		ApiName: v1.ApiCbShare,
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

func TestRepoDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.RepoDailyRequest, v1.RepoDailyResponse](ts, &Request[v1.RepoDailyRequest]{
		Params:  &v1.RepoDailyRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiRepoDaily,
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

func TestBondBlk(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.BondBlkRequest, v1.BondBlkResponse](ts, &Request[v1.BondBlkRequest]{
		Params:  &v1.BondBlkRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiBondBlk,
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

func TestBondBlkDetail(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.BondBlkDetailRequest, v1.BondBlkDetailResponse](ts, &Request[v1.BondBlkDetailRequest]{
		Params:  &v1.BondBlkDetailRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiBondBlkDetail,
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

func TestYcCb(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.YcCbRequest, v1.YcCbResponse](ts, &Request[v1.YcCbRequest]{
		Params:  &v1.YcCbRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1))},
		ApiName: v1.ApiYcCb,
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

func TestEcoCal(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.EcoCalRequest, v1.EcoCalResponse](ts, &Request[v1.EcoCalRequest]{
		Params:  &v1.EcoCalRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiEcoCal,
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

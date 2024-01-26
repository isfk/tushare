# tushare

> Tushare数据 Go SDK
>
> 使用泛型函数进行请求
>
> proto 定义了接口的 `数据结构` `Request` `Response` `Api***` `Fields***`, 方便使用

接口说明请参考[官网文档](https://tushare.pro/document/2)

## 使用

```sh
go get -u github.com/isfk/tushare
```

```go
package main

import (
	"log"
	"time"

	"github.com/isfk/tushare"
	v1 "github.com/isfk/tushare/gen/api/v1"
)

func main() {
	ts := tushare.NewTushare("your token")  // Get token from https://tushare.pro/user/token
	resp, err := tushare.RequestTushare[v1.TradeCalRequest, v1.TradeCalResponse](ts, &tushare.Request[v1.TradeCalRequest]{
		Params:  &v1.TradeCalRequest{CalDate: time.Now().Format("20060102"), Limit: "2", Offset: "0"},
		ApiName: v1.ApiTradeCal,
		Fields:  []string{},
	})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	if resp.Resp != nil {
		log.Printf("%v", resp.Resp)
	}
}

// go run main.go
// 2024/01/26 10:44:19 {"list":[{"exchange":"SSE","cal_date":"20240126","is_open":1,"pretrade_date":"20240125"}]}
```

## 开发

- 安装插件 `go install github.com/isfk/tushare/protoc-gen-go-tushare@latest`
- 使用快捷命令 `message_new` 添加定义
- 使用 `buf generate` 生成定义
- 在 `tushare_test.go` 中添加测试方法
- 修改 `README.md` 进度

## 进度

- [ ] 沪深股票
    - [x] 基础数据
        - [x] [股票列表](./tushare_test.go#L13)
        - [x] [交易日历](./tushare_test.go#L30)
        - [x] [股票曾用名](./tushare_test.go#L47)
        - [x] [沪深股通成分股](./tushare_test.go#L64)
        - [x] [上市公司基本信息](./tushare_test.go#L81)
        - [x] [上市公司管理层](./tushare_test.go#L98)
        - [x] [管理层薪酬和持股](./tushare_test.go#L115)
        - [x] [IPO新股上市](./tushare_test.go#L132)
        - [x] [备用列表](./tushare_test.go#L149)
    - [x] 行情数据
        - [x] [日线行情](./tushare_test.go#L166)
        - [x] [周线行情](./tushare_test.go#L183)
        - [x] [月线行情](./tushare_test.go#L200)
        - [x] ~~复权行情~~ <i>未提供HTTP接口</i>
        - [x] [复权因子](./tushare_test.go#L217)
        - [x] ~~实时快照（爬虫）~~  <i>未提供HTTP接口</i>
        - [x] ~~实时成交（爬虫）~~  <i>未提供HTTP接口</i>
        - [x] ~~实时排名（爬虫）~~  <i>未提供HTTP接口</i>
        - [x] [每日指标](./tushare_test.go#L234)
        - [x] ~~通用行情接口~~ <i>未提供HTTP接口</i>
        - [x] [个股资金流向](./tushare_test.go#L251)
        - [x] [每日涨跌停价格](./tushare_test.go#L268)
        - [x] [每日停复牌信息](./tushare_test.go#L285)
        - [x] [沪深港通资金流向](./tushare_test.go#L302)
        - [x] [沪深股通十大成交股](./tushare_test.go#L319)
        - [x] [港股通十大成交股](./tushare_test.go#L336)
        - [x] [港股通每日成交统计](./tushare_test.go#L353)
        - [x] [港股通每月成交统计](./tushare_test.go#L370) <i>因没有权限, 测试结果未验证</i>
        - [x] [备用行情](./tushare_test.go#L391)
    - [x] 财务数据
        - [x] [利润表](./tushare_test.go#L408)
        - [x] [资产负债表](./tushare_test.go#L425)
        - [x] [现金流量表](./tushare_test.go#L442)
        - [x] [业绩预告](./tushare_test.go#L459)
        - [x] [业绩快报](./tushare_test.go#L476)
        - [x] [分红送股数据](./tushare_test.go#L493)
        - [x] [财务指标数据](./tushare_test.go#L510)
        - [x] [财务审计意见](./tushare_test.go#L527)
        - [x] [主营业务构成](./tushare_test.go#L544)
        - [x] [财报披露日期表](./tushare_test.go#L561)
    - [ ] 参考数据
        - [ ] 融资融券交易汇总
        - [ ] 融资融券交易明细
        - [ ] 融资融券标的
        - [ ] 前十大股东
        - [ ] 前十大流通股东
        - [ ] 龙虎榜每日明细
        - [ ] 龙虎榜机构交易明细
        - [ ] 股权质押统计数据
        - [ ] 股权质押明细数据
        - [ ] 股票回购
        - [ ] 概念股分类表
        - [ ] 概念股明细列表
        - [ ] 限售股解禁
        - [ ] 大宗交易
        - [ ] 股票开户数据 （停）
        - [ ] 股票开户数据（旧）
        - [ ] 股东人数
        - [ ] 股东增减持
    - [ ] 特色数据
        - [ ] 券商盈利预测数据
        - [ ] 每日筹码及胜率
        - [ ] 每日筹码分布
        - [ ] 股票技术面因子
        - [ ] 中央结算系统持股统计
        - [ ] 中央结算系统持股明细
        - [ ] 沪深股通持股明细
        - [ ] 涨跌停和炸板数据
        - [ ] 机构调研数据
        - [ ] 券商月度金股
        - [ ] 游资名录
        - [ ] 游资每日明细
- [ ] 指数
    - [ ] 指数基本信息
    - [ ] 指数日线行情
    - [ ] 指数周线行情
    - [ ] 指数月线行情
    - [ ] 指数成分和权重
    - [ ] 大盘指数每日指标
    - [ ] 申万行业分类
    - [ ] 申万行业成分
    - [ ] 沪深市场每日交易统计
    - [ ] 深圳市场每日交易情况
    - [ ] 同花顺概念和行业列表
    - [ ] 同花顺概念和行业指数行情
    - [ ] 同花顺概念和行业指数成分
    - [ ] 中信行业指数日行情
    - [ ] 国际主要指数
- [ ] 公募基金
    - [ ] 基金列表
    - [ ] 基金管理人
    - [ ] 基金经理
    - [ ] 基金规模
    - [ ] 基金净值
    - [ ] 基金分红
    - [ ] 基金持仓
    - [ ] 基金行情
    - [ ] 复权因子
- [ ] 期货
    - [ ] 合约信息
    - [ ] 交易日历
    - [ ] 日线行情
    - [ ] 分钟行情
    - [ ] Tick行情
    - [ ] 仓单日报
    - [ ] 每日结算参数
    - [ ] 每日持仓排名
    - [ ] 南华期货指数行情
    - [ ] 期货主力与连续合约
    - [ ] 期货主要品种交易周报
- [ ] 现货
    - [ ] 上海黄金基础信息
    - [ ] 上海黄金现货日行情
- [ ] 期权
    - [ ] 期权合约信息
    - [ ] 期权日线行情
- [ ] 债券
    - [ ] 可转债基础信息
    - [ ] 可转债发行
    - [ ] 可转债赎回信息
    - [ ] 可转债票面利率
    - [ ] 可转债行情
    - [ ] 可转债转股价变动
    - [ ] 可转债转股结果
    - [ ] 债券回购日行情
    - [ ] 大宗交易
    - [ ] 大宗交易明细
    - [ ] 国债收益率曲线
    - [ ] 全球财经事件
- [ ] 外汇
    - [ ] 外汇基础信息（海外）
    - [ ] 外汇日线行情
- [ ] 港股
    - [ ] 港股基础信息
    - [ ] 港股交易日历
    - [ ] 港股日线行情
    - [ ] 港股分钟行情
- [ ] 美股
    - [ ] 美股基础信息
    - [ ] 美股交易日历
    - [ ] 美股日线行情
- [ ] 行业经济
- [ ] 宏观经济
- [ ] 另类数据
- [ ] 财富管理
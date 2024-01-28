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

## 支持下

<img src="https://github.com/isfk/tushare/assets/70327450/49e3673a-dddb-4e85-8b15-a332507e1b27" width="300px" />
<img src="https://github.com/isfk/tushare/assets/70327450/3d1cfcf7-5c7f-415c-a2b0-4a6f1c40c1bc" width="300px" />

## 开发

### 新方法

`tool` 目录放置了 `xlsx` 文件，从官网复制的数据，项目目录下执行 `make xlsx` 即可生成 `proto` 文件，然后 `make buf` 生成定义文件，编写测试方法即可


### 笨方法

- 安装插件 `go install github.com/isfk/tushare/protoc-gen-go-tushare@latest`
- 使用快捷命令 `message_new` 添加定义
- 使用 `buf generate` 生成定义
- 在 `stock_test.go` 中添加测试方法
- 修改 `README.md` 进度

## 进度

- [x] 沪深股票
    - [x] 基础数据
        - [x] [股票列表](./stock_test.go#L13)
        - [x] [交易日历](./stock_test.go#L34)
        - [x] [股票曾用名](./stock_test.go#L55)
        - [x] [沪深股通成分股](./stock_test.go#L76)
        - [x] [上市公司基本信息](./stock_test.go#L97)
        - [x] [上市公司管理层](./stock_test.go#L118)
        - [x] [管理层薪酬和持股](./stock_test.go#L139)
        - [x] [IPO新股上市](./stock_test.go#L160)
        - [x] [备用列表](./stock_test.go#L181)
    - [x] 行情数据
        - [x] [日线行情](./stock_test.go#L202)
        - [x] [周线行情](./stock_test.go#L223)
        - [x] [月线行情](./stock_test.go#L244)
        - [x] ~~复权行情~~ <i>未提供HTTP接口</i>
        - [x] [复权因子](./stock_test.go#L265)
        - [x] ~~实时快照（爬虫）~~ <i>未提供HTTP接口</i>
        - [x] ~~实时成交（爬虫）~~ <i>未提供HTTP接口</i>
        - [x] ~~实时排名（爬虫）~~ <i>未提供HTTP接口</i>
        - [x] [每日指标](./stock_test.go#L286)
        - [x] ~~通用行情接口~~ <i>未提供HTTP接口</i>
        - [x] [个股资金流向](./stock_test.go#L307)
        - [x] [每日涨跌停价格](./stock_test.go#L328)
        - [x] [每日停复牌信息](./stock_test.go#L349)
        - [x] [沪深港通资金流向](./stock_test.go#L370)
        - [x] [沪深股通十大成交股](./stock_test.go#L391)
        - [x] [港股通十大成交股](./stock_test.go#L412)
        - [x] [港股通每日成交统计](./stock_test.go#L433)
        - [x] [港股通每月成交统计](./stock_test.go#L454) <i>因没有权限, 测试结果未验证</i>
        - [x] [备用行情](./stock_test.go#L475)
    - [x] 财务数据
        - [x] [利润表](./stock_test.go#L496)
        - [x] [资产负债表](./stock_test.go#L517)
        - [x] [现金流量表](./stock_test.go#L538)
        - [x] [业绩预告](./stock_test.go#L559)
        - [x] [业绩快报](./stock_test.go#L580)
        - [x] [分红送股数据](./stock_test.go#L601)
        - [x] [财务指标数据](./stock_test.go#L622)
        - [x] [财务审计意见](./stock_test.go#L643)
        - [x] [主营业务构成](./stock_test.go#L664)
        - [x] [财报披露日期表](./stock_test.go#L685)
    - [x] 参考数据
        - [x] [融资融券交易汇总](./stock_test.go#L706)
        - [x] [融资融券交易明细](./stock_test.go#L727)
        - [x] [融资融券标的](./stock_test.go#L748) <i>因没有权限, 测试结果未验证</i>
        - [x] [前十大股东](./stock_test.go#L769)
        - [x] [前十大流通股东](./stock_test.go#L790)
        - [x] [龙虎榜每日明细](./stock_test.go#L811)
        - [x] [龙虎榜机构交易明细](./stock_test.go#L832)
        - [x] [股权质押统计数据](./stock_test.go#L853)
        - [x] [股权质押明细数据](./stock_test.go#L874)
        - [x] [股票回购](./stock_test.go#L895)
        - [x] [概念股分类表](./stock_test.go#L916)
        - [x] [概念股明细列表](./stock_test.go#L937)
        - [x] [限售股解禁](./stock_test.go#L958)
        - [x] [大宗交易](./stock_test.go#L979)
        - [x] ~~股票开户数据 （停）~~ <i>数据已停止更新</i>
        - [x] ~~股票开户数据（旧）~~ <i>数据已停止更新</i>
        - [x] [股东人数](./stock_test.go#L1000)
        - [x] [股东增减持](./stock_test.go#L1021)
    - [x] 特色数据
        - [x] [券商盈利预测数据](./stock_test.go#L1043)
        - [x] [每日筹码及胜率](./stock_test.go#L1064)
        - [x] [每日筹码分布](./stock_test.go#L1085)
        - [x] [股票技术面因子](./stock_test.go#L1106)
        - [x] [中央结算系统持股统计](./stock_test.go#L1127)
        - [x] [中央结算系统持股明细](./stock_test.go#L1148)
        - [x] [沪深股通持股明细](./stock_test.go#L1169)
        - [x] [涨跌停和炸板数据](./stock_test.go#L1190)
        - [x] [机构调研数据](./stock_test.go#L1211)
        - [x] [券商月度金股](./stock_test.go#L1232)
        - [x] [游资名录](./stock_test.go#L1253)
        - [x] [游资每日明细](./stock_test.go#L1274)
- [x] 指数
    - [x] [指数基本信息](./index_test.go#L12)
    - [x] [指数日线行情](./index_test.go#L33)
    - [x] [指数周线行情](./index_test.go#L54)
    - [x] [指数月线行情](./index_test.go#L75)
    - [x] [指数成分和权重](./index_test.go#L96)
    - [x] [大盘指数每日指标](./index_test.go#L117)
    - [x] [申万行业分类](./index_test.go#L138)
    - [x] [申万行业成分](./index_test.go#L159)
    - [x] [沪深市场每日交易统计](./index_test.go#L180)
    - [x] [深圳市场每日交易情况](./index_test.go#L201)
    - [x] [同花顺概念和行业列表](./index_test.go#L222)
    - [x] [同花顺概念和行业指数行情](./index_test.go#L243)
    - [x] [同花顺概念和行业指数成分](./index_test.go#L264) <i>因没有权限, 测试结果未验证</i>
    - [x] [中信行业指数日行情](./index_test.go#L285) <i>因没有权限, 测试结果未验证</i>
    - [x] [国际主要指数](./index_test.go#L306)
- [x] 公募基金
    - [x] [基金列表](./fund_test.go#L13)
    - [x] [基金管理人](./fund_test.go#L34)
    - [x] [基金经理](./fund_test.go#L55)
    - [x] [基金规模](./fund_test.go#L76)
    - [x] [基金净值](./fund_test.go#L97)
    - [x] [基金分红](./fund_test.go#L118)
    - [x] [基金持仓](./fund_test.go#L139)
    - [x] [基金行情](./fund_test.go#L160)
    - [x] [复权因子](./fund_test.go#L181)
- [x] 期货
    - [x] [合约信息](./fut_test.go#L13)
    - [x] [交易日历](./stock_test.go#L34) <i>参考股票交易日历</i>
    - [x] [日线行情](./fut_test.go#L34)
    - [x] [分钟行情](./fut_test.go#L55)
    - [x] ~~[Tick行情]~~ <i>未提供HTTP接口</i>
    - [x] [仓单日报](./fut_test.go#L76)
    - [x] [每日结算参数](./fut_test.go#L97)
    - [x] [每日持仓排名](./fut_test.go#L1118)
    - [x] [南华期货指数行情](./fut_test.go#L160)
    - [x] [期货主力与连续合约](./fut_test.go#L181)
    - [x] [期货主要品种交易周报](./fut_test.go#L181)
- [x] 现货
    - [x] [上海黄金基础信息](./sge_test.go#L12) <i>因没有权限, 测试结果未验证</i>
    - [x] [上海黄金现货日行情](./sge_test.go#L33)
- [x] 期权
    - [x] [期权合约信息](./opt_test.go#L12)
    - [x] [期权日线行情](./opt_test.go#L33)
- [ ] 债券
    - [ ] [可转债基础信息]
    - [ ] [可转债发行]
    - [ ] [可转债赎回信息]
    - [ ] [可转债票面利率]
    - [ ] [可转债行情]
    - [ ] [可转债转股价变动]
    - [ ] [可转债转股结果]
    - [ ] [债券回购日行情]
    - [ ] [大宗交易]
    - [ ] [大宗交易明细]
    - [ ] [国债收益率曲线]
    - [ ] [全球财经事件]
- [ ] 外汇
    - [ ] [外汇基础信息（海外）]
    - [ ] [外汇日线行情]
- [ ] 港股
    - [ ] [港股基础信息]
    - [ ] [港股交易日历]
    - [ ] [港股日线行情]
    - [ ] [港股分钟行情]
- [ ] 美股
    - [ ] [美股基础信息]
    - [ ] [美股交易日历]
    - [ ] [美股日线行情]
- [ ] 行业经济
- [ ] 宏观经济
- [ ] 另类数据
- [ ] 财富管理

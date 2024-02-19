# iboxchain

盒子支付 Golang SDK，好哒收款，慧掌柜
详细文档请参考 [盒子支付官方文档](https://open.iboxpay.com/doc)

## 安装

```bash
go get -u github.com/iboxpay/iboxchain
```

## 使用

```go
package main

import (
    "fmt"
    
    "github.com/iboxpay/iboxchain"
)

func main() {
    // 初始化
    iboxchain.Init("your_app_id", "your_app_secret", "your_merchant_id")
    
    // 创建订单
    order := iboxchain.Order{
        OutTradeNo: "your_out_trade_no",
        TotalFee:   100,
        Body:       "your_order_body",
        NotifyUrl:  "your_notify_url",
    }
    resp, err := iboxchain.CreateOrder(order)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(resp)
}
```

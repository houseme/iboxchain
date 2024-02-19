/*
 * Copyright IBoxChain Author(https://houseme.github.io/iboxchain/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/iboxchain.
 *
 */

package domain

// OrdersReq .交易查询 请求
type OrdersReq struct {
	Brh   string `json:"brh"`   // 机构号
	TrxID string `json:"trxId"` // 订单号，唯一
}

// OrdersResp . 交易查询 响应数据
type OrdersResp struct {
	Status    string `json:"status"`    // 处理状态，0-成功，1-失败，2-处理中
	TrxAmt    string `json:"trxAmt"`    // 提现金额，单位：分，成功时有值
	RejectAmt string `json:"rejectAmt"` // 退票金额，单位：分，成功时有值
	Desc      string `json:"desc"`      // 描述，失败时必填
}

// UnifiedOrderReq .
type UnifiedOrderReq struct {
	MchtNO          string `json:"mchtNo"`                       // 盒子的商户编号
	StoreNO         string `json:"storeNo"`                      // 盒子的门店编码
	OutOrderNO      uint64 `json:"outOrderNo,string"`            // 商户订单号
	TransAmount     uint64 `json:"transAmount,string"`           // 交易金额，单位：分
	CallbackURL     string `json:"callbackUrl,omitempty"`        // 交易完成之后，商户接收交易结果通知的地址
	RedirectURL     string `json:"redirectUrl,omitempty"`        // 支付后跳转的地址
	RedirectURLDesc string `json:"redirectUrlDesc,omitempty"`    // redirectUrl 的跳转描述：比如去开票
	ConfirmCode     string `json:"confirmCode"`                  // 支付确认码，由 4 位纯数字组成，每次请求时随机生成
	BoxSn           string `json:"boxSn,omitempty"`              // 盒子设备号
	HbFqNum         string `json:"hbFqNum,omitempty"`            // 花呗分期数：只支持传 3,6,12，仅仅支持支付宝有效
	OnlineType      string `json:"onlineType,omitempty"`         // 1 线下 2 线上，不传默认是线下
	QrValidTime     string `json:"qrValidTime,omitempty"`        // 码的有效时间，单位秒
	LedgerModel     uint   `json:"ledgerModel,string,omitempty"` // 分账类型:1 不分账 2 分账
}

// UnifiedOrderResp .
type UnifiedOrderResp struct {
	CommonResp
	OrderNO string `json:"orderNo"` // 盒子订单号
	QrCode  string `json:"qrCode"`  // 聚合码链接
}

// OrderQueryReq .无卡查询接口
type OrderQueryReq struct {
	MchtNo  string `json:"mchtNo"`
	StoreNo string `json:"storeNo"`
	OrderNo string `json:"orderNo"`
}

// OrderQueryResp .
type OrderQueryResp struct {
	CommonResp
	OrderNO           string `json:"orderNo"`
	TransAmount       uint64 `json:"transAmount,string"`
	TradeTime         string `json:"tradeTime"`         // 交易时间，格式：yyyy-MM-dd HH:mm:ss
	SubAppID          string `json:"subAppId"`          // 子商户公众账号 ID
	SubOpenID         string `json:"subOpenId"`         // 用户子标识
	ChannelTradeNO    string `json:"channelTradeNO"`    // 微信/支付宝账单上的商户订单号
	ConfirmCode       string `json:"confirmCode"`       // 支付确认码
	PayChannelOrderNO string `json:"payChannelOrderNo"` // 微信/支付宝的订单号
	PaymentType       string `json:"paymentType"`       // 支付方式，wechatActive：微信主扫 wechatPassive：微信被扫 wechatH5：微信公众号 wechatApplet：微信小程序 alipayActive：支付宝主扫 alipayPassive：支付宝被扫 alipayH5：支付宝服务窗 unionPassive：银联二维码
	UserPayAmount     string `json:"userPayAmount"`     // 用户支付金额，单位：分
	PlatformAmount    string `json:"platformAmount"`    // 微信/支付宝等平台的优惠金额，单位：分
	ShopCouponAmount  string `json:"shopCouponAmount"`  // 商家优惠金额，单位：分
	OutOrderNO        string `json:"outOrderNo"`        // 商户订单号
	TradeStatus       string `json:"tradeStatus"`       // 交易状态，1：支付中 2：支付成功 3：已部分退款 4：已退款 5：关闭 6：撤销 7：支付失败
}

// RefundReq 退款申请
type RefundReq struct {
	MchtNO       string `json:"mchtNo"`       // 盒子的商户编号
	StoreNO      string `json:"storeNo"`      // 盒子的门店编号
	OutRefundNO  string `json:"outRefundNo"`  // 商户退款订单号
	OutOrderNO   string `json:"outOrderNo"`   // 商户支付订单号
	RefundAmount string `json:"refundAmount"` // 退款金额，单位：分
}

// RefundResp .退款申请响应
type RefundResp struct {
	CommonResp
	TradeTime     string `json:"tradeTime"`     // 交易时间，格式：yyyy-MM-dd HH:mm:ss
	RefundOrderNO string `json:"refundOrderNo"` // 盒子的退款订单号
	RefundAmount  string `json:"refundAmount"`  // 退款金额，单位：分
	OutOrderNO    string `json:"outOrderNo"`    // 商家支付订单号
	OrderNO       string `json:"orderNo"`       // 盒子支付订单号
	TradeStatus   string `json:"tradeStatus"`   // 退款状态，1：退款中 2：退款成功 3：退款失败
}

// RefundQueryReq .退款查询参数
type RefundQueryReq struct {
	MchtNo      string `json:"mchtNo"`      //  盒子的商户编号
	StoreNo     string `json:"storeNo"`     // 盒子的门店编号
	OutRefundNo string `json:"outRefundNo"` // 商户退款订单号
}

// AsyncNotify .
type AsyncNotify struct {
	MchtNO                 string `json:"mchtNo"`                      // 盒子的商户编号
	StoreNO                string `json:"storeNo"`                     // 盒子的门店编码
	OutOrderNO             uint64 `json:"outOrderNo,string"`           // 商户订单号
	OrderNo                string `json:"orderNo"`                     // 盒子订单号
	ConfirmCode            string `json:"confirmCode"`                 // 支付确认码，由 4 位纯数字组成，每次请求时随机生成
	TransAmount            uint64 `json:"transAmount,string"`          // 交易金额，单位：分
	TradeTime              string `json:"tradeTime"`                   // 交易时间，格式：yyyy-MM-dd HH:mm:ss
	PayChannelOrderNo      string `json:"payChannelOrderNo"`           // 微信/支付宝的订单号
	ChannelTradeNo         string `json:"channelTradeNo"`              // 微信/支付宝账单上的商户订单号
	SubOpenID              string `json:"subOpenId"`                   // 用户标识
	PaymentType            string `json:"paymentType"`                 // 支付方式，wechatActive：微信主扫 wechatPassive：微信被扫 wechatH5：微信公众号 wechatApplet：微信小程序 alipayActive：支付宝主扫 alipayPassive：支付宝被扫 alipayH5：支付宝服务窗 unionPassive：银联二维码被扫
	PlatformCouponAmount   uint64 `json:"platformCouponAmount,string"` // 微信/支付宝等平台的优惠金额，单位：分
	ShopCouponAmount       uint64 `json:"shopCouponAmount,string"`     // 商家优惠金额，单位：分
	GoodaShopCounponAmount uint64 `json:"goodaShopCounponAmount"`      // 好哒的商家优惠，单位：分
	SubAppID               string `json:"subAppId"`                    // 微信公众号
}

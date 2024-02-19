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

// WithdrawReq .
type WithdrawReq struct {
	Brh       string `json:"brh,omitempty"`       // 机构号
	StartTime string `json:"startTime,omitempty"` // 订单开始时间，格式:yyyy-MM-dd
	EndTime   string `json:"endTime,omitempty"`   // 订单结束时间，格式:yyyy-MM-dd
	TrxID     string `json:"trxId,omitempty"`     // 订单号，唯一 长度不大于 24，当传订单号时查询订单号
	Type      string `json:"type"`                // 对账类型，1-虚拟户，2-实体户
}

// WithdrawResp .
type WithdrawResp struct {
	TotalNum uint64          `json:"totalNum"` // 总条数
	Data     []*WithdrawData `json:"data"`     // 对账数据
}

// WithdrawData .
type WithdrawData struct {
	AccTranNO       string `json:"accTranNo"`       // 记账流水号
	TranCode        string `json:"tranCode"`        // 记账交易码
	RuleCode        string `json:"ruleCode"`        // 记账规则
	IPSOrderNO      string `json:"ipsOrderNo"`      // ips 订单号
	TrxID           string `json:"trxId"`           // 商户订单号
	OutAccCode      string `json:"outAccCode"`      // 虚拟账户号
	OutAccName      string `json:"outAccName"`      // 转出方账户名
	InAccCode       string `json:"inAccCode"`       // 转入方账户号
	InAccName       string `json:"inAccName"`       // 转入方账户名
	TranType        string `json:"tranType"`        // 交易类型
	TranAmt         uint64 `json:"tranAmt"`         // 交易金额，单位：分
	AppTranDate     string `json:"appTranDate"`     // 应用交易日期，格式:yyyy-MM-dd HH: mm:ss
	AccDate         string `json:"accDate"`         // 账务日期，格式:yyyy-MM-dd HH:mm: ss
	Status          string `json:"status"`          // 状态，8-处理中，9-失败，10-成功
	Des             string `json:"des"`             // 备注
	IPSWaterNO      string `json:"ipsWaterNo"`      // ips 流水号
	BankOrderNo     string `json:"bankOrderNo"`     // 银行订单号
	ProductNO       string `json:"productNo"`       // 产品编号
	EMchtNO         string `json:"eMchtNo"`         // 商户号
	EAcctNO         string `json:"eAcctNo"`         // 交易账户号
	TradeType       string `json:"tradeType"`       // 交易类型
	SendTime        string `json:"sendTime"`        // 商户订单发起时间，格式:yyyy-MM-dd HH:mm:ss
	PreTradeTime    string `json:"preTradeTime"`    // ips 预交易时间，格式:yyyy-MM-dd HH: mm:ss
	FinishTradeTime string `json:"finishTradeTime"` // ips 交易完成时间，格式:yyyy-MM-dd HH:mm:ss
	Amount          uint64 `json:"amount"`          // 订单金额
	TradeAmount     uint64 `json:"tradeAmount"`     // 交易金额
}

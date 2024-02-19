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

// MerSplitRegReq merchant split register request
type MerSplitRegReq struct {
	MchtNo  uint64 `json:"mchtNo,string" description:"慧掌柜商户号（进件接口返回的 outMchtNo 字段）"`
	StoreNo uint64 `json:"storeNo,string" description:"慧掌柜门店号（进件接口返回的 outStoreNo 字段）"`
}

// MerSplitRegResp merchant split register response
type MerSplitRegResp struct {
	CommonResp
}

// MerSplitAccountRegReq merchant split account register request
type MerSplitAccountRegReq struct {
	AcctCategory        string                   `json:"acctCategory"`
	UserType            int                      `json:"userType"`
	UserName            string                   `json:"userName"`
	CompNm              string                   `json:"compNm"`
	LegalCertType       int                      `json:"legalCertType"`
	UserCardNo          string                   `json:"userCardNo"`
	UserCardStartDate   string                   `json:"userCardStartDate"`
	UserCardEndDate     string                   `json:"userCardEndDate"`
	AcctNm              string                   `json:"acctNm"`
	BankCardNo          string                   `json:"bankCardNo"`
	AcctBankNo          string                   `json:"acctBankNo"`
	AcctBankNm          string                   `json:"acctBankNm"`
	AcctZbankNo         string                   `json:"acctZbankNo"`
	AcctZbankNm         string                   `json:"acctZbankNm"`
	AcctZbankCode       string                   `json:"acctZbankCode"`
	ReservedPhone       string                   `json:"reservedPhone"`
	LicNo               string                   `json:"licNo"`
	LicStartDate        string                   `json:"licStartDate"`
	LicEndDate          string                   `json:"licEndDate"`
	BusinessAddressCode string                   `json:"businessAddressCode"`
	BusinessAddress     string                   `json:"businessAddress"`
	RegAddressCode      string                   `json:"regAddressCode"`
	RegAddress          string                   `json:"regAddress"`
	CompTel             string                   `json:"compTel"`
	BusinessMcc         string                   `json:"businessMcc"`
	Images              []*MerchantRegisterImage `json:"images"`
}

// MerSplitAccountRegResp merchant split account register response
type MerSplitAccountRegResp struct {
	CommonResp
	Data *MerSplitAccountData `json:"data"`
}

// MerSplitAccountData merchant split account data
type MerSplitAccountData struct {
	AcctNo uint64 `json:"acctNo,string" description:"分账账户号"`
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

// MerSplitAccountQueryReq merchant split account query request
type MerSplitAccountQueryReq struct {
	AcctNo uint64 `json:"acctNo,string" description:"分账账户号"`
}

// MerSplitAccountQueryResp merchant split account query response
type MerSplitAccountQueryResp struct {
	CommonResp
	Data *MerSplitAccountQueryData `json:"data"`
}

// MerSplitAccountQueryData merchant split account query data
type MerSplitAccountQueryData struct {
	AcctNo              uint64 `json:"acctNo,string" description:"分账账户号"`
	AcctCategory        string `json:"acctCategory"`
	UserType            int    `json:"userType"`
	UserName            string `json:"userName"`
	CompName            string `json:"compName"`
	LegalCertType       int    `json:"legalCertType"`
	UserCardNo          string `json:"userCardNo"`
	UserCardStartDate   string `json:"userCardStartDate"`
	UserCardEndDate     string `json:"userCardEndDate"`
	AcctNm              string `json:"acctNm"`
	BankCardNo          string `json:"bankCardNo"`
	AcctBankNo          string `json:"acctBankNo"`
	AcctBankNm          string `json:"acctBankNm"`
	AcctZbankNo         string `json:"acctZbankNo"`
	AcctZbankNm         string `json:"acctZbankNm"`
	AcctZbankCode       string `json:"acctZbankCode"`
	ReservedPhone       string `json:"reservedPhone"`
	LicNo               string `json:"licNo"`
	LicStartDate        string `json:"licStartDate"`
	LicEndDate          string `json:"licEndDate"`
	BusinessAddressCode string `json:"businessAddressCode"`
	BusinessAddress     string `json:"businessAddress"`
	RegAddressCode      string `json:"regAddressCode"`
	RegAddress          string `json:"regAddress"`
	CompTel             string `json:"compTel"`
	BusinessMcc         string `json:"businessMcc"`
	Status              int    `json:"status"`
	StatusDesc          string `json:"statusDesc"`
}

// MerBindAccountReq merchant bind account request
type MerBindAccountReq struct {
	MchtNo  uint64 `json:"mchtNo,string" description:"慧掌柜商户号（进件接口返回的 outMchtNo 字段）"`
	StoreNo uint64 `json:"storeNo,string" description:"慧掌柜门店号（进件接口返回的 outStoreNo 字段）"`
	AcctNo  uint64 `json:"acctNo,string" description:"分账账户号"`
}

// MerSplitTradeQueryReq merchant split trade query request
type MerSplitTradeQueryReq struct {
	MchtNo  uint64 `json:"mchtNo,string" description:"慧掌柜商户号（进件接口返回的 outMchtNo 字段）"`
	OrderNo string `json:"orderNo" description:"商户订单号"`
}

// MerSplitTradeQueryResp merchant split trade query response
type MerSplitTradeQueryResp struct {
	CommonResp
	Data *MerSplitAccountQueryData `json:"data"`
}

// MerSplitTradeQueryData merchant split trade query data
type MerSplitTradeQueryData struct {
	MchtNo          uint64 `json:"mchtNo,string"`
	TradeMchtNo     string `json:"tradeMchtNo"`
	ChannelMchtNo   string `json:"channelMchtNo"`
	OrderNo         string `json:"orderNo"`
	ChannelOrderNo  string `json:"channelOrderNo"`
	TxSubCode       string `json:"txSubCode"`
	PaymentType     string `json:"paymentType"`
	Status          string `json:"status"`
	PartAmount      string `json:"partAmount"`
	PayAmount       string `json:"payAmount"`
	PayServiceFee   string `json:"payServiceFee"`
	SettleAmount    string `json:"settleAmount"`
	TotalPerkAmount string `json:"totalPerkAmount"`
	PayTime         string `json:"payTime"`
	AlgoDate        string `json:"algoDate"`
	BizDate         string `json:"bizDate"`
	DeviceSn        string `json:"deviceSn"`
	CupsNo          string `json:"cupsNo"`
	CreateTime      string `json:"createTime"`
	RealAmount      string `json:"realAmount"`
	NotPartAmount   string `json:"notPartAmount"`
}

// OrderSplitApplyReq order split qpply request
type OrderSplitApplyReq struct {
	ApplyNo string                   `json:"applyNo"`
	MchtNo  uint64                   `json:"mchtNo,string"`
	OrderNo string                   `json:"orderNo"`
	Detail  []*OrderSplitApplyDetail `json:"detail"`
}

// OrderSplitApplyDetail order split detail
type OrderSplitApplyDetail struct {
	AcctNo   uint64 `json:"acctNo,string" description:"分账账户号"`
	AcctType int    `json:"acctType" description:"账户类型（0：商户收款账户（大 B），1：分账接收方账户（小 B））"`
	Amount   int    `json:"amount" description:"金额（分）"`
}

// OrderSplitApplyResp order split response
type OrderSplitApplyResp struct {
	CommonResp
}

// OrderSplitQueryReq order split query request
type OrderSplitQueryReq struct {
	ApplyNo string `json:"applyNo"`
	MchtNo  uint64 `json:"mchtNo,string"`
}

// OrderSplitQueryResp order split query response
type OrderSplitQueryResp struct {
	CommonResp
	Data *OrderSplitQueryData `json:"data"`
}

// OrderSplitQueryData order split query data
type OrderSplitQueryData struct {
	PayTime string `json:"payTime"`
	Status  int    `json:"status"`
	Desc    string `json:"desc"`
}

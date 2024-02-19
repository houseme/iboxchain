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

// MerchantReq .
type MerchantReq struct {
	OprID               string                   `json:"oprId,omitempty"`               // 拓展员 id
	OprInvitedCode      string                   `json:"oprInvitedCode,omitempty"`      // 拓展员推荐码
	LoginCode           string                   `json:"loginCode,omitempty"`           // 拓展员登录账号
	Brh                 string                   `json:"brh,omitempty"`                 // 机构号
	MchtNO              string                   `json:"mchtNo,omitempty"`              // 商户号
	AccountNo           string                   `json:"accountNo"`                     // 银行卡号
	AccountName         string                   `json:"accountName"`                   //  开户名 (账户名称)
	AccountType         string                   `json:"accountType"`                   // 账户类型 1 对私 0 对公
	AccountProxy        string                   `json:"accountProxy,omitempty"`        //  是否是委托清算账户 0-否，1-是，默认 0 当账户类型 选择为 '委托结算到非法人账户' 时填 1
	AgentCardNO         string                   `json:"agentCardNo,omitempty"`         // 被委托清算人身份证号码
	AgentPeriodValidity string                   `json:"agentPeriodValidity,omitempty"` // 被委托人身份证有效期 当选择结算类型为委托结算时必填，格式:20150101-20200101，若结束日期为长期 填 20150101-长期
	ZbankNO             string                   `json:"zbankNo"`                       // 开户银行支行号
	ZbankRegionCode     string                   `json:"zbankRegionCode"`               // 开户银行支行区域码
	OpenBankAccName     string                   `json:"openBankAccName,omitempty"`     // 开户许可证 - 开户名 对公结算时填写
	OpenBankAccount     string                   `json:"openBankAccount,omitempty"`     // 开户许可证 - 银行卡号 对公结算时填写
	OpenBankRegionCode  string                   `json:"openBankRegionCode,omitempty"`  // 开户许可证 - 银行地区码 对公结算时填写
	OpenCollectBankCode string                   `json:"openCollectBankCode,omitempty"` // 开户许可证 - 银行代码 对公结算时填写
	OpenUnionNO         string                   `json:"openUnionNo"`                   // 开户许可证 - 开户支行 对公结算时填写
	ChannelKind         string                   `json:"channelKind"`                   // 商户类型 慧掌柜商户：有执照填 000002 无执照填 000001，其他商户类型
	OutMchtNO           string                   `json:"outMchtNo,omitempty"`           // 慧掌柜商户号
	MchtName            string                   `json:"mchtName"`                      // 商户名称
	MchtCnShortName     string                   `json:"mchtCnShortName"`               // 商户简称
	Address             string                   `json:"address"`                       // 商户门店地址
	NoCardDiscID        string                   `json:"noCardDiscId,omitempty"`        // 值 0.38，0.33 等
	TNDiscID            string                   `json:"tNDiscId,omitempty"`            // 费率模型
	DisTNDiscID         string                   `json:"disTnDiscId,omitempty"`         // 优惠类费率模型
	DepositGear         string                   `json:"depositGear,omitempty"`         // 押金档位，单位：分
	Application         string                   `json:"application,omitempty"`         // 终端用途
	AreaNO              string                   `json:"areaNo"`                        // 经营地区区域码 参考区域码表
	Gender              string                   `json:"gender,omitempty"`              //  性别
	BusinessID          string                   `json:"businessId"`                    // 经营范围 参考经营范围表
	StorePhone          string                   `json:"storePhone,omitempty"`          // 店铺联系电话
	LicNO               string                   `json:"licNo,omitempty"`               // 营业执照号
	BusinessTerm        string                   `json:"businessTerm,omitempty"`        // 营业执照有效期 格式:20200101-20300101，长期的填 20200101-99990101
	LicenceType         string                   `json:"licenceType,omitempty"`         // 营业执照类型 0-企业法人营业执照 1-个体工商户营业执照 2-党 政，机关及事业单位 3-其他组织
	TaxNO               string                   `json:"taxNo,omitempty"`               // 税务登记号
	UserCardNO          string                   `json:"userCardNo"`                    // 商户法人身份证号
	UserPhone           string                   `json:"userPhone"`                     // 法人手机号
	UserName            string                   `json:"userName"`                      //  商户法人姓名
	ResserveMobile      string                   `json:"resserveMobile"`                // 银行预留手机号 channelkind 为 000001 时必填
	PeriodValidity      string                   `json:"periodValidity,omitempty"`      // 身份证有效期 格式:20150101-20200101，若结束日期为长期 填 20150101-长期
	Terminal            string                   `json:"terminal,omitempty"`            // 需要绑定的 sn 号列表
	Images              []*MerchantRegisterImage `json:"images,omitempty"`              // 图片的 map 集合，下面是具体的图片类型
}

// MerchantRegisterImage .
type MerchantRegisterImage struct {
	Name  string `json:"name"`  // 图片名称
	Value string `json:"value"` // 图片 code
}

// MerchantResp . 创建商户响应内容
type MerchantResp struct {
	MchtNO string `json:"mchtNo"` // 商户号
}

// MobileCheckReq . 慧掌柜校验手机号是否可用 请求参数
type MobileCheckReq struct {
	Phone       string `json:"phone"`
	ChannelKind string `json:"channelKind"` // 商户类型
}

// MobileCheckResp .慧掌柜校验手机号 返回数据
type MobileCheckResp struct {
	CanRegister string `json:"canRegister"` // 是否可注册，0-不可注册，1-可注册
	MchtNO      string `json:"mchtNo"`      // 已注册且机构号为当前机构或下级机构时返回注册的商户号
}

// MerchantQueryReq .
type MerchantQueryReq struct {
	InstID string `json:"instId"`
	MchtNo string `json:"mchtNo"`
}

// UploadFile .
type UploadFile struct {
	*UploadFileReq
}

// MerchantQueryResp 。
type MerchantQueryResp struct {
	Data       *MerchantQueryData `json:"data"`
	ErrorDesc  string             `json:"errorDesc"`
	ResultCode string             `json:"resultCode"`
	Timestamp  string             `json:"timestamp"`
}

// MerchantQueryData merchant_query 查询商户信息接口
type MerchantQueryData struct {
	Code          string `json:"code"`
	Desc          string `json:"desc"`
	MchtNo        string `json:"mchtNo"`
	StoreNo       string `json:"storeNo"`
	Option        string `json:"option"`
	Time          string `json:"time"`
	OutMchtNo     string `json:"outMchtNo"`
	ChannelMchtNo string `json:"channelMchtNo"`
}

// MerchantSubCreateReq .
type MerchantSubCreateReq struct {
	MchtNo       string        `json:"mchtNo"`
	StoreNo      string        `json:"storeNo"`
	PlatUserName string        `json:"platUserName"` // 平台用户名称，请求方保证唯一性
	UserType     string        `json:"userType" dc:"用户类型，1：企业，2：个人."`
	UserInfo     *UserInfo     `json:"userInfo"`
	BusinessInfo *BusinessInfo `json:"businessInfo"`
}

// UserInfo .
type UserInfo struct {
	UserName           string `json:"userName"`
	CertID             string `json:"certId"`
	UserMobile         string `json:"userMobile"`
	CardNo             string `json:"cardNo"`
	BankCode           string `json:"bankCode"`
	BankName           string `json:"bankName"`
	BankBranchName     string `json:"bankBranchName"`
	BankBranchCode     string `json:"bankBranchCode"`
	ValiDate           string `json:"valiDate"`
	CustAddress        string `json:"custAddress"`
	Professional       string `json:"professional"`
	CountyAreaCode     string `json:"countyAreaCode"`
	BankBranchAreaCode string `json:"bankBranchAreaCode"`
}

// MerchantSubCreateResp .
type MerchantSubCreateResp struct {
	ResultCode  string `json:"resultCode"`
	UserAccount string `json:"userAccount"`
	UserID      string `json:"userId"`
	IpsUserName string `json:"ipsUserName"`
	CreateDate  string `json:"createDate"`
	IpsMchtNo   string `json:"ipsMchtNo"`
	ErrorDesc   string `json:"errorDesc"`
	ErrorCode   string `json:"errorCode"`
	Timestamp   string `json:"timestamp"`
}

// BusinessInfo .
type BusinessInfo struct {
	CorpName            string `json:"corpName" dc:"企业名称"`                                                                                           // 企业名称
	BusinessCode        string `json:"businessCode,omitempty"`                                                                                       // 营业执照注册号
	InstitutionCode     string `json:"institutionCode,omitempty"`                                                                                    // 组织机构代码
	TaxCode             string `json:"taxCode,omitempty"`                                                                                            // 企业的税务登记号
	SocialCreditCode    string `json:"socialCreditCode" dc:"统一社会信用代码"`                                                                               // 统一社会信用代码
	LicenseStartDate    string `json:"licenseStartDate" dc:"企业营业执照起始日期"`                                                                             // 企业营业执照起始日期
	LicenseEndDate      string `json:"licenseEndDate" dc:"企业营业执照结束日期，证件有效期为长期填写:99991231"`                                                           //  企业营业执照结束日期，证件有效期为长期填写:99991231
	CorpBusinessAddress string `json:"corpBusinessAddress" dc:"企业经营地址"`                                                                              // 企业经营地址
	CorpRegAddress      string `json:"corpRegAddress" dc:"企业的注册地址"`                                                                                  //  企业的注册地址
	CorpFixedTelephone  string `json:"corpFixedTelephone" dc:"固定电话"`                                                                                 // 固定电话
	BusinessScope       string `json:"businessScope" dc:"经营范围"`                                                                                      // 经营范围
	LegalName           string `json:"legalName" dc:"法人姓名"`                                                                                          // 法人姓名
	LegalCertType       string `json:"legalCertType" dc:"法人证明类型 01 身份证 02 护照 03 港澳台居民通行证（废弃不用）04 外国人永久居留证 05 香港来往内地通行证 06 澳门来往内地通行证 07 台湾同胞来往内地通行证"` // 法人证明类型 01 身份证 02 护照 03 港澳台居民通行证（废弃不用）04 外国人永久居留证 05 香港来往内地通行证 06 澳门来往内地通行证 07 台湾同胞来往内地通行证
	LegalCertID         string `json:"legalCertId" dc:"法人证件号码"`                                                                                      // 法人证件号码
	LegalCertStartDate  string `json:"legalCertStartDate" dc:"法人证件开始日期"`                                                                             // 法人证件开始日期
	LegalCertEndDate    string `json:"legalCertEndDate" dc:"法人证件结束日期"`                                                                               // 法人证件结束日期
	LegalMobile         string `json:"legalMobile" dc:"法人手机号"`                                                                                       // 法人手机号
	ContactName         string `json:"contactName" dc:"联系人"`                                                                                         // 联系人
	ContactMobile       string `json:"contactMobile" dc:"联系人电话"`                                                                                     // 联系人电话
	ContractEmail       string `json:"contractEmail,omitempty" dc:"联系人邮箱"`                                                                           // 联系人邮箱
	BankAcctType        string `json:"bankAcctType" dc:"开户银行账户类型，userType=3 个体工商户必填；企业用户只能对公，个人用户只能对私 C-对公; P-对私"`                                   // 开户银行账户类型，userType=3 个体工商户必填；企业用户只能对公，个人用户只能对私 C-对公; P-对私
	BankAcctName        string `json:"bankAcctName" dc:"开户银行账户名，bankAcctType（或默认值）=P 时，与法人姓名一致；bankAcctType 或默认值=C 时，与 企业名称 corpName 一致"`            // 开户银行账户名，bankAcctType（或默认值）=P 时，与法人姓名一致；bankAcctType 或默认值=C 时，与 企业名称 corpName 一致
	BankCode            string `json:"bankCode" dc:"开户银行"`                                                                                           // 开户银行
	BankName            string `json:"bankName" dc:"开户银行名称"`                                                                                         // 开户银行名称
	BankAcctNo          string `json:"bankAcctNo" dc:"费率模型"`                                                                                         // 费率模型
	BankBranchName      string `json:"bankBranchName" dc:"开户银行支行名称"`                                                                                 // 开户银行支行名称
	BankBranchCode      string `json:"bankBranchCode" dc:"开户银行支行码"`                                                                                  // 开户银行支行码
	BankProv            string `json:"bankProv" dc:"开户行省份"`                                                                                          // 开户行省份
	BankArea            string `json:"bankArea" dc:"开户银行地区"`                                                                                         // 开户银行地区
	Industry            string `json:"industry,omitempty" dc:"行业名称"`                                                                                 // 行业名称
	CountyAreaCode      string `json:"countyAreaCode" dc:"支行地区码"`                                                                                    // 支行地区码
}

// AccountPartQueryReq account_part_query 获取订单分账结果请求
type AccountPartQueryReq struct {
	MchtNo     string `json:"mchtNo"`
	UserID     string `json:"userId"`
	OutOrderNo string `json:"outOrderNo"`
	TxSubCode  string `json:"txSubCode"`
}

// AccountPartQueryResp account_part_query 获取订单分账结果响应
type AccountPartQueryResp struct {
	ResultCode  string             `json:"resultCode"`
	ErrorDesc   string             `json:"errorDesc"`
	ErrorCode   string             `json:"errorCode"`
	TradeStatus string             `json:"tradeStatus"`
	PartList    []*AccountPartItem `json:"partList"`
}

// AccountPartItem 订单分账明细单体
type AccountPartItem struct {
	UserType string `json:"userType"` // 0：ips 用户账户（小 b）1：商户账户（大 b）
	UserNo   string `json:"userNo"`   // 用户编号：小 b 用户填用户号，大 b 用户填 5.0 商户号
	PartAmt  uint64 `json:"partAmt"`  // 分账金额：单位分
}

// AccountPartApplyReq account_part_apply 订单分账申请请求
type AccountPartApplyReq struct {
	MchtNo        string             `json:"mchtNo"`
	OutOrderNo    string             `json:"outOrderNo"`              // txSubCode 为 31 时原交易的商户订单号 (交易时的商户单号) 为 35 时为外部退款单号
	TxSubCode     string             `json:"txSubCode"`               // 类型：31 交易 35 退款
	OriOutOrderNo string             `json:"oriOutOrderNo,omitempty"` // 原交易外部单号 txSubCode=35 时必传
	PartList      []*AccountPartItem `json:"partList"`
}

// AccountPartApplyResp account_part_apply 订单分账申请响应
type AccountPartApplyResp struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrorDesc   string `json:"errorDesc"`
	TradeStatus string `json:"tradeStatus"`
}

// PictureUploadReq picture_upload 图片上传请求
type PictureUploadReq struct {
	MchtNo    string       `json:"mchtNo"`
	StoreNo   string       `json:"storeNo"`
	UserID    string       `json:"userId"`
	NotifyURL string       `json:"notifyUrl"`
	MediaInfo []*MediaInfo `json:"mediaInfo"`
}

// CertInfo 证件信息
type CertInfo struct {
	IdentityFront     string `json:"identityFront,omitempty"`
	IdentityContrary  string `json:"identityContrary,omitempty"`
	PersonWithIDPhoto string `json:"personWithIdPhoto,omitempty"`
	EnterpriseFront   string `json:"enterpriseFront,omitempty"`
}

// MediaInfo 媒体信息
type MediaInfo struct {
	TrxID    uint64    `json:"trxId,string"`
	CopyType uint      `json:"copyType,string"`
	CertInfo *CertInfo `json:"certInfo"`
}

// PictureUploadResp picture_upload 图片上传响应
type PictureUploadResp struct {
	ResultCode string  `json:"resultCode"`
	ErrorCode  string  `json:"errorCode"`
	ErrorDesc  string  `json:"errorDesc"`
	Infos      []*Info `json:"infos"`
}

// Info 图片上传响应
type Info struct {
	TrxID      string `json:"trxId"`      // 图片申请 id，每张图片保证唯一标识
	TrxStatus  string `json:"trxStatus"`  // 认证状态：0 未认证 1 认证中 2 认证成功 3 认证失败
	TrxReason  string `json:"trxReason"`  // 失败的原因
	TrxDtTm    string `json:"trxDtTm"`    // 日期:yyyyMMdd
	IpsTrxID   string `json:"ipsTrxId"`   // 认证流水号
	IpsTrxDtTm string `json:"ipsTrxDtTm"` // 认证时间
}

// PictureUploadAsyncResp 图片上传响应异步通知参数
type PictureUploadAsyncResp struct {
	TrxID        string `json:"trxId"`        // 图片申请 id，每张图片保证唯一标识
	TrxStatus    string `json:"trxStatus"`    // 认证状态：0 未认证 1 认证中 2 认证成功 3 认证失败
	TrxReason    string `json:"trxReason"`    // 失败的原因
	TrxDtTm      string `json:"trxDtTm"`      // 日期:yyyyMMdd
	IpsTrxID     string `json:"ipsTrxId"`     // 认证流水号
	IpsTrxDtTm   string `json:"ipsTrxDtTm"`   // 认证时间
	UserID       string `json:"userId"`       // 用户 id
	PaltUserName string `json:"paltUserName"` // 用户名
}

// AccountQueryReq account_query 开户查询信息接口
type AccountQueryReq struct {
	MchtNo       string `json:"mchtNo"`
	StoreNo      string `json:"storeNo"`
	UserID       string `json:"userId,omitempty"`
	PlatUserName string `json:"platUserName,omitempty"`
}

// AccountQueryResp account_query 开户查询信息接口
type AccountQueryResp struct {
	ResultCode   string        `json:"resultCode"`
	ErrorCode    string        `json:"errorCode"`
	ErrorDesc    string        `json:"errorDesc"`
	UserID       string        `json:"userId"`
	UserType     string        `json:"userType"`
	UserAccount  string        `json:"userAccount"`
	CreateDate   string        `json:"createDate"`
	UserInfo     *UserInfo     `json:"userInfo"`
	BusinessInfo *BusinessInfo `json:"businessInfo"`
	IpsMchtNo    string        `json:"ipsMchtNo"`
}

// QueryCertInfoReq query_certinfo  查询分账账户认证信息状态请求
type QueryCertInfoReq struct {
	MchtNo string `json:"mchtNo"`
	UserID string `json:"userId"`
}

// QueryCertInfoResp query_certinfo  查询分账账户认证信息状态响应
type QueryCertInfoResp struct {
	ResultCode string      `json:"resultCode"`
	CertList   []*CertItem `json:"certList"`
	ErrorCode  string      `json:"errorCode"`
	ErrorDesc  string      `json:"errorDesc"`
}

// CertItem 认证信息
type CertItem struct {
	StatusCode int    `json:"statusCode,string"` // 状态码，0：未认证，1：认证中，2：认证成功，3：认证失败。
	SrcType    uint   `json:"srcType,string"`    // 认证方式，1：身份证认证，2：影印件，3：人脸识别，4：手机认证，5：银行卡，6：二元认证，7：三元认证，8：法人影印件认证，9：企业法人身份认证，21：手持身份证认证。
	CertID     string `json:"certId"`            // 认证 ID.
	CopyType   int    `json:"copyType,string"`   // 影印件类型:1 个人 (身份证正反面) 2 企业 (营业执照) 3 企业法人 (法人身份证正反面) 4 个人手持身份证照片
	CertInfo   string `json:"certInfo"`          // 认证信息。json 格式的字符串，具体的认证信息
}

// OrderChargeQueryReq order_charge_query 单笔交易手续费查询请求
type OrderChargeQueryReq struct {
	MchtNo     string `json:"mchtNo"`
	TxSubCode  string `json:"txSubCode"`
	OutOrderNo string `json:"outOrderNo"`
}

// OrderChargeQueryResp order_charge_query 单笔交易手续费查询响应
type OrderChargeQueryResp struct {
	ResultCode string `json:"resultCode"` // 处理状态，0：失败 1：成功
	ErrorCode  string `json:"errorCode"`
	ErrorDesc  string `json:"errorDesc"`
	FeeAmount  uint64 `json:"feeAmount"` // 手续费金额单位分
}

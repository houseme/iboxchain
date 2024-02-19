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

package config

const (
	// unitedTrade 无卡聚合码下单接口
	unitedTrade = "api/hzg/v2/unitedtrade/aggre_order"

	// queryStatus 无卡查询接口
	queryStatus = "api/hzg/v2/unitedtrade/query_status"

	// refund . 无卡退款接口
	refund = "api/hzg/v2/unitedtrade/refund"

	// refundQuery  无卡退款查询
	refundQuery = "api/hzg/v2/unitedtrade/refund_query"

	// createMerchant 创建商户 /mcht/new_create_mcht
	createMerchant = "api/inst/mcht/new_create_mcht"

	// queryMerchant 查询商户审核状态
	queryMerchant = "api/inst/mcht/query_mcht_audit"

	// getFileToken 获取文件令牌接口
	getFileToken = "api/auth/v1/get_file_token"

	// uploadFile 上传文件接口
	uploadFile = "api/file/v1/upload_file"

	// 设备绑定接口
	bindDevice = "api/v3/hzg/manager/device/bind"

	// 设备绑定接口
	unbindDevice = "api/v3/hzg/manager/device/unbind"

	// 查询设备信息接口 .
	queryDeviceDetail = "api/v3/hzg/manager/device/detail"

	// 查询对账单下载接口
	downloadBillURL = "api/hzg/v2/bill/query_download_url"

	// 分账结果查询
	accountPartQuery = "api/hzg/v2/ledger/account_part_query"

	// 申请分账
	accountPartApply = "api/hzg/v2/ledger/account_part_apply"

	// 单笔交易手续费查询接口
	orderChargeQuery = "api/hzg/v2/ledger/order_charge_query"

	// 开户注册
	accountRegister = "api/hzg/v2/ledger/account_register"

	// 开户注册查询
	accountQuery = "api/hzg/v2/ledger/account_query"

	// 图片上传认证
	pictureUpload = "api/hzg/v2/ledger/picture_upload"

	// 用户认证信息查询
	queryCertInfo = "api/hzg/v2/ledger/query_certinfo"

	// SHA-256
	sha256 = "SHA-256"

	// md5
	md5 = "MD5"

	haoDaUnitedTrade    = "haoDaUnitedTrade"
	haoDaQueryStatus    = "haoDaQueryStatus"
	haoDaRefund         = "haoDaRefund"
	haoDaRefundQuery    = "haoDaRefundQuery"
	haoDaCreateMerchant = "haoDaCreateMerchant"
	haoDaQueryMerchant  = "haoDaQueryMerchant"
	haoDaGetFileToken   = "haoDaGetFileToken"
	haoDaUploadFile     = "haoDaUploadFile"

	haoDaBindDevice = "haoDaBindDevice"

	haoDaUnbindDevice = "haoDaUnbindDevice"

	haoDaQueryDeviceDetail = "haoDaQueryDeviceDetail"

	haoDaDownloadBillURL = "haoDaDownloadBillURL"

	haoDaAccountRegister = "haoDaAccountRegister"

	haoDaAccountQuery = "haoDaAccountQuery"

	haoDaPictureUpload = "haoDaPictureUpload"

	haoDaQueryCertInfo = "haoDaQueryCertInfo"

	haoDaOrderChargeQuery = "haoDaOrderChargeQuery"

	haoDaAccountPartApply = "haoDaAccountPartApply"

	haoDaAccountPartQuery = "haoDaAccountPartQuery"

	brh = "030085" // 机构码

	noCardDiscID = "0.3" // 支付宝，微信，银联二维码无卡费率

	channelKind = "000002" // 商户类型

	// userAgent .
	httpHeaderUserAgent = `Mozilla/5.0 (Alliance; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36`
)

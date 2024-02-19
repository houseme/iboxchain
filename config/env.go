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

import (
	"context"
)

// Env parameter
type Env struct {
}

// UnitedTrade 聚合下单
func (e *Env) UnitedTrade(_ context.Context) string {
	return haoDaUnitedTrade
}

// QueryStatus .
func (e *Env) QueryStatus(_ context.Context) string {
	return haoDaQueryStatus
}

// Refund .
func (e *Env) Refund(_ context.Context) string {
	return haoDaRefund
}

// RefundQuery .
func (e *Env) RefundQuery(_ context.Context) string {
	return haoDaRefundQuery
}

// CreateMerchant .
func (e *Env) CreateMerchant(_ context.Context) string {
	return haoDaCreateMerchant
}

// QueryMerchant .
func (e *Env) QueryMerchant(_ context.Context) string {
	return haoDaQueryMerchant
}

// GetFileToken .
func (e *Env) GetFileToken(_ context.Context) string {
	return haoDaGetFileToken
}

// UploadFile .
func (e *Env) UploadFile(_ context.Context) string {
	return haoDaUploadFile
}

// UnbindDevice .
func (e *Env) UnbindDevice(_ context.Context) string {
	return haoDaUnbindDevice
}

// BindDevice .
func (e *Env) BindDevice(_ context.Context) string {
	return haoDaBindDevice
}

// QueryDeviceDetail .
func (e *Env) QueryDeviceDetail(_ context.Context) string {
	return haoDaQueryDeviceDetail
}

// DownloadBillURL .下载账单地址
func (e *Env) DownloadBillURL(_ context.Context) string {
	return haoDaDownloadBillURL
}

// AccountRegister .
func (e *Env) AccountRegister(_ context.Context) string {
	return haoDaAccountRegister
}

// AccountQuery .
func (e *Env) AccountQuery(_ context.Context) string {
	return haoDaAccountQuery
}

// PictureUpload .
func (e *Env) PictureUpload(_ context.Context) string {
	return haoDaPictureUpload
}

// QueryCertInfo .
func (e *Env) QueryCertInfo(_ context.Context) string {
	return haoDaQueryCertInfo
}

// OrderChargeQuery .
func (e *Env) OrderChargeQuery(_ context.Context) string {
	return haoDaOrderChargeQuery
}

// AccountPartApply .申请分账
func (e *Env) AccountPartApply(_ context.Context) string {
	return haoDaAccountPartApply
}

// AccountPartQuery .查询分账
func (e *Env) AccountPartQuery(_ context.Context) string {
	return haoDaAccountPartQuery
}

// BRH .
func BRH(_ context.Context) string {
	return brh
}

// NoCardDiscID .无卡费率
func NoCardDiscID(_ context.Context) string {
	return noCardDiscID
}

// ChannelKind .商户类型
func ChannelKind(_ context.Context) string {
	return channelKind
}

// UserAgent .
func (e *Env) UserAgent(_ context.Context) string {
	return httpHeaderUserAgent
}

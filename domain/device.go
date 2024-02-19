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

// DeviceBindReq 。
type DeviceBindReq struct {
	MchtNo  string `json:"mchtNo"`
	StoreNo string `json:"storeNo"`
	Sn      string `json:"sn"`
	UserID  string `json:"userId,omitempty"`
}

// DeviceCommonResp .
type DeviceCommonResp struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

// DeviceBindResp .
type DeviceBindResp struct {
	*DeviceCommonResp
}

// DeviceUnbindReq .
type DeviceUnbindReq struct {
	MchtNo string `json:"mchtNo"`
	Sn     string `json:"sn"`
}

// DeviceUnbindResp .
type DeviceUnbindResp struct {
	*DeviceCommonResp
}

// DeviceQueryReq .
type DeviceQueryReq struct {
	Sn     string `json:"sn"`
	MchtNo string `json:"mchtNo"`
}

// DeviceQueryResp . device query response
type DeviceQueryResp struct {
	Code      int        `json:"code"`
	ErrorCode string     `json:"errorCode,omitempty"`
	ErrorDesc string     `json:"errorDesc,omitempty"`
	Data      DeviceItem `json:"data,omitempty"`
}

// DeviceItem device item
type DeviceItem struct {
	ID             int    `json:"id"`
	Sn             string `json:"sn"`
	DeviceID       int    `json:"deviceId"`
	MchtNo         string `json:"mchtNo"`
	StoreNo        string `json:"storeNo"`
	EquipmentType  string `json:"equipmentType"`
	EquipmentModel string `json:"equipmentModel"`
	Status         int    `json:"status"`
	CreateTime     string `json:"createTime"`
	CreatorID      int    `json:"creatorId"`
	UpdateTime     string `json:"updateTime"`
	StoreName      string `json:"storeName"`
	MchtName       string `json:"mchtName"`
	BindTime       string `json:"bindTime"`
	TermNo         string `json:"termNo"`
	Name           string `json:"name"`
	TermType       int    `json:"termType"`
}

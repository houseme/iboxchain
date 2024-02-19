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

// UploadFileReq upload file
type UploadFileReq struct {
	Token    string `json:"token"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FilePath string `json:"filePath"`
}

// UploadFileResp upload file response
type UploadFileResp struct {
	Code      string              `json:"code"`                // 处理状态，0：失败 1：成功
	ErrorCode string              `json:"errorCode,omitempty"` // 错误码
	ErrorDesc string              `json:"errorDesc,omitempty"` // 错误描述
	Data      *UploadFileRespData `json:"data,omitempty"`
}

// UploadFileRespData upload file response data
type UploadFileRespData struct {
	FileName string `json:"fileName"`
}

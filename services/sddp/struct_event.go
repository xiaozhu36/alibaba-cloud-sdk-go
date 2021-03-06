package sddp

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Event is a nested struct in sddp response
type Event struct {
	Status            int    `json:"Status" xml:"Status"`
	DealTime          int    `json:"DealTime" xml:"DealTime"`
	EventTime         int    `json:"EventTime" xml:"EventTime"`
	TypeName          string `json:"TypeName" xml:"TypeName"`
	DealUserId        int    `json:"DealUserId" xml:"DealUserId"`
	SubTypeCode       string `json:"SubTypeCode" xml:"SubTypeCode"`
	DepartName        string `json:"DepartName" xml:"DepartName"`
	AlertTime         int    `json:"AlertTime" xml:"AlertTime"`
	DealReason        string `json:"DealReason" xml:"DealReason"`
	TargetProductCode string `json:"TargetProductCode" xml:"TargetProductCode"`
	Id                int    `json:"Id" xml:"Id"`
	DealDisplayName   string `json:"DealDisplayName" xml:"DealDisplayName"`
	ProductCode       string `json:"ProductCode" xml:"ProductCode"`
	DisplayName       string `json:"DisplayName" xml:"DisplayName"`
	StatusName        string `json:"StatusName" xml:"StatusName"`
	Backed            bool   `json:"Backed" xml:"Backed"`
	LoginName         string `json:"LoginName" xml:"LoginName"`
	SubTypeName       string `json:"SubTypeName" xml:"SubTypeName"`
	DealLoginName     string `json:"DealLoginName" xml:"DealLoginName"`
	UserId            int    `json:"UserId" xml:"UserId"`
	TypeCode          string `json:"TypeCode" xml:"TypeCode"`
	TelephoneNum      string `json:"TelephoneNum" xml:"TelephoneNum"`
	DataInstance      string `json:"DataInstance" xml:"DataInstance"`
	Detail            Detail `json:"Detail" xml:"Detail"`
}

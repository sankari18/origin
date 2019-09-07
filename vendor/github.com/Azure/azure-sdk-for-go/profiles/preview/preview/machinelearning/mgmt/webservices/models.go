// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package webservices

import original "github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/webservices"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type AssetType = original.AssetType

const (
	AssetTypeModule   AssetType = original.AssetTypeModule
	AssetTypeResource AssetType = original.AssetTypeResource
)

type ColumnFormat = original.ColumnFormat

const (
	Byte           ColumnFormat = original.Byte
	Char           ColumnFormat = original.Char
	Complex128     ColumnFormat = original.Complex128
	Complex64      ColumnFormat = original.Complex64
	DateTime       ColumnFormat = original.DateTime
	DateTimeOffset ColumnFormat = original.DateTimeOffset
	Double         ColumnFormat = original.Double
	Duration       ColumnFormat = original.Duration
	Float          ColumnFormat = original.Float
	Int16          ColumnFormat = original.Int16
	Int32          ColumnFormat = original.Int32
	Int64          ColumnFormat = original.Int64
	Int8           ColumnFormat = original.Int8
	Uint16         ColumnFormat = original.Uint16
	Uint32         ColumnFormat = original.Uint32
	Uint64         ColumnFormat = original.Uint64
	Uint8          ColumnFormat = original.Uint8
)

type ColumnType = original.ColumnType

const (
	Boolean ColumnType = original.Boolean
	Integer ColumnType = original.Integer
	Number  ColumnType = original.Number
	String  ColumnType = original.String
)

type DiagnosticsLevel = original.DiagnosticsLevel

const (
	All   DiagnosticsLevel = original.All
	Error DiagnosticsLevel = original.Error
	None  DiagnosticsLevel = original.None
)

type InputPortType = original.InputPortType

const (
	Dataset InputPortType = original.Dataset
)

type OutputPortType = original.OutputPortType

const (
	OutputPortTypeDataset OutputPortType = original.OutputPortTypeDataset
)

type PackageType = original.PackageType

const (
	PackageTypeGraph                PackageType = original.PackageTypeGraph
	PackageTypeWebServiceProperties PackageType = original.PackageTypeWebServiceProperties
)

type ParameterType = original.ParameterType

const (
	ParameterTypeBoolean         ParameterType = original.ParameterTypeBoolean
	ParameterTypeColumnPicker    ParameterType = original.ParameterTypeColumnPicker
	ParameterTypeCredential      ParameterType = original.ParameterTypeCredential
	ParameterTypeDataGatewayName ParameterType = original.ParameterTypeDataGatewayName
	ParameterTypeDouble          ParameterType = original.ParameterTypeDouble
	ParameterTypeEnumerated      ParameterType = original.ParameterTypeEnumerated
	ParameterTypeFloat           ParameterType = original.ParameterTypeFloat
	ParameterTypeInt             ParameterType = original.ParameterTypeInt
	ParameterTypeMode            ParameterType = original.ParameterTypeMode
	ParameterTypeParameterRange  ParameterType = original.ParameterTypeParameterRange
	ParameterTypeScript          ParameterType = original.ParameterTypeScript
	ParameterTypeString          ParameterType = original.ParameterTypeString
)

type ProvisioningState = original.ProvisioningState

const (
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
	Unknown      ProvisioningState = original.Unknown
)

type AssetItem = original.AssetItem
type AssetLocation = original.AssetLocation
type ColumnSpecification = original.ColumnSpecification
type CommitmentPlan = original.CommitmentPlan
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DiagnosticsConfiguration = original.DiagnosticsConfiguration
type ExampleRequest = original.ExampleRequest
type GraphEdge = original.GraphEdge
type GraphNode = original.GraphNode
type GraphPackage = original.GraphPackage
type GraphParameter = original.GraphParameter
type GraphParameterLink = original.GraphParameterLink
type InputPort = original.InputPort
type Keys = original.Keys
type MachineLearningWorkspace = original.MachineLearningWorkspace
type ModeValueInfo = original.ModeValueInfo
type ModuleAssetParameter = original.ModuleAssetParameter
type OutputPort = original.OutputPort
type PaginatedWebServicesList = original.PaginatedWebServicesList
type PaginatedWebServicesListIterator = original.PaginatedWebServicesListIterator
type PaginatedWebServicesListPage = original.PaginatedWebServicesListPage
type PatchFuture = original.PatchFuture
type BasicProperties = original.BasicProperties
type Properties = original.Properties
type PropertiesForGraph = original.PropertiesForGraph
type RealtimeConfiguration = original.RealtimeConfiguration
type RemoveFuture = original.RemoveFuture
type Resource = original.Resource
type ServiceInputOutputSpecification = original.ServiceInputOutputSpecification
type StorageAccount = original.StorageAccount
type TableSpecification = original.TableSpecification
type WebService = original.WebService
type Client = original.Client

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAssetTypeValues() []AssetType {
	return original.PossibleAssetTypeValues()
}
func PossibleColumnFormatValues() []ColumnFormat {
	return original.PossibleColumnFormatValues()
}
func PossibleColumnTypeValues() []ColumnType {
	return original.PossibleColumnTypeValues()
}
func PossibleDiagnosticsLevelValues() []DiagnosticsLevel {
	return original.PossibleDiagnosticsLevelValues()
}
func PossibleInputPortTypeValues() []InputPortType {
	return original.PossibleInputPortTypeValues()
}
func PossibleOutputPortTypeValues() []OutputPortType {
	return original.PossibleOutputPortTypeValues()
}
func PossiblePackageTypeValues() []PackageType {
	return original.PossiblePackageTypeValues()
}
func PossibleParameterTypeValues() []ParameterType {
	return original.PossibleParameterTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}

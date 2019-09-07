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

package catalog

import original "github.com/Azure/azure-sdk-for-go/services/datalake/analytics/2016-11-01-preview/catalog"

type Client = original.Client

const (
	DefaultAdlaCatalogDNSSuffix = original.DefaultAdlaCatalogDNSSuffix
)

type BaseClient = original.BaseClient
type ACLType = original.ACLType

const (
	Group    ACLType = original.Group
	GroupObj ACLType = original.GroupObj
	Other    ACLType = original.Other
	User     ACLType = original.User
	UserObj  ACLType = original.UserObj
)

type FileType = original.FileType

const (
	Assembly FileType = original.Assembly
	Nodeploy FileType = original.Nodeploy
	Resource FileType = original.Resource
)

type PermissionType = original.PermissionType

const (
	All    PermissionType = original.All
	Alter  PermissionType = original.Alter
	Create PermissionType = original.Create
	Drop   PermissionType = original.Drop
	None   PermissionType = original.None
	Use    PermissionType = original.Use
	Write  PermissionType = original.Write
)

type ACL = original.ACL
type ACLCreateOrUpdateParameters = original.ACLCreateOrUpdateParameters
type ACLDeleteParameters = original.ACLDeleteParameters
type ACLList = original.ACLList
type ACLListIterator = original.ACLListIterator
type ACLListPage = original.ACLListPage
type DataLakeAnalyticsCatalogCredentialCreateParameters = original.DataLakeAnalyticsCatalogCredentialCreateParameters
type DataLakeAnalyticsCatalogCredentialDeleteParameters = original.DataLakeAnalyticsCatalogCredentialDeleteParameters
type DataLakeAnalyticsCatalogCredentialUpdateParameters = original.DataLakeAnalyticsCatalogCredentialUpdateParameters
type DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters = original.DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters
type DdlName = original.DdlName
type EntityID = original.EntityID
type ExternalTable = original.ExternalTable
type Item = original.Item
type ItemList = original.ItemList
type TypeFieldInfo = original.TypeFieldInfo
type USQLAssembly = original.USQLAssembly
type USQLAssemblyClr = original.USQLAssemblyClr
type USQLAssemblyDependencyInfo = original.USQLAssemblyDependencyInfo
type USQLAssemblyFileInfo = original.USQLAssemblyFileInfo
type USQLAssemblyList = original.USQLAssemblyList
type USQLAssemblyListIterator = original.USQLAssemblyListIterator
type USQLAssemblyListPage = original.USQLAssemblyListPage
type USQLCredential = original.USQLCredential
type USQLCredentialList = original.USQLCredentialList
type USQLCredentialListIterator = original.USQLCredentialListIterator
type USQLCredentialListPage = original.USQLCredentialListPage
type USQLDatabase = original.USQLDatabase
type USQLDatabaseList = original.USQLDatabaseList
type USQLDatabaseListIterator = original.USQLDatabaseListIterator
type USQLDatabaseListPage = original.USQLDatabaseListPage
type USQLDirectedColumn = original.USQLDirectedColumn
type USQLDistributionInfo = original.USQLDistributionInfo
type USQLExternalDataSource = original.USQLExternalDataSource
type USQLExternalDataSourceList = original.USQLExternalDataSourceList
type USQLExternalDataSourceListIterator = original.USQLExternalDataSourceListIterator
type USQLExternalDataSourceListPage = original.USQLExternalDataSourceListPage
type USQLIndex = original.USQLIndex
type USQLPackage = original.USQLPackage
type USQLPackageList = original.USQLPackageList
type USQLPackageListIterator = original.USQLPackageListIterator
type USQLPackageListPage = original.USQLPackageListPage
type USQLProcedure = original.USQLProcedure
type USQLProcedureList = original.USQLProcedureList
type USQLProcedureListIterator = original.USQLProcedureListIterator
type USQLProcedureListPage = original.USQLProcedureListPage
type USQLSchema = original.USQLSchema
type USQLSchemaList = original.USQLSchemaList
type USQLSchemaListIterator = original.USQLSchemaListIterator
type USQLSchemaListPage = original.USQLSchemaListPage
type USQLSecret = original.USQLSecret
type USQLTable = original.USQLTable
type USQLTableColumn = original.USQLTableColumn
type USQLTableFragment = original.USQLTableFragment
type USQLTableFragmentList = original.USQLTableFragmentList
type USQLTableFragmentListIterator = original.USQLTableFragmentListIterator
type USQLTableFragmentListPage = original.USQLTableFragmentListPage
type USQLTableList = original.USQLTableList
type USQLTableListIterator = original.USQLTableListIterator
type USQLTableListPage = original.USQLTableListPage
type USQLTablePartition = original.USQLTablePartition
type USQLTablePartitionList = original.USQLTablePartitionList
type USQLTablePartitionListIterator = original.USQLTablePartitionListIterator
type USQLTablePartitionListPage = original.USQLTablePartitionListPage
type USQLTablePreview = original.USQLTablePreview
type USQLTableStatistics = original.USQLTableStatistics
type USQLTableStatisticsList = original.USQLTableStatisticsList
type USQLTableStatisticsListIterator = original.USQLTableStatisticsListIterator
type USQLTableStatisticsListPage = original.USQLTableStatisticsListPage
type USQLTableType = original.USQLTableType
type USQLTableTypeList = original.USQLTableTypeList
type USQLTableTypeListIterator = original.USQLTableTypeListIterator
type USQLTableTypeListPage = original.USQLTableTypeListPage
type USQLTableValuedFunction = original.USQLTableValuedFunction
type USQLTableValuedFunctionList = original.USQLTableValuedFunctionList
type USQLTableValuedFunctionListIterator = original.USQLTableValuedFunctionListIterator
type USQLTableValuedFunctionListPage = original.USQLTableValuedFunctionListPage
type USQLType = original.USQLType
type USQLTypeList = original.USQLTypeList
type USQLTypeListIterator = original.USQLTypeListIterator
type USQLTypeListPage = original.USQLTypeListPage
type USQLView = original.USQLView
type USQLViewList = original.USQLViewList
type USQLViewListIterator = original.USQLViewListIterator
type USQLViewListPage = original.USQLViewListPage

func NewClient() Client {
	return original.NewClient()
}
func New() BaseClient {
	return original.New()
}
func NewWithoutDefaults(adlaCatalogDNSSuffix string) BaseClient {
	return original.NewWithoutDefaults(adlaCatalogDNSSuffix)
}
func PossibleACLTypeValues() []ACLType {
	return original.PossibleACLTypeValues()
}
func PossibleFileTypeValues() []FileType {
	return original.PossibleFileTypeValues()
}
func PossiblePermissionTypeValues() []PermissionType {
	return original.PossiblePermissionTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

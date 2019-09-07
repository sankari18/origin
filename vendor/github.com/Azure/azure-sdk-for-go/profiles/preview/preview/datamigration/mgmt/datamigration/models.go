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

package datamigration

import original "github.com/Azure/azure-sdk-for-go/services/preview/datamigration/mgmt/2018-03-31-preview/datamigration"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type AuthenticationType = original.AuthenticationType

const (
	ActiveDirectoryIntegrated AuthenticationType = original.ActiveDirectoryIntegrated
	ActiveDirectoryPassword   AuthenticationType = original.ActiveDirectoryPassword
	None                      AuthenticationType = original.None
	SQLAuthentication         AuthenticationType = original.SQLAuthentication
	WindowsAuthentication     AuthenticationType = original.WindowsAuthentication
)

type DatabaseCompatLevel = original.DatabaseCompatLevel

const (
	CompatLevel100 DatabaseCompatLevel = original.CompatLevel100
	CompatLevel110 DatabaseCompatLevel = original.CompatLevel110
	CompatLevel120 DatabaseCompatLevel = original.CompatLevel120
	CompatLevel130 DatabaseCompatLevel = original.CompatLevel130
	CompatLevel140 DatabaseCompatLevel = original.CompatLevel140
	CompatLevel80  DatabaseCompatLevel = original.CompatLevel80
	CompatLevel90  DatabaseCompatLevel = original.CompatLevel90
)

type DatabaseFileType = original.DatabaseFileType

const (
	Filestream   DatabaseFileType = original.Filestream
	Fulltext     DatabaseFileType = original.Fulltext
	Log          DatabaseFileType = original.Log
	NotSupported DatabaseFileType = original.NotSupported
	Rows         DatabaseFileType = original.Rows
)

type DatabaseMigrationStage = original.DatabaseMigrationStage

const (
	DatabaseMigrationStageBackup     DatabaseMigrationStage = original.DatabaseMigrationStageBackup
	DatabaseMigrationStageCompleted  DatabaseMigrationStage = original.DatabaseMigrationStageCompleted
	DatabaseMigrationStageFileCopy   DatabaseMigrationStage = original.DatabaseMigrationStageFileCopy
	DatabaseMigrationStageInitialize DatabaseMigrationStage = original.DatabaseMigrationStageInitialize
	DatabaseMigrationStageNone       DatabaseMigrationStage = original.DatabaseMigrationStageNone
	DatabaseMigrationStageRestore    DatabaseMigrationStage = original.DatabaseMigrationStageRestore
)

type DatabaseState = original.DatabaseState

const (
	Copying          DatabaseState = original.Copying
	Emergency        DatabaseState = original.Emergency
	Offline          DatabaseState = original.Offline
	OfflineSecondary DatabaseState = original.OfflineSecondary
	Online           DatabaseState = original.Online
	Recovering       DatabaseState = original.Recovering
	RecoveryPending  DatabaseState = original.RecoveryPending
	Restoring        DatabaseState = original.Restoring
	Suspect          DatabaseState = original.Suspect
)

type ErrorType = original.ErrorType

const (
	ErrorTypeDefault ErrorType = original.ErrorTypeDefault
	ErrorTypeError   ErrorType = original.ErrorTypeError
	ErrorTypeWarning ErrorType = original.ErrorTypeWarning
)

type LoginMigrationStage = original.LoginMigrationStage

const (
	LoginMigrationStageAssignRoleMembership       LoginMigrationStage = original.LoginMigrationStageAssignRoleMembership
	LoginMigrationStageAssignRoleOwnership        LoginMigrationStage = original.LoginMigrationStageAssignRoleOwnership
	LoginMigrationStageCompleted                  LoginMigrationStage = original.LoginMigrationStageCompleted
	LoginMigrationStageEstablishObjectPermissions LoginMigrationStage = original.LoginMigrationStageEstablishObjectPermissions
	LoginMigrationStageEstablishServerPermissions LoginMigrationStage = original.LoginMigrationStageEstablishServerPermissions
	LoginMigrationStageEstablishUserMapping       LoginMigrationStage = original.LoginMigrationStageEstablishUserMapping
	LoginMigrationStageInitialize                 LoginMigrationStage = original.LoginMigrationStageInitialize
	LoginMigrationStageLoginMigration             LoginMigrationStage = original.LoginMigrationStageLoginMigration
	LoginMigrationStageNone                       LoginMigrationStage = original.LoginMigrationStageNone
)

type LoginType = original.LoginType

const (
	AsymmetricKey LoginType = original.AsymmetricKey
	Certificate   LoginType = original.Certificate
	ExternalGroup LoginType = original.ExternalGroup
	ExternalUser  LoginType = original.ExternalUser
	SQLLogin      LoginType = original.SQLLogin
	WindowsGroup  LoginType = original.WindowsGroup
	WindowsUser   LoginType = original.WindowsUser
)

type MigrationState = original.MigrationState

const (
	MigrationStateCompleted  MigrationState = original.MigrationStateCompleted
	MigrationStateFailed     MigrationState = original.MigrationStateFailed
	MigrationStateInProgress MigrationState = original.MigrationStateInProgress
	MigrationStateNone       MigrationState = original.MigrationStateNone
	MigrationStateSkipped    MigrationState = original.MigrationStateSkipped
	MigrationStateStopped    MigrationState = original.MigrationStateStopped
	MigrationStateWarning    MigrationState = original.MigrationStateWarning
)

type MigrationStatus = original.MigrationStatus

const (
	MigrationStatusCompleted               MigrationStatus = original.MigrationStatusCompleted
	MigrationStatusCompletedWithWarnings   MigrationStatus = original.MigrationStatusCompletedWithWarnings
	MigrationStatusConfigured              MigrationStatus = original.MigrationStatusConfigured
	MigrationStatusConnecting              MigrationStatus = original.MigrationStatusConnecting
	MigrationStatusDefault                 MigrationStatus = original.MigrationStatusDefault
	MigrationStatusError                   MigrationStatus = original.MigrationStatusError
	MigrationStatusRunning                 MigrationStatus = original.MigrationStatusRunning
	MigrationStatusSelectLogins            MigrationStatus = original.MigrationStatusSelectLogins
	MigrationStatusSourceAndTargetSelected MigrationStatus = original.MigrationStatusSourceAndTargetSelected
	MigrationStatusStopped                 MigrationStatus = original.MigrationStatusStopped
)

type NameCheckFailureReason = original.NameCheckFailureReason

const (
	AlreadyExists NameCheckFailureReason = original.AlreadyExists
	Invalid       NameCheckFailureReason = original.Invalid
)

type ObjectType = original.ObjectType

const (
	Function         ObjectType = original.Function
	StoredProcedures ObjectType = original.StoredProcedures
	Table            ObjectType = original.Table
	User             ObjectType = original.User
	View             ObjectType = original.View
)

type ProjectProvisioningState = original.ProjectProvisioningState

const (
	Deleting  ProjectProvisioningState = original.Deleting
	Succeeded ProjectProvisioningState = original.Succeeded
)

type ProjectSourcePlatform = original.ProjectSourcePlatform

const (
	SQL     ProjectSourcePlatform = original.SQL
	Unknown ProjectSourcePlatform = original.Unknown
)

type ProjectTargetPlatform = original.ProjectTargetPlatform

const (
	ProjectTargetPlatformSQLDB   ProjectTargetPlatform = original.ProjectTargetPlatformSQLDB
	ProjectTargetPlatformSQLMI   ProjectTargetPlatform = original.ProjectTargetPlatformSQLMI
	ProjectTargetPlatformUnknown ProjectTargetPlatform = original.ProjectTargetPlatformUnknown
)

type ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleType

const (
	ResourceSkuCapacityScaleTypeAutomatic ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeAutomatic
	ResourceSkuCapacityScaleTypeManual    ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeManual
	ResourceSkuCapacityScaleTypeNone      ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeNone
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
)

type ResultType = original.ResultType

const (
	ResultTypeDatabaseLevelOutput             ResultType = original.ResultTypeDatabaseLevelOutput
	ResultTypeErrorOutput                     ResultType = original.ResultTypeErrorOutput
	ResultTypeMigrateSQLServerSQLDbTaskOutput ResultType = original.ResultTypeMigrateSQLServerSQLDbTaskOutput
	ResultTypeMigrationLevelOutput            ResultType = original.ResultTypeMigrationLevelOutput
	ResultTypeTableLevelOutput                ResultType = original.ResultTypeTableLevelOutput
)

type ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutput

const (
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeAgentJobLevelOutput                ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeAgentJobLevelOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeConnectToSourceSQLServerTaskOutput ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeConnectToSourceSQLServerTaskOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeDatabaseLevelOutput                ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeDatabaseLevelOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeLoginLevelOutput                   ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeLoginLevelOutput
	ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeTaskLevelOutput                    ResultTypeBasicConnectToSourceSQLServerTaskOutput = original.ResultTypeBasicConnectToSourceSQLServerTaskOutputResultTypeTaskLevelOutput
)

type ResultTypeBasicMigrateSQLServerSQLMITaskOutput = original.ResultTypeBasicMigrateSQLServerSQLMITaskOutput

const (
	ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeAgentJobLevelOutput             ResultTypeBasicMigrateSQLServerSQLMITaskOutput = original.ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeAgentJobLevelOutput
	ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeDatabaseLevelOutput             ResultTypeBasicMigrateSQLServerSQLMITaskOutput = original.ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeDatabaseLevelOutput
	ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeErrorOutput                     ResultTypeBasicMigrateSQLServerSQLMITaskOutput = original.ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeErrorOutput
	ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeLoginLevelOutput                ResultTypeBasicMigrateSQLServerSQLMITaskOutput = original.ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeLoginLevelOutput
	ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeMigrateSQLServerSQLMITaskOutput ResultTypeBasicMigrateSQLServerSQLMITaskOutput = original.ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeMigrateSQLServerSQLMITaskOutput
	ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeMigrationLevelOutput            ResultTypeBasicMigrateSQLServerSQLMITaskOutput = original.ResultTypeBasicMigrateSQLServerSQLMITaskOutputResultTypeMigrationLevelOutput
)

type ServerLevelPermissionsGroup = original.ServerLevelPermissionsGroup

const (
	Default                         ServerLevelPermissionsGroup = original.Default
	MigrationFromSQLServerToAzureDB ServerLevelPermissionsGroup = original.MigrationFromSQLServerToAzureDB
)

type ServiceProvisioningState = original.ServiceProvisioningState

const (
	ServiceProvisioningStateAccepted      ServiceProvisioningState = original.ServiceProvisioningStateAccepted
	ServiceProvisioningStateDeleting      ServiceProvisioningState = original.ServiceProvisioningStateDeleting
	ServiceProvisioningStateDeploying     ServiceProvisioningState = original.ServiceProvisioningStateDeploying
	ServiceProvisioningStateFailed        ServiceProvisioningState = original.ServiceProvisioningStateFailed
	ServiceProvisioningStateFailedToStart ServiceProvisioningState = original.ServiceProvisioningStateFailedToStart
	ServiceProvisioningStateFailedToStop  ServiceProvisioningState = original.ServiceProvisioningStateFailedToStop
	ServiceProvisioningStateStarting      ServiceProvisioningState = original.ServiceProvisioningStateStarting
	ServiceProvisioningStateStopped       ServiceProvisioningState = original.ServiceProvisioningStateStopped
	ServiceProvisioningStateStopping      ServiceProvisioningState = original.ServiceProvisioningStateStopping
	ServiceProvisioningStateSucceeded     ServiceProvisioningState = original.ServiceProvisioningStateSucceeded
)

type ServiceScalability = original.ServiceScalability

const (
	ServiceScalabilityAutomatic ServiceScalability = original.ServiceScalabilityAutomatic
	ServiceScalabilityManual    ServiceScalability = original.ServiceScalabilityManual
	ServiceScalabilityNone      ServiceScalability = original.ServiceScalabilityNone
)

type Severity = original.Severity

const (
	SeverityError   Severity = original.SeverityError
	SeverityMessage Severity = original.SeverityMessage
	SeverityWarning Severity = original.SeverityWarning
)

type TaskState = original.TaskState

const (
	TaskStateCanceled              TaskState = original.TaskStateCanceled
	TaskStateFailed                TaskState = original.TaskStateFailed
	TaskStateFailedInputValidation TaskState = original.TaskStateFailedInputValidation
	TaskStateFaulted               TaskState = original.TaskStateFaulted
	TaskStateQueued                TaskState = original.TaskStateQueued
	TaskStateRunning               TaskState = original.TaskStateRunning
	TaskStateSucceeded             TaskState = original.TaskStateSucceeded
	TaskStateUnknown               TaskState = original.TaskStateUnknown
)

type TaskType = original.TaskType

const (
	TaskTypeConnectToSourceSQLServer                    TaskType = original.TaskTypeConnectToSourceSQLServer
	TaskTypeConnectToTargetAzureSQLDbMI                 TaskType = original.TaskTypeConnectToTargetAzureSQLDbMI
	TaskTypeConnectToTargetSQLDb                        TaskType = original.TaskTypeConnectToTargetSQLDb
	TaskTypeGetUserTablesSQL                            TaskType = original.TaskTypeGetUserTablesSQL
	TaskTypeMigrateSQLServerAzureSQLDbMI                TaskType = original.TaskTypeMigrateSQLServerAzureSQLDbMI
	TaskTypeMigrateSQLServerSQLDb                       TaskType = original.TaskTypeMigrateSQLServerSQLDb
	TaskTypeUnknown                                     TaskType = original.TaskTypeUnknown
	TaskTypeValidateMigrationInputSQLServerAzureSQLDbMI TaskType = original.TaskTypeValidateMigrationInputSQLServerAzureSQLDbMI
)

type Type = original.Type

const (
	TypeSQLConnectionInfo Type = original.TypeSQLConnectionInfo
	TypeUnknown           Type = original.TypeUnknown
)

type UpdateActionType = original.UpdateActionType

const (
	AddedOnTarget   UpdateActionType = original.AddedOnTarget
	ChangedOnTarget UpdateActionType = original.ChangedOnTarget
	DeletedOnTarget UpdateActionType = original.DeletedOnTarget
)

type ValidationStatus = original.ValidationStatus

const (
	ValidationStatusCompleted           ValidationStatus = original.ValidationStatusCompleted
	ValidationStatusCompletedWithIssues ValidationStatus = original.ValidationStatusCompletedWithIssues
	ValidationStatusDefault             ValidationStatus = original.ValidationStatusDefault
	ValidationStatusFailed              ValidationStatus = original.ValidationStatusFailed
	ValidationStatusInitialized         ValidationStatus = original.ValidationStatusInitialized
	ValidationStatusInProgress          ValidationStatus = original.ValidationStatusInProgress
	ValidationStatusNotStarted          ValidationStatus = original.ValidationStatusNotStarted
	ValidationStatusStopped             ValidationStatus = original.ValidationStatusStopped
)

type APIError = original.APIError
type AvailableServiceSku = original.AvailableServiceSku
type AvailableServiceSkuCapacity = original.AvailableServiceSkuCapacity
type AvailableServiceSkuSku = original.AvailableServiceSkuSku
type BlobShare = original.BlobShare
type BasicConnectionInfo = original.BasicConnectionInfo
type ConnectionInfo = original.ConnectionInfo
type ConnectToSourceSQLServerTaskInput = original.ConnectToSourceSQLServerTaskInput
type BasicConnectToSourceSQLServerTaskOutput = original.BasicConnectToSourceSQLServerTaskOutput
type ConnectToSourceSQLServerTaskOutput = original.ConnectToSourceSQLServerTaskOutput
type ConnectToSourceSQLServerTaskOutputAgentJobLevel = original.ConnectToSourceSQLServerTaskOutputAgentJobLevel
type ConnectToSourceSQLServerTaskOutputDatabaseLevel = original.ConnectToSourceSQLServerTaskOutputDatabaseLevel
type ConnectToSourceSQLServerTaskOutputLoginLevel = original.ConnectToSourceSQLServerTaskOutputLoginLevel
type ConnectToSourceSQLServerTaskOutputTaskLevel = original.ConnectToSourceSQLServerTaskOutputTaskLevel
type ConnectToSourceSQLServerTaskProperties = original.ConnectToSourceSQLServerTaskProperties
type ConnectToTargetSQLDbTaskInput = original.ConnectToTargetSQLDbTaskInput
type ConnectToTargetSQLDbTaskOutput = original.ConnectToTargetSQLDbTaskOutput
type ConnectToTargetSQLDbTaskProperties = original.ConnectToTargetSQLDbTaskProperties
type ConnectToTargetSQLMITaskInput = original.ConnectToTargetSQLMITaskInput
type ConnectToTargetSQLMITaskOutput = original.ConnectToTargetSQLMITaskOutput
type ConnectToTargetSQLMITaskProperties = original.ConnectToTargetSQLMITaskProperties
type Database = original.Database
type DatabaseFileInfo = original.DatabaseFileInfo
type DatabaseFileInput = original.DatabaseFileInput
type DatabaseInfo = original.DatabaseInfo
type DatabaseObjectName = original.DatabaseObjectName
type DatabaseSummaryResult = original.DatabaseSummaryResult
type DatabaseTable = original.DatabaseTable
type DataIntegrityValidationResult = original.DataIntegrityValidationResult
type DataItemMigrationSummaryResult = original.DataItemMigrationSummaryResult
type Error = original.Error
type ExecutionStatistics = original.ExecutionStatistics
type FileShare = original.FileShare
type GetUserTablesSQLTaskInput = original.GetUserTablesSQLTaskInput
type GetUserTablesSQLTaskOutput = original.GetUserTablesSQLTaskOutput
type GetUserTablesSQLTaskProperties = original.GetUserTablesSQLTaskProperties
type MigrateSQLServerSQLDbDatabaseInput = original.MigrateSQLServerSQLDbDatabaseInput
type MigrateSQLServerSQLDbTaskInput = original.MigrateSQLServerSQLDbTaskInput
type BasicMigrateSQLServerSQLDbTaskOutput = original.BasicMigrateSQLServerSQLDbTaskOutput
type MigrateSQLServerSQLDbTaskOutput = original.MigrateSQLServerSQLDbTaskOutput
type MigrateSQLServerSQLDbTaskOutputDatabaseLevel = original.MigrateSQLServerSQLDbTaskOutputDatabaseLevel
type MigrateSQLServerSQLDbTaskOutputError = original.MigrateSQLServerSQLDbTaskOutputError
type MigrateSQLServerSQLDbTaskOutputMigrationLevel = original.MigrateSQLServerSQLDbTaskOutputMigrationLevel
type MigrateSQLServerSQLDbTaskOutputTableLevel = original.MigrateSQLServerSQLDbTaskOutputTableLevel
type MigrateSQLServerSQLDbTaskProperties = original.MigrateSQLServerSQLDbTaskProperties
type MigrateSQLServerSQLMIDatabaseInput = original.MigrateSQLServerSQLMIDatabaseInput
type MigrateSQLServerSQLMITaskInput = original.MigrateSQLServerSQLMITaskInput
type BasicMigrateSQLServerSQLMITaskOutput = original.BasicMigrateSQLServerSQLMITaskOutput
type MigrateSQLServerSQLMITaskOutput = original.MigrateSQLServerSQLMITaskOutput
type MigrateSQLServerSQLMITaskOutputAgentJobLevel = original.MigrateSQLServerSQLMITaskOutputAgentJobLevel
type MigrateSQLServerSQLMITaskOutputDatabaseLevel = original.MigrateSQLServerSQLMITaskOutputDatabaseLevel
type MigrateSQLServerSQLMITaskOutputError = original.MigrateSQLServerSQLMITaskOutputError
type MigrateSQLServerSQLMITaskOutputLoginLevel = original.MigrateSQLServerSQLMITaskOutputLoginLevel
type MigrateSQLServerSQLMITaskOutputMigrationLevel = original.MigrateSQLServerSQLMITaskOutputMigrationLevel
type MigrateSQLServerSQLMITaskProperties = original.MigrateSQLServerSQLMITaskProperties
type MigrateSQLServerSQLServerDatabaseInput = original.MigrateSQLServerSQLServerDatabaseInput
type MigrationEligibilityInfo = original.MigrationEligibilityInfo
type MigrationReportResult = original.MigrationReportResult
type MigrationTableMetadata = original.MigrationTableMetadata
type MigrationValidationDatabaseLevelResult = original.MigrationValidationDatabaseLevelResult
type MigrationValidationDatabaseSummaryResult = original.MigrationValidationDatabaseSummaryResult
type MigrationValidationOptions = original.MigrationValidationOptions
type MigrationValidationResult = original.MigrationValidationResult
type NameAvailabilityRequest = original.NameAvailabilityRequest
type NameAvailabilityResponse = original.NameAvailabilityResponse
type ODataError = original.ODataError
type Project = original.Project
type ProjectList = original.ProjectList
type ProjectListIterator = original.ProjectListIterator
type ProjectListPage = original.ProjectListPage
type ProjectMetadata = original.ProjectMetadata
type ProjectProperties = original.ProjectProperties
type ProjectTask = original.ProjectTask
type BasicProjectTaskProperties = original.BasicProjectTaskProperties
type ProjectTaskProperties = original.ProjectTaskProperties
type QueryAnalysisValidationResult = original.QueryAnalysisValidationResult
type QueryExecutionResult = original.QueryExecutionResult
type Quota = original.Quota
type QuotaList = original.QuotaList
type QuotaListIterator = original.QuotaListIterator
type QuotaListPage = original.QuotaListPage
type QuotaName = original.QuotaName
type ReportableException = original.ReportableException
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuCapacity = original.ResourceSkuCapacity
type ResourceSkuCosts = original.ResourceSkuCosts
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkusResult = original.ResourceSkusResult
type ResourceSkusResultIterator = original.ResourceSkusResultIterator
type ResourceSkusResultPage = original.ResourceSkusResultPage
type SchemaComparisonValidationResult = original.SchemaComparisonValidationResult
type SchemaComparisonValidationResultType = original.SchemaComparisonValidationResultType
type Service = original.Service
type ServiceList = original.ServiceList
type ServiceListIterator = original.ServiceListIterator
type ServiceListPage = original.ServiceListPage
type ServiceOperation = original.ServiceOperation
type ServiceOperationDisplay = original.ServiceOperationDisplay
type ServiceOperationList = original.ServiceOperationList
type ServiceOperationListIterator = original.ServiceOperationListIterator
type ServiceOperationListPage = original.ServiceOperationListPage
type ServiceProperties = original.ServiceProperties
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServiceSku = original.ServiceSku
type ServiceSkuList = original.ServiceSkuList
type ServiceSkuListIterator = original.ServiceSkuListIterator
type ServiceSkuListPage = original.ServiceSkuListPage
type ServicesStartFuture = original.ServicesStartFuture
type ServicesStopFuture = original.ServicesStopFuture
type ServiceStatusResponse = original.ServiceStatusResponse
type ServicesUpdateFuture = original.ServicesUpdateFuture
type SQLConnectionInfo = original.SQLConnectionInfo
type SQLMigrationTaskInput = original.SQLMigrationTaskInput
type StartMigrationScenarioServerRoleResult = original.StartMigrationScenarioServerRoleResult
type TaskList = original.TaskList
type TaskListIterator = original.TaskListIterator
type TaskListPage = original.TaskListPage
type TrackedResource = original.TrackedResource
type ValidateMigrationInputSQLServerSQLMITaskInput = original.ValidateMigrationInputSQLServerSQLMITaskInput
type ValidateMigrationInputSQLServerSQLMITaskOutput = original.ValidateMigrationInputSQLServerSQLMITaskOutput
type ValidateMigrationInputSQLServerSQLMITaskProperties = original.ValidateMigrationInputSQLServerSQLMITaskProperties
type ValidationError = original.ValidationError
type WaitStatistics = original.WaitStatistics
type OperationsClient = original.OperationsClient
type ProjectsClient = original.ProjectsClient
type ResourceSkusClient = original.ResourceSkusClient
type ServicesClient = original.ServicesClient
type TasksClient = original.TasksClient
type UsagesClient = original.UsagesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return original.PossibleAuthenticationTypeValues()
}
func PossibleDatabaseCompatLevelValues() []DatabaseCompatLevel {
	return original.PossibleDatabaseCompatLevelValues()
}
func PossibleDatabaseFileTypeValues() []DatabaseFileType {
	return original.PossibleDatabaseFileTypeValues()
}
func PossibleDatabaseMigrationStageValues() []DatabaseMigrationStage {
	return original.PossibleDatabaseMigrationStageValues()
}
func PossibleDatabaseStateValues() []DatabaseState {
	return original.PossibleDatabaseStateValues()
}
func PossibleErrorTypeValues() []ErrorType {
	return original.PossibleErrorTypeValues()
}
func PossibleLoginMigrationStageValues() []LoginMigrationStage {
	return original.PossibleLoginMigrationStageValues()
}
func PossibleLoginTypeValues() []LoginType {
	return original.PossibleLoginTypeValues()
}
func PossibleMigrationStateValues() []MigrationState {
	return original.PossibleMigrationStateValues()
}
func PossibleMigrationStatusValues() []MigrationStatus {
	return original.PossibleMigrationStatusValues()
}
func PossibleNameCheckFailureReasonValues() []NameCheckFailureReason {
	return original.PossibleNameCheckFailureReasonValues()
}
func PossibleObjectTypeValues() []ObjectType {
	return original.PossibleObjectTypeValues()
}
func PossibleProjectProvisioningStateValues() []ProjectProvisioningState {
	return original.PossibleProjectProvisioningStateValues()
}
func PossibleProjectSourcePlatformValues() []ProjectSourcePlatform {
	return original.PossibleProjectSourcePlatformValues()
}
func PossibleProjectTargetPlatformValues() []ProjectTargetPlatform {
	return original.PossibleProjectTargetPlatformValues()
}
func PossibleResourceSkuCapacityScaleTypeValues() []ResourceSkuCapacityScaleType {
	return original.PossibleResourceSkuCapacityScaleTypeValues()
}
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return original.PossibleResourceSkuRestrictionsReasonCodeValues()
}
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return original.PossibleResourceSkuRestrictionsTypeValues()
}
func PossibleResultTypeValues() []ResultType {
	return original.PossibleResultTypeValues()
}
func PossibleResultTypeBasicConnectToSourceSQLServerTaskOutputValues() []ResultTypeBasicConnectToSourceSQLServerTaskOutput {
	return original.PossibleResultTypeBasicConnectToSourceSQLServerTaskOutputValues()
}
func PossibleResultTypeBasicMigrateSQLServerSQLMITaskOutputValues() []ResultTypeBasicMigrateSQLServerSQLMITaskOutput {
	return original.PossibleResultTypeBasicMigrateSQLServerSQLMITaskOutputValues()
}
func PossibleServerLevelPermissionsGroupValues() []ServerLevelPermissionsGroup {
	return original.PossibleServerLevelPermissionsGroupValues()
}
func PossibleServiceProvisioningStateValues() []ServiceProvisioningState {
	return original.PossibleServiceProvisioningStateValues()
}
func PossibleServiceScalabilityValues() []ServiceScalability {
	return original.PossibleServiceScalabilityValues()
}
func PossibleSeverityValues() []Severity {
	return original.PossibleSeverityValues()
}
func PossibleTaskStateValues() []TaskState {
	return original.PossibleTaskStateValues()
}
func PossibleTaskTypeValues() []TaskType {
	return original.PossibleTaskTypeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUpdateActionTypeValues() []UpdateActionType {
	return original.PossibleUpdateActionTypeValues()
}
func PossibleValidationStatusValues() []ValidationStatus {
	return original.PossibleValidationStatusValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProjectsClient(subscriptionID string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkusClient(subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClient(subscriptionID)
}
func NewResourceSkusClientWithBaseURI(baseURI string, subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTasksClient(subscriptionID string) TasksClient {
	return original.NewTasksClient(subscriptionID)
}
func NewTasksClientWithBaseURI(baseURI string, subscriptionID string) TasksClient {
	return original.NewTasksClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

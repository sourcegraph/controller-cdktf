package googleoracledatabaseautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabaseautonomousdatabase/internal"
)

type GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference interface {
	cdktf.ComplexObject
	ActualUsedDataStorageSizeTb() *float64
	AllocatedStorageSizeTb() *float64
	ApexDetails() GoogleOracleDatabaseAutonomousDatabasePropertiesApexDetailsList
	ArePrimaryAllowlistedIpsUsed() cdktf.IResolvable
	AutonomousContainerDatabaseId() *string
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodDays() *float64
	SetBackupRetentionPeriodDays(val *float64)
	BackupRetentionPeriodDaysInput() *float64
	CharacterSet() *string
	SetCharacterSet(val *string)
	CharacterSetInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComputeCount() *float64
	SetComputeCount(val *float64)
	ComputeCountInput() *float64
	ConnectionStrings() GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionStringsList
	ConnectionUrls() GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerContacts() GoogleOracleDatabaseAutonomousDatabasePropertiesCustomerContactsList
	CustomerContactsInput() interface{}
	DatabaseManagementState() *string
	DataSafeState() *string
	DataStorageSizeGb() *float64
	SetDataStorageSizeGb(val *float64)
	DataStorageSizeGbInput() *float64
	DataStorageSizeTb() *float64
	SetDataStorageSizeTb(val *float64)
	DataStorageSizeTbInput() *float64
	DbEdition() *string
	SetDbEdition(val *string)
	DbEditionInput() *string
	DbVersion() *string
	SetDbVersion(val *string)
	DbVersionInput() *string
	DbWorkload() *string
	SetDbWorkload(val *string)
	DbWorkloadInput() *string
	FailedDataRecoveryDuration() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleOracleDatabaseAutonomousDatabaseProperties
	SetInternalValue(val *GoogleOracleDatabaseAutonomousDatabaseProperties)
	IsAutoScalingEnabled() interface{}
	SetIsAutoScalingEnabled(val interface{})
	IsAutoScalingEnabledInput() interface{}
	IsLocalDataGuardEnabled() cdktf.IResolvable
	IsStorageAutoScalingEnabled() interface{}
	SetIsStorageAutoScalingEnabled(val interface{})
	IsStorageAutoScalingEnabledInput() interface{}
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	LifecycleDetails() *string
	LocalAdgAutoFailoverMaxDataLossLimit() *float64
	LocalDisasterRecoveryType() *string
	LocalStandbyDb() GoogleOracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList
	MaintenanceBeginTime() *string
	MaintenanceEndTime() *string
	MaintenanceScheduleType() *string
	SetMaintenanceScheduleType(val *string)
	MaintenanceScheduleTypeInput() *string
	MemoryPerOracleComputeUnitGbs() *float64
	MemoryTableGbs() *float64
	MtlsConnectionRequired() interface{}
	SetMtlsConnectionRequired(val interface{})
	MtlsConnectionRequiredInput() interface{}
	NCharacterSet() *string
	SetNCharacterSet(val *string)
	NCharacterSetInput() *string
	NextLongTermBackupTime() *string
	Ocid() *string
	OciUrl() *string
	OpenMode() *string
	OperationsInsightsState() *string
	SetOperationsInsightsState(val *string)
	OperationsInsightsStateInput() *string
	PeerDbIds() *[]*string
	PermissionLevel() *string
	PrivateEndpoint() *string
	PrivateEndpointIp() *string
	SetPrivateEndpointIp(val *string)
	PrivateEndpointIpInput() *string
	PrivateEndpointLabel() *string
	SetPrivateEndpointLabel(val *string)
	PrivateEndpointLabelInput() *string
	RefreshableMode() *string
	RefreshableState() *string
	Role() *string
	ScheduledOperationDetails() GoogleOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList
	SqlWebDeveloperUrl() *string
	State() *string
	SupportedCloneRegions() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalAutoBackupStorageSizeGbs() *float64
	UsedDataStorageSizeTbs() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCustomerContacts(value interface{})
	ResetBackupRetentionPeriodDays()
	ResetCharacterSet()
	ResetComputeCount()
	ResetCustomerContacts()
	ResetDataStorageSizeGb()
	ResetDataStorageSizeTb()
	ResetDbEdition()
	ResetDbVersion()
	ResetIsAutoScalingEnabled()
	ResetIsStorageAutoScalingEnabled()
	ResetMaintenanceScheduleType()
	ResetMtlsConnectionRequired()
	ResetNCharacterSet()
	ResetOperationsInsightsState()
	ResetPrivateEndpointIp()
	ResetPrivateEndpointLabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ActualUsedDataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) AllocatedStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ApexDetails() GoogleOracleDatabaseAutonomousDatabasePropertiesApexDetailsList {
	var returns GoogleOracleDatabaseAutonomousDatabasePropertiesApexDetailsList
	_jsii_.Get(
		j,
		"apexDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ArePrimaryAllowlistedIpsUsed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"arePrimaryAllowlistedIpsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) AutonomousContainerDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousContainerDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) BackupRetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) BackupRetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ConnectionStrings() GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionStringsList {
	var returns GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ConnectionUrls() GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList {
	var returns GoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList
	_jsii_.Get(
		j,
		"connectionUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CustomerContacts() GoogleOracleDatabaseAutonomousDatabasePropertiesCustomerContactsList {
	var returns GoogleOracleDatabaseAutonomousDatabasePropertiesCustomerContactsList
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CustomerContactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DatabaseManagementState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseManagementState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataSafeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSafeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeTbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbWorkloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) FailedDataRecoveryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failedDataRecoveryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) InternalValue() *GoogleOracleDatabaseAutonomousDatabaseProperties {
	var returns *GoogleOracleDatabaseAutonomousDatabaseProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsAutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsAutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsLocalDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLocalDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsStorageAutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStorageAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsStorageAutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStorageAutoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalAdgAutoFailoverMaxDataLossLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalDisasterRecoveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDisasterRecoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalStandbyDb() GoogleOracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList {
	var returns GoogleOracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList
	_jsii_.Get(
		j,
		"localStandbyDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceBeginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceBeginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceScheduleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceScheduleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceScheduleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceScheduleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MemoryPerOracleComputeUnitGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MemoryTableGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryTableGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MtlsConnectionRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MtlsConnectionRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mtlsConnectionRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) NCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) NCharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nCharacterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) NextLongTermBackupTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) OpenMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) OperationsInsightsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsInsightsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) OperationsInsightsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsInsightsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PeerDbIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDbIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PermissionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) RefreshableMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) RefreshableState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ScheduledOperationDetails() GoogleOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList {
	var returns GoogleOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList
	_jsii_.Get(
		j,
		"scheduledOperationDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) SupportedCloneRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCloneRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) TotalAutoBackupStorageSizeGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalAutoBackupStorageSizeGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) UsedDataStorageSizeTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeTbs",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseAutonomousDatabase.GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference_Override(g GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseAutonomousDatabase.GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetBackupRetentionPeriodDays(val *float64) {
	if err := j.validateSetBackupRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriodDays",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetCharacterSet(val *string) {
	if err := j.validateSetCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDataStorageSizeGb(val *float64) {
	if err := j.validateSetDataStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDataStorageSizeTb(val *float64) {
	if err := j.validateSetDataStorageSizeTbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeTb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDbEdition(val *string) {
	if err := j.validateSetDbEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbEdition",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDbVersion(val *string) {
	if err := j.validateSetDbVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDbWorkload(val *string) {
	if err := j.validateSetDbWorkloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbWorkload",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseAutonomousDatabaseProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetIsAutoScalingEnabled(val interface{}) {
	if err := j.validateSetIsAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetIsStorageAutoScalingEnabled(val interface{}) {
	if err := j.validateSetIsStorageAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isStorageAutoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetMaintenanceScheduleType(val *string) {
	if err := j.validateSetMaintenanceScheduleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceScheduleType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetMtlsConnectionRequired(val interface{}) {
	if err := j.validateSetMtlsConnectionRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtlsConnectionRequired",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetNCharacterSet(val *string) {
	if err := j.validateSetNCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nCharacterSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetOperationsInsightsState(val *string) {
	if err := j.validateSetOperationsInsightsStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationsInsightsState",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetPrivateEndpointIp(val *string) {
	if err := j.validateSetPrivateEndpointIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointIp",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetPrivateEndpointLabel(val *string) {
	if err := j.validateSetPrivateEndpointLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointLabel",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PutCustomerContacts(value interface{}) {
	if err := g.validatePutCustomerContactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomerContacts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetBackupRetentionPeriodDays() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupRetentionPeriodDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetCharacterSet() {
	_jsii_.InvokeVoid(
		g,
		"resetCharacterSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetComputeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetCustomerContacts() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomerContacts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDataStorageSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStorageSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDataStorageSizeTb() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStorageSizeTb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDbEdition() {
	_jsii_.InvokeVoid(
		g,
		"resetDbEdition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDbVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetDbVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetIsAutoScalingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsAutoScalingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetIsStorageAutoScalingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsStorageAutoScalingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetMaintenanceScheduleType() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceScheduleType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetMtlsConnectionRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetMtlsConnectionRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetNCharacterSet() {
	_jsii_.InvokeVoid(
		g,
		"resetNCharacterSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetOperationsInsightsState() {
	_jsii_.InvokeVoid(
		g,
		"resetOperationsInsightsState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetPrivateEndpointIp() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateEndpointIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetPrivateEndpointLabel() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateEndpointLabel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


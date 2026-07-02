package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/netappvolume/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/netapp_volume google_netapp_volume}.
type NetappVolume interface {
	cdktf.TerraformResource
	ActiveDirectory() *string
	BackupConfig() NetappVolumeBackupConfigOutputReference
	BackupConfigInput() *NetappVolumeBackupConfig
	BlockDevices() NetappVolumeBlockDevicesList
	BlockDevicesInput() any
	CacheParameters() NetappVolumeCacheParametersOutputReference
	CacheParametersInput() *NetappVolumeCacheParameters
	CapacityGib() *string
	SetCapacityGib(val *string)
	CapacityGibInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ColdTierSizeGib() *string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreateTime() *string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	EncryptionType() *string
	ExportPolicy() NetappVolumeExportPolicyOutputReference
	ExportPolicyInput() *NetappVolumeExportPolicy
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HasReplication() cdktf.IResolvable
	HotTierSizeUsedGib() *string
	HybridReplicationParameters() NetappVolumeHybridReplicationParametersOutputReference
	HybridReplicationParametersInput() *NetappVolumeHybridReplicationParameters
	Id() *string
	SetId(val *string)
	IdInput() *string
	KerberosEnabled() any
	SetKerberosEnabled(val any)
	KerberosEnabledInput() any
	KmsConfig() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LargeCapacity() any
	SetLargeCapacity(val any)
	LargeCapacityConfig() NetappVolumeLargeCapacityConfigOutputReference
	LargeCapacityConfigInput() *NetappVolumeLargeCapacityConfig
	LargeCapacityInput() any
	LdapEnabled() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MountOptions() NetappVolumeMountOptionsList
	MultipleEndpoints() any
	SetMultipleEndpoints(val any)
	MultipleEndpointsInput() any
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	PsaRange() *string
	// Experimental.
	RawOverrides() any
	ReplicaZone() *string
	RestoreParameters() NetappVolumeRestoreParametersOutputReference
	RestoreParametersInput() *NetappVolumeRestoreParameters
	RestrictedActions() *[]*string
	SetRestrictedActions(val *[]*string)
	RestrictedActionsInput() *[]*string
	SecurityStyle() *string
	SetSecurityStyle(val *string)
	SecurityStyleInput() *string
	ServiceLevel() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
	SmbSettings() *[]*string
	SetSmbSettings(val *[]*string)
	SmbSettingsInput() *[]*string
	SnapshotDirectory() any
	SetSnapshotDirectory(val any)
	SnapshotDirectoryInput() any
	SnapshotPolicy() NetappVolumeSnapshotPolicyOutputReference
	SnapshotPolicyInput() *NetappVolumeSnapshotPolicy
	State() *string
	StateDetails() *string
	StoragePool() *string
	SetStoragePool(val *string)
	StoragePoolInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	ThroughputMibps() *float64
	SetThroughputMibps(val *float64)
	ThroughputMibpsInput() *float64
	TieringPolicy() NetappVolumeTieringPolicyOutputReference
	TieringPolicyInput() *NetappVolumeTieringPolicy
	Timeouts() NetappVolumeTimeoutsOutputReference
	TimeoutsInput() any
	UnixPermissions() *string
	SetUnixPermissions(val *string)
	UnixPermissionsInput() *string
	UsedGib() *string
	Zone() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBackupConfig(value *NetappVolumeBackupConfig)
	PutBlockDevices(value any)
	PutCacheParameters(value *NetappVolumeCacheParameters)
	PutExportPolicy(value *NetappVolumeExportPolicy)
	PutHybridReplicationParameters(value *NetappVolumeHybridReplicationParameters)
	PutLargeCapacityConfig(value *NetappVolumeLargeCapacityConfig)
	PutRestoreParameters(value *NetappVolumeRestoreParameters)
	PutSnapshotPolicy(value *NetappVolumeSnapshotPolicy)
	PutTieringPolicy(value *NetappVolumeTieringPolicy)
	PutTimeouts(value *NetappVolumeTimeouts)
	ResetBackupConfig()
	ResetBlockDevices()
	ResetCacheParameters()
	ResetDeletionPolicy()
	ResetDescription()
	ResetExportPolicy()
	ResetHybridReplicationParameters()
	ResetId()
	ResetKerberosEnabled()
	ResetLabels()
	ResetLargeCapacity()
	ResetLargeCapacityConfig()
	ResetMultipleEndpoints()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRestoreParameters()
	ResetRestrictedActions()
	ResetSecurityStyle()
	ResetShareName()
	ResetSmbSettings()
	ResetSnapshotDirectory()
	ResetSnapshotPolicy()
	ResetThroughputMibps()
	ResetTieringPolicy()
	ResetTimeouts()
	ResetUnixPermissions()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for NetappVolume
type jsiiProxy_NetappVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetappVolume) ActiveDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) BackupConfig() NetappVolumeBackupConfigOutputReference {
	var returns NetappVolumeBackupConfigOutputReference
	_jsii_.Get(
		j,
		"backupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) BackupConfigInput() *NetappVolumeBackupConfig {
	var returns *NetappVolumeBackupConfig
	_jsii_.Get(
		j,
		"backupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) BlockDevices() NetappVolumeBlockDevicesList {
	var returns NetappVolumeBlockDevicesList
	_jsii_.Get(
		j,
		"blockDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) BlockDevicesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"blockDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CacheParameters() NetappVolumeCacheParametersOutputReference {
	var returns NetappVolumeCacheParametersOutputReference
	_jsii_.Get(
		j,
		"cacheParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CacheParametersInput() *NetappVolumeCacheParameters {
	var returns *NetappVolumeCacheParameters
	_jsii_.Get(
		j,
		"cacheParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CapacityGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CapacityGibInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ColdTierSizeGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coldTierSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ExportPolicy() NetappVolumeExportPolicyOutputReference {
	var returns NetappVolumeExportPolicyOutputReference
	_jsii_.Get(
		j,
		"exportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ExportPolicyInput() *NetappVolumeExportPolicy {
	var returns *NetappVolumeExportPolicy
	_jsii_.Get(
		j,
		"exportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) HasReplication() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) HotTierSizeUsedGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hotTierSizeUsedGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) HybridReplicationParameters() NetappVolumeHybridReplicationParametersOutputReference {
	var returns NetappVolumeHybridReplicationParametersOutputReference
	_jsii_.Get(
		j,
		"hybridReplicationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) HybridReplicationParametersInput() *NetappVolumeHybridReplicationParameters {
	var returns *NetappVolumeHybridReplicationParameters
	_jsii_.Get(
		j,
		"hybridReplicationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) KerberosEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberosEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) KerberosEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberosEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) KmsConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LargeCapacity() any {
	var returns any
	_jsii_.Get(
		j,
		"largeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LargeCapacityConfig() NetappVolumeLargeCapacityConfigOutputReference {
	var returns NetappVolumeLargeCapacityConfigOutputReference
	_jsii_.Get(
		j,
		"largeCapacityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LargeCapacityConfigInput() *NetappVolumeLargeCapacityConfig {
	var returns *NetappVolumeLargeCapacityConfig
	_jsii_.Get(
		j,
		"largeCapacityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LargeCapacityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"largeCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LdapEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ldapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) MountOptions() NetappVolumeMountOptionsList {
	var returns NetappVolumeMountOptionsList
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) MultipleEndpoints() any {
	var returns any
	_jsii_.Get(
		j,
		"multipleEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) MultipleEndpointsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"multipleEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) PsaRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"psaRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ReplicaZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) RestoreParameters() NetappVolumeRestoreParametersOutputReference {
	var returns NetappVolumeRestoreParametersOutputReference
	_jsii_.Get(
		j,
		"restoreParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) RestoreParametersInput() *NetappVolumeRestoreParameters {
	var returns *NetappVolumeRestoreParameters
	_jsii_.Get(
		j,
		"restoreParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) RestrictedActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) RestrictedActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ServiceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SmbSettings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"smbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SmbSettingsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"smbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SnapshotDirectory() any {
	var returns any
	_jsii_.Get(
		j,
		"snapshotDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SnapshotDirectoryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"snapshotDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SnapshotPolicy() NetappVolumeSnapshotPolicyOutputReference {
	var returns NetappVolumeSnapshotPolicyOutputReference
	_jsii_.Get(
		j,
		"snapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SnapshotPolicyInput() *NetappVolumeSnapshotPolicy {
	var returns *NetappVolumeSnapshotPolicy
	_jsii_.Get(
		j,
		"snapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) StoragePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) StoragePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ThroughputMibps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputMibps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ThroughputMibpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputMibpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TieringPolicy() NetappVolumeTieringPolicyOutputReference {
	var returns NetappVolumeTieringPolicyOutputReference
	_jsii_.Get(
		j,
		"tieringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TieringPolicyInput() *NetappVolumeTieringPolicy {
	var returns *NetappVolumeTieringPolicy
	_jsii_.Get(
		j,
		"tieringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Timeouts() NetappVolumeTimeoutsOutputReference {
	var returns NetappVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) UnixPermissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unixPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) UnixPermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unixPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) UsedGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usedGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/netapp_volume google_netapp_volume} Resource.
func NewNetappVolume(scope constructs.Construct, id *string, config *NetappVolumeConfig) NetappVolume {
	_init_.Initialize()

	if err := validateNewNetappVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolume{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolume",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/netapp_volume google_netapp_volume} Resource.
func NewNetappVolume_Override(n NetappVolume, scope constructs.Construct, id *string, config *NetappVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolume",
		[]any{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetappVolume) SetCapacityGib(val *string) {
	if err := j.validateSetCapacityGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityGib",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetKerberosEnabled(val any) {
	if err := j.validateSetKerberosEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetLargeCapacity(val any) {
	if err := j.validateSetLargeCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"largeCapacity",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetMultipleEndpoints(val any) {
	if err := j.validateSetMultipleEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multipleEndpoints",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetRestrictedActions(val *[]*string) {
	if err := j.validateSetRestrictedActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedActions",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetShareName(val *string) {
	if err := j.validateSetShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetSmbSettings(val *[]*string) {
	if err := j.validateSetSmbSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbSettings",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetSnapshotDirectory(val any) {
	if err := j.validateSetSnapshotDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotDirectory",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetStoragePool(val *string) {
	if err := j.validateSetStoragePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePool",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetThroughputMibps(val *float64) {
	if err := j.validateSetThroughputMibpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputMibps",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetUnixPermissions(val *string) {
	if err := j.validateSetUnixPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unixPermissions",
		val,
	)
}

// Generates CDKTF code for importing a NetappVolume resource upon running "cdktf plan <stack-name>".
func NetappVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetappVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolume.NetappVolume",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NetappVolume_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateNetappVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolume.NetappVolume",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappVolume_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateNetappVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolume.NetappVolume",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappVolume_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateNetappVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolume.NetappVolume",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func NetappVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.netappVolume.NetappVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetappVolume) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (n *jsiiProxy_NetappVolume) AddOverride(path *string, value any) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]any{path, value},
	)
}

func (n *jsiiProxy_NetappVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]any{id, provider},
	)
}

func (n *jsiiProxy_NetappVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetappVolume) MoveTo(moveTarget *string, index any) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (n *jsiiProxy_NetappVolume) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetappVolume) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (n *jsiiProxy_NetappVolume) PutBackupConfig(value *NetappVolumeBackupConfig) {
	if err := n.validatePutBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putBackupConfig",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutBlockDevices(value any) {
	if err := n.validatePutBlockDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putBlockDevices",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutCacheParameters(value *NetappVolumeCacheParameters) {
	if err := n.validatePutCacheParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCacheParameters",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutExportPolicy(value *NetappVolumeExportPolicy) {
	if err := n.validatePutExportPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putExportPolicy",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutHybridReplicationParameters(value *NetappVolumeHybridReplicationParameters) {
	if err := n.validatePutHybridReplicationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHybridReplicationParameters",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutLargeCapacityConfig(value *NetappVolumeLargeCapacityConfig) {
	if err := n.validatePutLargeCapacityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLargeCapacityConfig",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutRestoreParameters(value *NetappVolumeRestoreParameters) {
	if err := n.validatePutRestoreParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putRestoreParameters",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutSnapshotPolicy(value *NetappVolumeSnapshotPolicy) {
	if err := n.validatePutSnapshotPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSnapshotPolicy",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutTieringPolicy(value *NetappVolumeTieringPolicy) {
	if err := n.validatePutTieringPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTieringPolicy",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutTimeouts(value *NetappVolumeTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]any{value},
	)
}

func (n *jsiiProxy_NetappVolume) ResetBackupConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetBackupConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetBlockDevices() {
	_jsii_.InvokeVoid(
		n,
		"resetBlockDevices",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetCacheParameters() {
	_jsii_.InvokeVoid(
		n,
		"resetCacheParameters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetExportPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetExportPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetHybridReplicationParameters() {
	_jsii_.InvokeVoid(
		n,
		"resetHybridReplicationParameters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetKerberosEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberosEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetLargeCapacity() {
	_jsii_.InvokeVoid(
		n,
		"resetLargeCapacity",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetLargeCapacityConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetLargeCapacityConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetMultipleEndpoints() {
	_jsii_.InvokeVoid(
		n,
		"resetMultipleEndpoints",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetRestoreParameters() {
	_jsii_.InvokeVoid(
		n,
		"resetRestoreParameters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetRestrictedActions() {
	_jsii_.InvokeVoid(
		n,
		"resetRestrictedActions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSecurityStyle() {
	_jsii_.InvokeVoid(
		n,
		"resetSecurityStyle",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetShareName() {
	_jsii_.InvokeVoid(
		n,
		"resetShareName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSmbSettings() {
	_jsii_.InvokeVoid(
		n,
		"resetSmbSettings",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSnapshotDirectory() {
	_jsii_.InvokeVoid(
		n,
		"resetSnapshotDirectory",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSnapshotPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetSnapshotPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetThroughputMibps() {
	_jsii_.InvokeVoid(
		n,
		"resetThroughputMibps",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetTieringPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetTieringPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetUnixPermissions() {
	_jsii_.InvokeVoid(
		n,
		"resetUnixPermissions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

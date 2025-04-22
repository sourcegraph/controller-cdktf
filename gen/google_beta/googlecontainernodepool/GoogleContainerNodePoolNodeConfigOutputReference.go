package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainernodepool/internal"
)

type GoogleContainerNodePoolNodeConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedMachineFeatures() GoogleContainerNodePoolNodeConfigAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *GoogleContainerNodePoolNodeConfigAdvancedMachineFeatures
	BootDiskKmsKey() *string
	SetBootDiskKmsKey(val *string)
	BootDiskKmsKeyInput() *string
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
	ConfidentialNodes() GoogleContainerNodePoolNodeConfigConfidentialNodesOutputReference
	ConfidentialNodesInput() *GoogleContainerNodePoolNodeConfigConfidentialNodes
	ContainerdConfig() GoogleContainerNodePoolNodeConfigContainerdConfigOutputReference
	ContainerdConfigInput() *GoogleContainerNodePoolNodeConfigContainerdConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	EffectiveTaints() GoogleContainerNodePoolNodeConfigEffectiveTaintsList
	EnableConfidentialStorage() interface{}
	SetEnableConfidentialStorage(val interface{})
	EnableConfidentialStorageInput() interface{}
	EphemeralStorageConfig() GoogleContainerNodePoolNodeConfigEphemeralStorageConfigOutputReference
	EphemeralStorageConfigInput() *GoogleContainerNodePoolNodeConfigEphemeralStorageConfig
	EphemeralStorageLocalSsdConfig() GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfigOutputReference
	EphemeralStorageLocalSsdConfigInput() *GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig
	FastSocket() GoogleContainerNodePoolNodeConfigFastSocketOutputReference
	FastSocketInput() *GoogleContainerNodePoolNodeConfigFastSocket
	// Experimental.
	Fqn() *string
	GcfsConfig() GoogleContainerNodePoolNodeConfigGcfsConfigOutputReference
	GcfsConfigInput() *GoogleContainerNodePoolNodeConfigGcfsConfig
	GuestAccelerator() GoogleContainerNodePoolNodeConfigGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Gvnic() GoogleContainerNodePoolNodeConfigGvnicOutputReference
	GvnicInput() *GoogleContainerNodePoolNodeConfigGvnic
	HostMaintenancePolicy() GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOutputReference
	HostMaintenancePolicyInput() *GoogleContainerNodePoolNodeConfigHostMaintenancePolicy
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *GoogleContainerNodePoolNodeConfig
	SetInternalValue(val *GoogleContainerNodePoolNodeConfig)
	KubeletConfig() GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference
	KubeletConfigInput() *GoogleContainerNodePoolNodeConfigKubeletConfig
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LinuxNodeConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference
	LinuxNodeConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfig
	LocalNvmeSsdBlockConfig() GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfigOutputReference
	LocalNvmeSsdBlockConfigInput() *GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig
	LocalSsdCount() *float64
	SetLocalSsdCount(val *float64)
	LocalSsdCountInput() *float64
	LocalSsdEncryptionMode() *string
	SetLocalSsdEncryptionMode(val *string)
	LocalSsdEncryptionModeInput() *string
	LoggingVariant() *string
	SetLoggingVariant(val *string)
	LoggingVariantInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MaxRunDuration() *string
	SetMaxRunDuration(val *string)
	MaxRunDurationInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	NodeGroup() *string
	SetNodeGroup(val *string)
	NodeGroupInput() *string
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	Preemptible() interface{}
	SetPreemptible(val interface{})
	PreemptibleInput() interface{}
	ReservationAffinity() GoogleContainerNodePoolNodeConfigReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleContainerNodePoolNodeConfigReservationAffinity
	ResourceLabels() *map[string]*string
	SetResourceLabels(val *map[string]*string)
	ResourceLabelsInput() *map[string]*string
	ResourceManagerTags() *map[string]*string
	SetResourceManagerTags(val *map[string]*string)
	ResourceManagerTagsInput() *map[string]*string
	SandboxConfig() GoogleContainerNodePoolNodeConfigSandboxConfigOutputReference
	SandboxConfigInput() *GoogleContainerNodePoolNodeConfigSandboxConfig
	SecondaryBootDisks() GoogleContainerNodePoolNodeConfigSecondaryBootDisksList
	SecondaryBootDisksInput() interface{}
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() GoogleContainerNodePoolNodeConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleContainerNodePoolNodeConfigShieldedInstanceConfig
	SoleTenantConfig() GoogleContainerNodePoolNodeConfigSoleTenantConfigOutputReference
	SoleTenantConfigInput() *GoogleContainerNodePoolNodeConfigSoleTenantConfig
	Spot() interface{}
	SetSpot(val interface{})
	SpotInput() interface{}
	StoragePools() *[]*string
	SetStoragePools(val *[]*string)
	StoragePoolsInput() *[]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Taint() GoogleContainerNodePoolNodeConfigTaintList
	TaintInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsNodeConfig() GoogleContainerNodePoolNodeConfigWindowsNodeConfigOutputReference
	WindowsNodeConfigInput() *GoogleContainerNodePoolNodeConfigWindowsNodeConfig
	WorkloadMetadataConfig() GoogleContainerNodePoolNodeConfigWorkloadMetadataConfigOutputReference
	WorkloadMetadataConfigInput() *GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig
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
	PutAdvancedMachineFeatures(value *GoogleContainerNodePoolNodeConfigAdvancedMachineFeatures)
	PutConfidentialNodes(value *GoogleContainerNodePoolNodeConfigConfidentialNodes)
	PutContainerdConfig(value *GoogleContainerNodePoolNodeConfigContainerdConfig)
	PutEphemeralStorageConfig(value *GoogleContainerNodePoolNodeConfigEphemeralStorageConfig)
	PutEphemeralStorageLocalSsdConfig(value *GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig)
	PutFastSocket(value *GoogleContainerNodePoolNodeConfigFastSocket)
	PutGcfsConfig(value *GoogleContainerNodePoolNodeConfigGcfsConfig)
	PutGuestAccelerator(value interface{})
	PutGvnic(value *GoogleContainerNodePoolNodeConfigGvnic)
	PutHostMaintenancePolicy(value *GoogleContainerNodePoolNodeConfigHostMaintenancePolicy)
	PutKubeletConfig(value *GoogleContainerNodePoolNodeConfigKubeletConfig)
	PutLinuxNodeConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfig)
	PutLocalNvmeSsdBlockConfig(value *GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig)
	PutReservationAffinity(value *GoogleContainerNodePoolNodeConfigReservationAffinity)
	PutSandboxConfig(value *GoogleContainerNodePoolNodeConfigSandboxConfig)
	PutSecondaryBootDisks(value interface{})
	PutShieldedInstanceConfig(value *GoogleContainerNodePoolNodeConfigShieldedInstanceConfig)
	PutSoleTenantConfig(value *GoogleContainerNodePoolNodeConfigSoleTenantConfig)
	PutTaint(value interface{})
	PutWindowsNodeConfig(value *GoogleContainerNodePoolNodeConfigWindowsNodeConfig)
	PutWorkloadMetadataConfig(value *GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig)
	ResetAdvancedMachineFeatures()
	ResetBootDiskKmsKey()
	ResetConfidentialNodes()
	ResetContainerdConfig()
	ResetDiskSizeGb()
	ResetDiskType()
	ResetEnableConfidentialStorage()
	ResetEphemeralStorageConfig()
	ResetEphemeralStorageLocalSsdConfig()
	ResetFastSocket()
	ResetGcfsConfig()
	ResetGuestAccelerator()
	ResetGvnic()
	ResetHostMaintenancePolicy()
	ResetImageType()
	ResetKubeletConfig()
	ResetLabels()
	ResetLinuxNodeConfig()
	ResetLocalNvmeSsdBlockConfig()
	ResetLocalSsdCount()
	ResetLocalSsdEncryptionMode()
	ResetLoggingVariant()
	ResetMachineType()
	ResetMaxRunDuration()
	ResetMetadata()
	ResetMinCpuPlatform()
	ResetNodeGroup()
	ResetOauthScopes()
	ResetPreemptible()
	ResetReservationAffinity()
	ResetResourceLabels()
	ResetResourceManagerTags()
	ResetSandboxConfig()
	ResetSecondaryBootDisks()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetSoleTenantConfig()
	ResetSpot()
	ResetStoragePools()
	ResetTags()
	ResetTaint()
	ResetWindowsNodeConfig()
	ResetWorkloadMetadataConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerNodePoolNodeConfigOutputReference
type jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) AdvancedMachineFeatures() GoogleContainerNodePoolNodeConfigAdvancedMachineFeaturesOutputReference {
	var returns GoogleContainerNodePoolNodeConfigAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) AdvancedMachineFeaturesInput() *GoogleContainerNodePoolNodeConfigAdvancedMachineFeatures {
	var returns *GoogleContainerNodePoolNodeConfigAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) BootDiskKmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ConfidentialNodes() GoogleContainerNodePoolNodeConfigConfidentialNodesOutputReference {
	var returns GoogleContainerNodePoolNodeConfigConfidentialNodesOutputReference
	_jsii_.Get(
		j,
		"confidentialNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ConfidentialNodesInput() *GoogleContainerNodePoolNodeConfigConfidentialNodes {
	var returns *GoogleContainerNodePoolNodeConfigConfidentialNodes
	_jsii_.Get(
		j,
		"confidentialNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ContainerdConfig() GoogleContainerNodePoolNodeConfigContainerdConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigContainerdConfigOutputReference
	_jsii_.Get(
		j,
		"containerdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ContainerdConfigInput() *GoogleContainerNodePoolNodeConfigContainerdConfig {
	var returns *GoogleContainerNodePoolNodeConfigContainerdConfig
	_jsii_.Get(
		j,
		"containerdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) EffectiveTaints() GoogleContainerNodePoolNodeConfigEffectiveTaintsList {
	var returns GoogleContainerNodePoolNodeConfigEffectiveTaintsList
	_jsii_.Get(
		j,
		"effectiveTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) EnableConfidentialStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) EnableConfidentialStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) EphemeralStorageConfig() GoogleContainerNodePoolNodeConfigEphemeralStorageConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigEphemeralStorageConfigOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) EphemeralStorageConfigInput() *GoogleContainerNodePoolNodeConfigEphemeralStorageConfig {
	var returns *GoogleContainerNodePoolNodeConfigEphemeralStorageConfig
	_jsii_.Get(
		j,
		"ephemeralStorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) EphemeralStorageLocalSsdConfig() GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfigOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) EphemeralStorageLocalSsdConfigInput() *GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig {
	var returns *GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) FastSocket() GoogleContainerNodePoolNodeConfigFastSocketOutputReference {
	var returns GoogleContainerNodePoolNodeConfigFastSocketOutputReference
	_jsii_.Get(
		j,
		"fastSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) FastSocketInput() *GoogleContainerNodePoolNodeConfigFastSocket {
	var returns *GoogleContainerNodePoolNodeConfigFastSocket
	_jsii_.Get(
		j,
		"fastSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GcfsConfig() GoogleContainerNodePoolNodeConfigGcfsConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigGcfsConfigOutputReference
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GcfsConfigInput() *GoogleContainerNodePoolNodeConfigGcfsConfig {
	var returns *GoogleContainerNodePoolNodeConfigGcfsConfig
	_jsii_.Get(
		j,
		"gcfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GuestAccelerator() GoogleContainerNodePoolNodeConfigGuestAcceleratorList {
	var returns GoogleContainerNodePoolNodeConfigGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Gvnic() GoogleContainerNodePoolNodeConfigGvnicOutputReference {
	var returns GoogleContainerNodePoolNodeConfigGvnicOutputReference
	_jsii_.Get(
		j,
		"gvnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GvnicInput() *GoogleContainerNodePoolNodeConfigGvnic {
	var returns *GoogleContainerNodePoolNodeConfigGvnic
	_jsii_.Get(
		j,
		"gvnicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) HostMaintenancePolicy() GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOutputReference {
	var returns GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"hostMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) HostMaintenancePolicyInput() *GoogleContainerNodePoolNodeConfigHostMaintenancePolicy {
	var returns *GoogleContainerNodePoolNodeConfigHostMaintenancePolicy
	_jsii_.Get(
		j,
		"hostMaintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) InternalValue() *GoogleContainerNodePoolNodeConfig {
	var returns *GoogleContainerNodePoolNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) KubeletConfig() GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) KubeletConfigInput() *GoogleContainerNodePoolNodeConfigKubeletConfig {
	var returns *GoogleContainerNodePoolNodeConfigKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LinuxNodeConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference
	_jsii_.Get(
		j,
		"linuxNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LinuxNodeConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfig {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"linuxNodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LocalNvmeSsdBlockConfig() GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfigOutputReference
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LocalNvmeSsdBlockConfigInput() *GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig {
	var returns *GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LocalSsdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LocalSsdEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LocalSsdEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) LoggingVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) MaxRunDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRunDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) MaxRunDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRunDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) NodeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) NodeGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ReservationAffinity() GoogleContainerNodePoolNodeConfigReservationAffinityOutputReference {
	var returns GoogleContainerNodePoolNodeConfigReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ReservationAffinityInput() *GoogleContainerNodePoolNodeConfigReservationAffinity {
	var returns *GoogleContainerNodePoolNodeConfigReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResourceLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResourceLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResourceManagerTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResourceManagerTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) SandboxConfig() GoogleContainerNodePoolNodeConfigSandboxConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigSandboxConfigOutputReference
	_jsii_.Get(
		j,
		"sandboxConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) SandboxConfigInput() *GoogleContainerNodePoolNodeConfigSandboxConfig {
	var returns *GoogleContainerNodePoolNodeConfigSandboxConfig
	_jsii_.Get(
		j,
		"sandboxConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) SecondaryBootDisks() GoogleContainerNodePoolNodeConfigSecondaryBootDisksList {
	var returns GoogleContainerNodePoolNodeConfigSecondaryBootDisksList
	_jsii_.Get(
		j,
		"secondaryBootDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) SecondaryBootDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryBootDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ShieldedInstanceConfig() GoogleContainerNodePoolNodeConfigShieldedInstanceConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ShieldedInstanceConfigInput() *GoogleContainerNodePoolNodeConfigShieldedInstanceConfig {
	var returns *GoogleContainerNodePoolNodeConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) SoleTenantConfig() GoogleContainerNodePoolNodeConfigSoleTenantConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigSoleTenantConfigOutputReference
	_jsii_.Get(
		j,
		"soleTenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) SoleTenantConfigInput() *GoogleContainerNodePoolNodeConfigSoleTenantConfig {
	var returns *GoogleContainerNodePoolNodeConfigSoleTenantConfig
	_jsii_.Get(
		j,
		"soleTenantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Spot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) SpotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) StoragePools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storagePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) StoragePoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storagePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Taint() GoogleContainerNodePoolNodeConfigTaintList {
	var returns GoogleContainerNodePoolNodeConfigTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) WindowsNodeConfig() GoogleContainerNodePoolNodeConfigWindowsNodeConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigWindowsNodeConfigOutputReference
	_jsii_.Get(
		j,
		"windowsNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) WindowsNodeConfigInput() *GoogleContainerNodePoolNodeConfigWindowsNodeConfig {
	var returns *GoogleContainerNodePoolNodeConfigWindowsNodeConfig
	_jsii_.Get(
		j,
		"windowsNodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) WorkloadMetadataConfig() GoogleContainerNodePoolNodeConfigWorkloadMetadataConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigWorkloadMetadataConfigOutputReference
	_jsii_.Get(
		j,
		"workloadMetadataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) WorkloadMetadataConfigInput() *GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig {
	var returns *GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig
	_jsii_.Get(
		j,
		"workloadMetadataConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerNodePoolNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerNodePoolNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerNodePoolNodeConfigOutputReference_Override(g GoogleContainerNodePoolNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetBootDiskKmsKey(val *string) {
	if err := j.validateSetBootDiskKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskKmsKey",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetEnableConfidentialStorage(val interface{}) {
	if err := j.validateSetEnableConfidentialStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidentialStorage",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetInternalValue(val *GoogleContainerNodePoolNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetLocalSsdCount(val *float64) {
	if err := j.validateSetLocalSsdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetLocalSsdEncryptionMode(val *string) {
	if err := j.validateSetLocalSsdEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetLoggingVariant(val *string) {
	if err := j.validateSetLoggingVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingVariant",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetMaxRunDuration(val *string) {
	if err := j.validateSetMaxRunDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRunDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetNodeGroup(val *string) {
	if err := j.validateSetNodeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetOauthScopes(val *[]*string) {
	if err := j.validateSetOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetResourceLabels(val *map[string]*string) {
	if err := j.validateSetResourceLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetResourceManagerTags(val *map[string]*string) {
	if err := j.validateSetResourceManagerTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceManagerTags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetSpot(val interface{}) {
	if err := j.validateSetSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spot",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetStoragePools(val *[]*string) {
	if err := j.validateSetStoragePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePools",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutAdvancedMachineFeatures(value *GoogleContainerNodePoolNodeConfigAdvancedMachineFeatures) {
	if err := g.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutConfidentialNodes(value *GoogleContainerNodePoolNodeConfigConfidentialNodes) {
	if err := g.validatePutConfidentialNodesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialNodes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutContainerdConfig(value *GoogleContainerNodePoolNodeConfigContainerdConfig) {
	if err := g.validatePutContainerdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerdConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutEphemeralStorageConfig(value *GoogleContainerNodePoolNodeConfigEphemeralStorageConfig) {
	if err := g.validatePutEphemeralStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralStorageConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutEphemeralStorageLocalSsdConfig(value *GoogleContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig) {
	if err := g.validatePutEphemeralStorageLocalSsdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralStorageLocalSsdConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutFastSocket(value *GoogleContainerNodePoolNodeConfigFastSocket) {
	if err := g.validatePutFastSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFastSocket",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutGcfsConfig(value *GoogleContainerNodePoolNodeConfigGcfsConfig) {
	if err := g.validatePutGcfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcfsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutGuestAccelerator(value interface{}) {
	if err := g.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutGvnic(value *GoogleContainerNodePoolNodeConfigGvnic) {
	if err := g.validatePutGvnicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGvnic",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutHostMaintenancePolicy(value *GoogleContainerNodePoolNodeConfigHostMaintenancePolicy) {
	if err := g.validatePutHostMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHostMaintenancePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutKubeletConfig(value *GoogleContainerNodePoolNodeConfigKubeletConfig) {
	if err := g.validatePutKubeletConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutLinuxNodeConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfig) {
	if err := g.validatePutLinuxNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutLocalNvmeSsdBlockConfig(value *GoogleContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig) {
	if err := g.validatePutLocalNvmeSsdBlockConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocalNvmeSsdBlockConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutReservationAffinity(value *GoogleContainerNodePoolNodeConfigReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutSandboxConfig(value *GoogleContainerNodePoolNodeConfigSandboxConfig) {
	if err := g.validatePutSandboxConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSandboxConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutSecondaryBootDisks(value interface{}) {
	if err := g.validatePutSecondaryBootDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecondaryBootDisks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutShieldedInstanceConfig(value *GoogleContainerNodePoolNodeConfigShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutSoleTenantConfig(value *GoogleContainerNodePoolNodeConfigSoleTenantConfig) {
	if err := g.validatePutSoleTenantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoleTenantConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutTaint(value interface{}) {
	if err := g.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTaint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutWindowsNodeConfig(value *GoogleContainerNodePoolNodeConfigWindowsNodeConfig) {
	if err := g.validatePutWindowsNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowsNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) PutWorkloadMetadataConfig(value *GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig) {
	if err := g.validatePutWorkloadMetadataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkloadMetadataConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetBootDiskKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetConfidentialNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetContainerdConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerdConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetEnableConfidentialStorage() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableConfidentialStorage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetEphemeralStorageConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralStorageConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetEphemeralStorageLocalSsdConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralStorageLocalSsdConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetFastSocket() {
	_jsii_.InvokeVoid(
		g,
		"resetFastSocket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetGcfsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGcfsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetGvnic() {
	_jsii_.InvokeVoid(
		g,
		"resetGvnic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetHostMaintenancePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetHostMaintenancePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		g,
		"resetImageType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetLinuxNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetLocalNvmeSsdBlockConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalNvmeSsdBlockConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetLocalSsdCount() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetLocalSsdEncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdEncryptionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetLoggingVariant() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingVariant",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetMaxRunDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRunDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetNodeGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetResourceLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetSandboxConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSandboxConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetSecondaryBootDisks() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryBootDisks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetSoleTenantConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSoleTenantConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetSpot() {
	_jsii_.InvokeVoid(
		g,
		"resetSpot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetStoragePools() {
	_jsii_.InvokeVoid(
		g,
		"resetStoragePools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetTaint() {
	_jsii_.InvokeVoid(
		g,
		"resetTaint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetWindowsNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ResetWorkloadMetadataConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadMetadataConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


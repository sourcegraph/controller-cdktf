package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodePoolNodeConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedMachineFeatures() GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeatures
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
	ConfidentialNodes() GoogleContainerClusterNodePoolNodeConfigConfidentialNodesOutputReference
	ConfidentialNodesInput() *GoogleContainerClusterNodePoolNodeConfigConfidentialNodes
	ContainerdConfig() GoogleContainerClusterNodePoolNodeConfigContainerdConfigOutputReference
	ContainerdConfigInput() *GoogleContainerClusterNodePoolNodeConfigContainerdConfig
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
	EffectiveTaints() GoogleContainerClusterNodePoolNodeConfigEffectiveTaintsList
	EnableConfidentialStorage() interface{}
	SetEnableConfidentialStorage(val interface{})
	EnableConfidentialStorageInput() interface{}
	EphemeralStorageConfig() GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfigOutputReference
	EphemeralStorageConfigInput() *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfig
	EphemeralStorageLocalSsdConfig() GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfigOutputReference
	EphemeralStorageLocalSsdConfigInput() *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfig
	FastSocket() GoogleContainerClusterNodePoolNodeConfigFastSocketOutputReference
	FastSocketInput() *GoogleContainerClusterNodePoolNodeConfigFastSocket
	FlexStart() interface{}
	SetFlexStart(val interface{})
	FlexStartInput() interface{}
	// Experimental.
	Fqn() *string
	GcfsConfig() GoogleContainerClusterNodePoolNodeConfigGcfsConfigOutputReference
	GcfsConfigInput() *GoogleContainerClusterNodePoolNodeConfigGcfsConfig
	GuestAccelerator() GoogleContainerClusterNodePoolNodeConfigGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Gvnic() GoogleContainerClusterNodePoolNodeConfigGvnicOutputReference
	GvnicInput() *GoogleContainerClusterNodePoolNodeConfigGvnic
	HostMaintenancePolicy() GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicyOutputReference
	HostMaintenancePolicyInput() *GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicy
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *GoogleContainerClusterNodePoolNodeConfig
	SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfig)
	KubeletConfig() GoogleContainerClusterNodePoolNodeConfigKubeletConfigOutputReference
	KubeletConfigInput() *GoogleContainerClusterNodePoolNodeConfigKubeletConfig
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LinuxNodeConfig() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference
	LinuxNodeConfigInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig
	LocalNvmeSsdBlockConfig() GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigOutputReference
	LocalNvmeSsdBlockConfigInput() *GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig
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
	ReservationAffinity() GoogleContainerClusterNodePoolNodeConfigReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleContainerClusterNodePoolNodeConfigReservationAffinity
	ResourceLabels() *map[string]*string
	SetResourceLabels(val *map[string]*string)
	ResourceLabelsInput() *map[string]*string
	ResourceManagerTags() *map[string]*string
	SetResourceManagerTags(val *map[string]*string)
	ResourceManagerTagsInput() *map[string]*string
	SandboxConfig() GoogleContainerClusterNodePoolNodeConfigSandboxConfigOutputReference
	SandboxConfigInput() *GoogleContainerClusterNodePoolNodeConfigSandboxConfig
	SecondaryBootDisks() GoogleContainerClusterNodePoolNodeConfigSecondaryBootDisksList
	SecondaryBootDisksInput() interface{}
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfig
	SoleTenantConfig() GoogleContainerClusterNodePoolNodeConfigSoleTenantConfigOutputReference
	SoleTenantConfigInput() *GoogleContainerClusterNodePoolNodeConfigSoleTenantConfig
	Spot() interface{}
	SetSpot(val interface{})
	SpotInput() interface{}
	StoragePools() *[]*string
	SetStoragePools(val *[]*string)
	StoragePoolsInput() *[]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Taint() GoogleContainerClusterNodePoolNodeConfigTaintList
	TaintInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsNodeConfig() GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfigOutputReference
	WindowsNodeConfigInput() *GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfig
	WorkloadMetadataConfig() GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfigOutputReference
	WorkloadMetadataConfigInput() *GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfig
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
	PutAdvancedMachineFeatures(value *GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeatures)
	PutConfidentialNodes(value *GoogleContainerClusterNodePoolNodeConfigConfidentialNodes)
	PutContainerdConfig(value *GoogleContainerClusterNodePoolNodeConfigContainerdConfig)
	PutEphemeralStorageConfig(value *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfig)
	PutEphemeralStorageLocalSsdConfig(value *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfig)
	PutFastSocket(value *GoogleContainerClusterNodePoolNodeConfigFastSocket)
	PutGcfsConfig(value *GoogleContainerClusterNodePoolNodeConfigGcfsConfig)
	PutGuestAccelerator(value interface{})
	PutGvnic(value *GoogleContainerClusterNodePoolNodeConfigGvnic)
	PutHostMaintenancePolicy(value *GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicy)
	PutKubeletConfig(value *GoogleContainerClusterNodePoolNodeConfigKubeletConfig)
	PutLinuxNodeConfig(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig)
	PutLocalNvmeSsdBlockConfig(value *GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig)
	PutReservationAffinity(value *GoogleContainerClusterNodePoolNodeConfigReservationAffinity)
	PutSandboxConfig(value *GoogleContainerClusterNodePoolNodeConfigSandboxConfig)
	PutSecondaryBootDisks(value interface{})
	PutShieldedInstanceConfig(value *GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfig)
	PutSoleTenantConfig(value *GoogleContainerClusterNodePoolNodeConfigSoleTenantConfig)
	PutTaint(value interface{})
	PutWindowsNodeConfig(value *GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfig)
	PutWorkloadMetadataConfig(value *GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfig)
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
	ResetFlexStart()
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

// The jsii proxy struct for GoogleContainerClusterNodePoolNodeConfigOutputReference
type jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) AdvancedMachineFeatures() GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeaturesOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) AdvancedMachineFeaturesInput() *GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeatures {
	var returns *GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) BootDiskKmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ConfidentialNodes() GoogleContainerClusterNodePoolNodeConfigConfidentialNodesOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigConfidentialNodesOutputReference
	_jsii_.Get(
		j,
		"confidentialNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ConfidentialNodesInput() *GoogleContainerClusterNodePoolNodeConfigConfidentialNodes {
	var returns *GoogleContainerClusterNodePoolNodeConfigConfidentialNodes
	_jsii_.Get(
		j,
		"confidentialNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ContainerdConfig() GoogleContainerClusterNodePoolNodeConfigContainerdConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigContainerdConfigOutputReference
	_jsii_.Get(
		j,
		"containerdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ContainerdConfigInput() *GoogleContainerClusterNodePoolNodeConfigContainerdConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigContainerdConfig
	_jsii_.Get(
		j,
		"containerdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) EffectiveTaints() GoogleContainerClusterNodePoolNodeConfigEffectiveTaintsList {
	var returns GoogleContainerClusterNodePoolNodeConfigEffectiveTaintsList
	_jsii_.Get(
		j,
		"effectiveTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) EnableConfidentialStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) EnableConfidentialStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) EphemeralStorageConfig() GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfigOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) EphemeralStorageConfigInput() *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfig
	_jsii_.Get(
		j,
		"ephemeralStorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) EphemeralStorageLocalSsdConfig() GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfigOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) EphemeralStorageLocalSsdConfigInput() *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfig
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) FastSocket() GoogleContainerClusterNodePoolNodeConfigFastSocketOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigFastSocketOutputReference
	_jsii_.Get(
		j,
		"fastSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) FastSocketInput() *GoogleContainerClusterNodePoolNodeConfigFastSocket {
	var returns *GoogleContainerClusterNodePoolNodeConfigFastSocket
	_jsii_.Get(
		j,
		"fastSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) FlexStart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flexStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) FlexStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flexStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GcfsConfig() GoogleContainerClusterNodePoolNodeConfigGcfsConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigGcfsConfigOutputReference
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GcfsConfigInput() *GoogleContainerClusterNodePoolNodeConfigGcfsConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigGcfsConfig
	_jsii_.Get(
		j,
		"gcfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GuestAccelerator() GoogleContainerClusterNodePoolNodeConfigGuestAcceleratorList {
	var returns GoogleContainerClusterNodePoolNodeConfigGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Gvnic() GoogleContainerClusterNodePoolNodeConfigGvnicOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigGvnicOutputReference
	_jsii_.Get(
		j,
		"gvnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GvnicInput() *GoogleContainerClusterNodePoolNodeConfigGvnic {
	var returns *GoogleContainerClusterNodePoolNodeConfigGvnic
	_jsii_.Get(
		j,
		"gvnicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) HostMaintenancePolicy() GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicyOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"hostMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) HostMaintenancePolicyInput() *GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicy {
	var returns *GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicy
	_jsii_.Get(
		j,
		"hostMaintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) InternalValue() *GoogleContainerClusterNodePoolNodeConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) KubeletConfig() GoogleContainerClusterNodePoolNodeConfigKubeletConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) KubeletConfigInput() *GoogleContainerClusterNodePoolNodeConfigKubeletConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LinuxNodeConfig() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference
	_jsii_.Get(
		j,
		"linuxNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LinuxNodeConfigInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"linuxNodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LocalNvmeSsdBlockConfig() GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigOutputReference
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LocalNvmeSsdBlockConfigInput() *GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LocalSsdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LocalSsdEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LocalSsdEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) LoggingVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) MaxRunDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRunDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) MaxRunDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRunDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) NodeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) NodeGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ReservationAffinity() GoogleContainerClusterNodePoolNodeConfigReservationAffinityOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ReservationAffinityInput() *GoogleContainerClusterNodePoolNodeConfigReservationAffinity {
	var returns *GoogleContainerClusterNodePoolNodeConfigReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResourceLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResourceLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResourceManagerTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResourceManagerTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) SandboxConfig() GoogleContainerClusterNodePoolNodeConfigSandboxConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigSandboxConfigOutputReference
	_jsii_.Get(
		j,
		"sandboxConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) SandboxConfigInput() *GoogleContainerClusterNodePoolNodeConfigSandboxConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigSandboxConfig
	_jsii_.Get(
		j,
		"sandboxConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) SecondaryBootDisks() GoogleContainerClusterNodePoolNodeConfigSecondaryBootDisksList {
	var returns GoogleContainerClusterNodePoolNodeConfigSecondaryBootDisksList
	_jsii_.Get(
		j,
		"secondaryBootDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) SecondaryBootDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryBootDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ShieldedInstanceConfig() GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ShieldedInstanceConfigInput() *GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) SoleTenantConfig() GoogleContainerClusterNodePoolNodeConfigSoleTenantConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigSoleTenantConfigOutputReference
	_jsii_.Get(
		j,
		"soleTenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) SoleTenantConfigInput() *GoogleContainerClusterNodePoolNodeConfigSoleTenantConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigSoleTenantConfig
	_jsii_.Get(
		j,
		"soleTenantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Spot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) SpotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) StoragePools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storagePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) StoragePoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storagePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Taint() GoogleContainerClusterNodePoolNodeConfigTaintList {
	var returns GoogleContainerClusterNodePoolNodeConfigTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) WindowsNodeConfig() GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfigOutputReference
	_jsii_.Get(
		j,
		"windowsNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) WindowsNodeConfigInput() *GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfig
	_jsii_.Get(
		j,
		"windowsNodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) WorkloadMetadataConfig() GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfigOutputReference
	_jsii_.Get(
		j,
		"workloadMetadataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) WorkloadMetadataConfigInput() *GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfig
	_jsii_.Get(
		j,
		"workloadMetadataConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodePoolNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodePoolNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodePoolNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodePoolNodeConfigOutputReference_Override(g GoogleContainerClusterNodePoolNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetBootDiskKmsKey(val *string) {
	if err := j.validateSetBootDiskKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskKmsKey",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetEnableConfidentialStorage(val interface{}) {
	if err := j.validateSetEnableConfidentialStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidentialStorage",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetFlexStart(val interface{}) {
	if err := j.validateSetFlexStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flexStart",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetLocalSsdCount(val *float64) {
	if err := j.validateSetLocalSsdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetLocalSsdEncryptionMode(val *string) {
	if err := j.validateSetLocalSsdEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetLoggingVariant(val *string) {
	if err := j.validateSetLoggingVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingVariant",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetMaxRunDuration(val *string) {
	if err := j.validateSetMaxRunDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRunDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetNodeGroup(val *string) {
	if err := j.validateSetNodeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetOauthScopes(val *[]*string) {
	if err := j.validateSetOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetResourceLabels(val *map[string]*string) {
	if err := j.validateSetResourceLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetResourceManagerTags(val *map[string]*string) {
	if err := j.validateSetResourceManagerTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceManagerTags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetSpot(val interface{}) {
	if err := j.validateSetSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spot",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetStoragePools(val *[]*string) {
	if err := j.validateSetStoragePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePools",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutAdvancedMachineFeatures(value *GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeatures) {
	if err := g.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutConfidentialNodes(value *GoogleContainerClusterNodePoolNodeConfigConfidentialNodes) {
	if err := g.validatePutConfidentialNodesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialNodes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutContainerdConfig(value *GoogleContainerClusterNodePoolNodeConfigContainerdConfig) {
	if err := g.validatePutContainerdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerdConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutEphemeralStorageConfig(value *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageConfig) {
	if err := g.validatePutEphemeralStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralStorageConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutEphemeralStorageLocalSsdConfig(value *GoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfig) {
	if err := g.validatePutEphemeralStorageLocalSsdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralStorageLocalSsdConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutFastSocket(value *GoogleContainerClusterNodePoolNodeConfigFastSocket) {
	if err := g.validatePutFastSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFastSocket",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutGcfsConfig(value *GoogleContainerClusterNodePoolNodeConfigGcfsConfig) {
	if err := g.validatePutGcfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcfsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutGuestAccelerator(value interface{}) {
	if err := g.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutGvnic(value *GoogleContainerClusterNodePoolNodeConfigGvnic) {
	if err := g.validatePutGvnicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGvnic",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutHostMaintenancePolicy(value *GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicy) {
	if err := g.validatePutHostMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHostMaintenancePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutKubeletConfig(value *GoogleContainerClusterNodePoolNodeConfigKubeletConfig) {
	if err := g.validatePutKubeletConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutLinuxNodeConfig(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig) {
	if err := g.validatePutLinuxNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutLocalNvmeSsdBlockConfig(value *GoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig) {
	if err := g.validatePutLocalNvmeSsdBlockConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocalNvmeSsdBlockConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutReservationAffinity(value *GoogleContainerClusterNodePoolNodeConfigReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutSandboxConfig(value *GoogleContainerClusterNodePoolNodeConfigSandboxConfig) {
	if err := g.validatePutSandboxConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSandboxConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutSecondaryBootDisks(value interface{}) {
	if err := g.validatePutSecondaryBootDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecondaryBootDisks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutShieldedInstanceConfig(value *GoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutSoleTenantConfig(value *GoogleContainerClusterNodePoolNodeConfigSoleTenantConfig) {
	if err := g.validatePutSoleTenantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoleTenantConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutTaint(value interface{}) {
	if err := g.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTaint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutWindowsNodeConfig(value *GoogleContainerClusterNodePoolNodeConfigWindowsNodeConfig) {
	if err := g.validatePutWindowsNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowsNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) PutWorkloadMetadataConfig(value *GoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfig) {
	if err := g.validatePutWorkloadMetadataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkloadMetadataConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetBootDiskKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetConfidentialNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetContainerdConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerdConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetEnableConfidentialStorage() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableConfidentialStorage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetEphemeralStorageConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralStorageConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetEphemeralStorageLocalSsdConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralStorageLocalSsdConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetFastSocket() {
	_jsii_.InvokeVoid(
		g,
		"resetFastSocket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetFlexStart() {
	_jsii_.InvokeVoid(
		g,
		"resetFlexStart",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetGcfsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGcfsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetGvnic() {
	_jsii_.InvokeVoid(
		g,
		"resetGvnic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetHostMaintenancePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetHostMaintenancePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		g,
		"resetImageType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetLinuxNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetLocalNvmeSsdBlockConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalNvmeSsdBlockConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetLocalSsdCount() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetLocalSsdEncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdEncryptionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetLoggingVariant() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingVariant",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetMaxRunDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRunDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetNodeGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetResourceLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetSandboxConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSandboxConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetSecondaryBootDisks() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryBootDisks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetSoleTenantConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSoleTenantConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetSpot() {
	_jsii_.InvokeVoid(
		g,
		"resetSpot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetStoragePools() {
	_jsii_.InvokeVoid(
		g,
		"resetStoragePools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetTaint() {
	_jsii_.InvokeVoid(
		g,
		"resetTaint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetWindowsNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ResetWorkloadMetadataConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadMetadataConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


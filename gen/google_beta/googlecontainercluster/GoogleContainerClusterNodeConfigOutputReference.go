package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodeConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedMachineFeatures() GoogleContainerClusterNodeConfigAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *GoogleContainerClusterNodeConfigAdvancedMachineFeatures
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
	EphemeralStorageConfig() GoogleContainerClusterNodeConfigEphemeralStorageConfigOutputReference
	EphemeralStorageConfigInput() *GoogleContainerClusterNodeConfigEphemeralStorageConfig
	EphemeralStorageLocalSsdConfig() GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfigOutputReference
	EphemeralStorageLocalSsdConfigInput() *GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfig
	// Experimental.
	Fqn() *string
	GcfsConfig() GoogleContainerClusterNodeConfigGcfsConfigOutputReference
	GcfsConfigInput() *GoogleContainerClusterNodeConfigGcfsConfig
	GuestAccelerator() GoogleContainerClusterNodeConfigGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Gvnic() GoogleContainerClusterNodeConfigGvnicOutputReference
	GvnicInput() *GoogleContainerClusterNodeConfigGvnic
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *GoogleContainerClusterNodeConfig
	SetInternalValue(val *GoogleContainerClusterNodeConfig)
	KubeletConfig() GoogleContainerClusterNodeConfigKubeletConfigOutputReference
	KubeletConfigInput() *GoogleContainerClusterNodeConfigKubeletConfig
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LinuxNodeConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference
	LinuxNodeConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfig
	LocalNvmeSsdBlockConfig() GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfigOutputReference
	LocalNvmeSsdBlockConfigInput() *GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfig
	LocalSsdCount() *float64
	SetLocalSsdCount(val *float64)
	LocalSsdCountInput() *float64
	LoggingVariant() *string
	SetLoggingVariant(val *string)
	LoggingVariantInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
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
	ReservationAffinity() GoogleContainerClusterNodeConfigReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleContainerClusterNodeConfigReservationAffinity
	ResourceLabels() *map[string]*string
	SetResourceLabels(val *map[string]*string)
	ResourceLabelsInput() *map[string]*string
	SandboxConfig() GoogleContainerClusterNodeConfigSandboxConfigOutputReference
	SandboxConfigInput() *GoogleContainerClusterNodeConfigSandboxConfig
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() GoogleContainerClusterNodeConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleContainerClusterNodeConfigShieldedInstanceConfig
	SoleTenantConfig() GoogleContainerClusterNodeConfigSoleTenantConfigOutputReference
	SoleTenantConfigInput() *GoogleContainerClusterNodeConfigSoleTenantConfig
	Spot() interface{}
	SetSpot(val interface{})
	SpotInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Taint() GoogleContainerClusterNodeConfigTaintList
	TaintInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkloadMetadataConfig() GoogleContainerClusterNodeConfigWorkloadMetadataConfigOutputReference
	WorkloadMetadataConfigInput() *GoogleContainerClusterNodeConfigWorkloadMetadataConfig
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
	PutAdvancedMachineFeatures(value *GoogleContainerClusterNodeConfigAdvancedMachineFeatures)
	PutEphemeralStorageConfig(value *GoogleContainerClusterNodeConfigEphemeralStorageConfig)
	PutEphemeralStorageLocalSsdConfig(value *GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfig)
	PutGcfsConfig(value *GoogleContainerClusterNodeConfigGcfsConfig)
	PutGuestAccelerator(value interface{})
	PutGvnic(value *GoogleContainerClusterNodeConfigGvnic)
	PutKubeletConfig(value *GoogleContainerClusterNodeConfigKubeletConfig)
	PutLinuxNodeConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfig)
	PutLocalNvmeSsdBlockConfig(value *GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfig)
	PutReservationAffinity(value *GoogleContainerClusterNodeConfigReservationAffinity)
	PutSandboxConfig(value *GoogleContainerClusterNodeConfigSandboxConfig)
	PutShieldedInstanceConfig(value *GoogleContainerClusterNodeConfigShieldedInstanceConfig)
	PutSoleTenantConfig(value *GoogleContainerClusterNodeConfigSoleTenantConfig)
	PutTaint(value interface{})
	PutWorkloadMetadataConfig(value *GoogleContainerClusterNodeConfigWorkloadMetadataConfig)
	ResetAdvancedMachineFeatures()
	ResetBootDiskKmsKey()
	ResetDiskSizeGb()
	ResetDiskType()
	ResetEphemeralStorageConfig()
	ResetEphemeralStorageLocalSsdConfig()
	ResetGcfsConfig()
	ResetGuestAccelerator()
	ResetGvnic()
	ResetImageType()
	ResetKubeletConfig()
	ResetLabels()
	ResetLinuxNodeConfig()
	ResetLocalNvmeSsdBlockConfig()
	ResetLocalSsdCount()
	ResetLoggingVariant()
	ResetMachineType()
	ResetMetadata()
	ResetMinCpuPlatform()
	ResetNodeGroup()
	ResetOauthScopes()
	ResetPreemptible()
	ResetReservationAffinity()
	ResetResourceLabels()
	ResetSandboxConfig()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetSoleTenantConfig()
	ResetSpot()
	ResetTags()
	ResetTaint()
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

// The jsii proxy struct for GoogleContainerClusterNodeConfigOutputReference
type jsiiProxy_GoogleContainerClusterNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) AdvancedMachineFeatures() GoogleContainerClusterNodeConfigAdvancedMachineFeaturesOutputReference {
	var returns GoogleContainerClusterNodeConfigAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) AdvancedMachineFeaturesInput() *GoogleContainerClusterNodeConfigAdvancedMachineFeatures {
	var returns *GoogleContainerClusterNodeConfigAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) BootDiskKmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) EphemeralStorageConfig() GoogleContainerClusterNodeConfigEphemeralStorageConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigEphemeralStorageConfigOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) EphemeralStorageConfigInput() *GoogleContainerClusterNodeConfigEphemeralStorageConfig {
	var returns *GoogleContainerClusterNodeConfigEphemeralStorageConfig
	_jsii_.Get(
		j,
		"ephemeralStorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) EphemeralStorageLocalSsdConfig() GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfigOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) EphemeralStorageLocalSsdConfigInput() *GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfig {
	var returns *GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfig
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GcfsConfig() GoogleContainerClusterNodeConfigGcfsConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigGcfsConfigOutputReference
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GcfsConfigInput() *GoogleContainerClusterNodeConfigGcfsConfig {
	var returns *GoogleContainerClusterNodeConfigGcfsConfig
	_jsii_.Get(
		j,
		"gcfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GuestAccelerator() GoogleContainerClusterNodeConfigGuestAcceleratorList {
	var returns GoogleContainerClusterNodeConfigGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Gvnic() GoogleContainerClusterNodeConfigGvnicOutputReference {
	var returns GoogleContainerClusterNodeConfigGvnicOutputReference
	_jsii_.Get(
		j,
		"gvnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GvnicInput() *GoogleContainerClusterNodeConfigGvnic {
	var returns *GoogleContainerClusterNodeConfigGvnic
	_jsii_.Get(
		j,
		"gvnicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) InternalValue() *GoogleContainerClusterNodeConfig {
	var returns *GoogleContainerClusterNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) KubeletConfig() GoogleContainerClusterNodeConfigKubeletConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) KubeletConfigInput() *GoogleContainerClusterNodeConfigKubeletConfig {
	var returns *GoogleContainerClusterNodeConfigKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LinuxNodeConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference
	_jsii_.Get(
		j,
		"linuxNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LinuxNodeConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfig {
	var returns *GoogleContainerClusterNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"linuxNodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LocalNvmeSsdBlockConfig() GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfigOutputReference
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LocalNvmeSsdBlockConfigInput() *GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfig {
	var returns *GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfig
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LocalSsdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) LoggingVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) NodeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) NodeGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ReservationAffinity() GoogleContainerClusterNodeConfigReservationAffinityOutputReference {
	var returns GoogleContainerClusterNodeConfigReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ReservationAffinityInput() *GoogleContainerClusterNodeConfigReservationAffinity {
	var returns *GoogleContainerClusterNodeConfigReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResourceLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResourceLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) SandboxConfig() GoogleContainerClusterNodeConfigSandboxConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigSandboxConfigOutputReference
	_jsii_.Get(
		j,
		"sandboxConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) SandboxConfigInput() *GoogleContainerClusterNodeConfigSandboxConfig {
	var returns *GoogleContainerClusterNodeConfigSandboxConfig
	_jsii_.Get(
		j,
		"sandboxConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ShieldedInstanceConfig() GoogleContainerClusterNodeConfigShieldedInstanceConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ShieldedInstanceConfigInput() *GoogleContainerClusterNodeConfigShieldedInstanceConfig {
	var returns *GoogleContainerClusterNodeConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) SoleTenantConfig() GoogleContainerClusterNodeConfigSoleTenantConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigSoleTenantConfigOutputReference
	_jsii_.Get(
		j,
		"soleTenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) SoleTenantConfigInput() *GoogleContainerClusterNodeConfigSoleTenantConfig {
	var returns *GoogleContainerClusterNodeConfigSoleTenantConfig
	_jsii_.Get(
		j,
		"soleTenantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Spot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) SpotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Taint() GoogleContainerClusterNodeConfigTaintList {
	var returns GoogleContainerClusterNodeConfigTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) WorkloadMetadataConfig() GoogleContainerClusterNodeConfigWorkloadMetadataConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigWorkloadMetadataConfigOutputReference
	_jsii_.Get(
		j,
		"workloadMetadataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) WorkloadMetadataConfigInput() *GoogleContainerClusterNodeConfigWorkloadMetadataConfig {
	var returns *GoogleContainerClusterNodeConfigWorkloadMetadataConfig
	_jsii_.Get(
		j,
		"workloadMetadataConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodeConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodeConfigOutputReference_Override(g GoogleContainerClusterNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetBootDiskKmsKey(val *string) {
	if err := j.validateSetBootDiskKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskKmsKey",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetInternalValue(val *GoogleContainerClusterNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetLocalSsdCount(val *float64) {
	if err := j.validateSetLocalSsdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetLoggingVariant(val *string) {
	if err := j.validateSetLoggingVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingVariant",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetNodeGroup(val *string) {
	if err := j.validateSetNodeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetOauthScopes(val *[]*string) {
	if err := j.validateSetOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetResourceLabels(val *map[string]*string) {
	if err := j.validateSetResourceLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetSpot(val interface{}) {
	if err := j.validateSetSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spot",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutAdvancedMachineFeatures(value *GoogleContainerClusterNodeConfigAdvancedMachineFeatures) {
	if err := g.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutEphemeralStorageConfig(value *GoogleContainerClusterNodeConfigEphemeralStorageConfig) {
	if err := g.validatePutEphemeralStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralStorageConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutEphemeralStorageLocalSsdConfig(value *GoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfig) {
	if err := g.validatePutEphemeralStorageLocalSsdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralStorageLocalSsdConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutGcfsConfig(value *GoogleContainerClusterNodeConfigGcfsConfig) {
	if err := g.validatePutGcfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcfsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutGuestAccelerator(value interface{}) {
	if err := g.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutGvnic(value *GoogleContainerClusterNodeConfigGvnic) {
	if err := g.validatePutGvnicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGvnic",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutKubeletConfig(value *GoogleContainerClusterNodeConfigKubeletConfig) {
	if err := g.validatePutKubeletConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutLinuxNodeConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfig) {
	if err := g.validatePutLinuxNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutLocalNvmeSsdBlockConfig(value *GoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfig) {
	if err := g.validatePutLocalNvmeSsdBlockConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocalNvmeSsdBlockConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutReservationAffinity(value *GoogleContainerClusterNodeConfigReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutSandboxConfig(value *GoogleContainerClusterNodeConfigSandboxConfig) {
	if err := g.validatePutSandboxConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSandboxConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutShieldedInstanceConfig(value *GoogleContainerClusterNodeConfigShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutSoleTenantConfig(value *GoogleContainerClusterNodeConfigSoleTenantConfig) {
	if err := g.validatePutSoleTenantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoleTenantConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutTaint(value interface{}) {
	if err := g.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTaint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) PutWorkloadMetadataConfig(value *GoogleContainerClusterNodeConfigWorkloadMetadataConfig) {
	if err := g.validatePutWorkloadMetadataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkloadMetadataConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetBootDiskKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetEphemeralStorageConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralStorageConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetEphemeralStorageLocalSsdConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralStorageLocalSsdConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetGcfsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGcfsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetGvnic() {
	_jsii_.InvokeVoid(
		g,
		"resetGvnic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		g,
		"resetImageType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetLinuxNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetLocalNvmeSsdBlockConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalNvmeSsdBlockConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetLocalSsdCount() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetLoggingVariant() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingVariant",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetNodeGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetResourceLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetSandboxConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSandboxConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetSoleTenantConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSoleTenantConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetSpot() {
	_jsii_.InvokeVoid(
		g,
		"resetSpot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetTaint() {
	_jsii_.InvokeVoid(
		g,
		"resetTaint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ResetWorkloadMetadataConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadMetadataConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

